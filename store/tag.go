package store

import (
	"github.com/c8112002/news-api/entities"
	"github.com/c8112002/news-api/models"
)

func newTags(ts models.TagSlice) []*entities.Tag {
	var tags []*entities.Tag
	for _, t := range ts {
		tag := newTag(t)
		tags = append(tags, tag)
	}
	return tags
}

func newTag(t *models.Tag) *entities.Tag {
	return entities.NewTag(t.ID, t.Name)
}
