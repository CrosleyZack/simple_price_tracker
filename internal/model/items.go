package model

import "fmt"

// Item is the defining items, the website they are on, and the uri path
type Item struct {
	Name    string `json:"name"`
	Website string `json:"website"`
	// expects items joined by "."
	URIPath string `json:"path"`
}

func (i Item) String() string {
	return fmt.Sprintf("%s : %s : %s", i.Name, i.Website, i.URIPath)
}
