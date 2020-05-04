package store

import (
	"github.com/c8112002/news-api/entities"
	"github.com/c8112002/news-api/models"
)

func newSite(s *models.Site) *entities.Site {
	return entities.NewSite(s.ID, s.Name, s.URL)
}