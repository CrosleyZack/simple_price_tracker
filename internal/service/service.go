package service

import (
	"errors"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/anaskhan96/soup"
	"github.com/cenkalti/backoff/v4"
	"github.com/crosleyzack/price_tracker/internal/events"
	event_fsjson "github.com/crosleyzack/price_tracker/internal/events/fsjson"
	"github.com/crosleyzack/price_tracker/internal/items"
	item_fsjson "github.com/crosleyzack/price_tracker/internal/items/fsjson"
	"github.com/crosleyzack/price_tracker/internal/model"
	"github.com/crosleyzack/price_tracker/internal/sites"
	site_fsjson "github.com/crosleyzack/price_tracker/internal/sites/fsjson"
	"github.com/sirupsen/logrus"
)

var (
	FindSplitToken   = "|"
	SearchSplitToken = "."
)

// Service is the main service struct
type Service struct {
	ItemStore  items.IFC
	SiteStore  sites.IFC
	EventStore events.IFC
	Logger     *logrus.Logger
}

// NewService creates a new service from config
func NewService(conf *Config) (s *Service, err error) {
	s = &Service{
		Logger: logrus.New(),
	}
	s.EventStore, err = event_fsjson.New(conf.Event)
	if err != nil {
		return nil, err
	}
	s.ItemStore, err = item_fsjson.New(conf.Item)
	if err != nil {
		return nil, err
	}
	s.SiteStore, err = site_fsjson.New(conf.Site)
	if err != nil {
		return nil, err
	}
	return s, nil
}

// NewEvents checks the current price of all items and adds an event if the price has changed
func (s *Service) NewEvents() error {
	items, err := s.ItemStore.ListItems()
	if err != nil {
		return err
	}
	for _, item := range items {
		price, err := s.EventStore.CurrentPrice(item.Name)
		if err != nil && !errors.Is(err, model.ErrNotFound) {
			s.Logger.Error(err)
			continue
		}
		site, err := s.SiteStore.GetSite(item.Website)
		if err != nil {
			s.Logger.Error(err)
			continue
		}
		siteData, err := getSite(site.URL, item.URIPath)
		if err != nil {
			s.Logger.Error(err)
			continue
		}
		doc := soup.HTMLParse(string(siteData))
		children := doc.Children()
		if len(children) == 0 {
			s.Logger.Error("No children found")
			continue
		}
		html := doc.HTML()
		if html == "" {
			s.Logger.Error("No HTML found")
			continue
		}
		newPrice, err := priceFromHTML(doc, site.PricePath)
		if err != nil {
			s.Logger.Error(err)
			continue
		}
		if float32(newPrice) != price {
			err = s.EventStore.AddEvent(item.Name, model.Event{
				Price: float32(newPrice),
				Date:  time.Now().UTC(),
			})
			if err != nil {
				s.Logger.Error(err)
				continue
			}
		}
	}
	return nil
}

// priceFromHTML finds the price from the HTML using the pricePath.]
//
//	pricePath is a series of separate searches separated by FindSplitToken
//	each search is a series of tags separated by SearchSplitToken
//	if the final item in the search is a number, it is the index of the html object to grab. defaults to 0
func priceFromHTML(html soup.Root, pricePath string) (float32, error) {
	searches := strings.Split(pricePath, FindSplitToken)
	for _, search := range searches {
		index := 0
		searchStrs := strings.Split(search, SearchSplitToken)
		// if final item in the search is a number, it is the index
		if i, err := strconv.ParseInt(searchStrs[len(searchStrs)-1], 10, 32); err == nil {
			searchStrs = searchStrs[:len(searchStrs)-1]
			index = int(i)
		}
		roots := html.FindAllStrict(searchStrs...)
		if len(roots) == 0 {
			return -1, errors.New("No roots found")
		}
		// minimum value is -1 * len(roots)
		if index < (-1 * len(roots)) {
			index = -1 * len(roots)
		}
		// maximum value is len(roots) - 1
		if index > len(roots)-1 {
			index = len(roots) - 1
		}
		// if index is negative, it is from the end of the search
		if index < 0 {
			index = len(roots) + index
		}
		html = roots[index]
	}
	text := cleanText(html.Text())
	newPrice, err := strconv.ParseFloat(text, 32)
	if err != nil {
		return -1, err
	}
	return float32(newPrice), nil
}

func getSite(baseURL, path string) ([]byte, error) {
	uri, err := url.JoinPath(baseURL, path)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:135.0) Gecko/20100101 Firefox/135.0")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")
	req.Header.Set("Accept", "image/avif,image/jxl,image/webp,image/png,image/svg+xml,image/*;q=0.8,*/*;q=0.5")
	req.Header.Set("Accept-Encoding", "identity")
	req.Header.Set("Host", baseURL)
	req.Header.Set("Sec-Fetch-Dest", "document")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Sec-Fetch-Site", "none")
	req.Header.Set("Sec-Fetch-User", "?1")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	var response *http.Response
	var cookies []*http.Cookie
	err = backoff.Retry(
		func() error {
			for _, cookie := range cookies {
				req.AddCookie(cookie)
			}
			response, err = http.DefaultClient.Do(req)
			if err != nil {
				return err
			}
			cookies = response.Cookies()
			return err
		},
		backoff.NewExponentialBackOff(),
	)
	if err != nil {
		return nil, err
	}
	resBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return resBody, nil
}

func cleanText(text string) string {
	text = strings.Trim(text, " \t\n")
	text = strings.ReplaceAll(text, "$", "")
	return text
}
