package model

import "fmt"

type Website struct {
	Name      string `json:"name"`
	URL       string `json:"url"`
	PricePath string `json:"price_path"`
}

func (w Website) String() string {
	return fmt.Sprintf("%s : %s : %s", w.Name, w.URL, w.PricePath)
}
