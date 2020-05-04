package entities

import (
	"fmt"
)

type Site struct {
	ID   uint
	Name string
	URL  string
}

func (s Site) String() string {
	return fmt.Sprintf("{id: %v, name: %v, url: %v}", s.ID, s.Name, s.URL)
}

func NewSite(id uint, name string, url string) *Site {
	return &Site{
		ID: id,
		Name: name,
		URL: url,
	}
}
