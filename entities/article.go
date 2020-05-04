package entities

import (
	"fmt"
	"time"
)

type Article struct {
	ID         uint
	OriginalID string
	Title      string
	URL        string
	Image      *ArticleImage
	CrawledAt  time.Time
	Site       *Site
	Tags       []*Tag
}

func NewArticle(id uint, originalID string, title string, url string, image *ArticleImage, crawledAt time.Time, site *Site, tags []*Tag) *Article {
	return &Article{
		ID: id,
		OriginalID: originalID,
		Title:      title,
		URL:        url,
		Image:      image,
		CrawledAt:  crawledAt,
		Site:       site,
		Tags:       tags,
	}
}

func (a Article) String() string {
	return fmt.Sprintf("{id: %v, title: %v, url: %v, image: %v, crawledAt: %v, site: %v, tags: %v}", a.ID, a.Title, a.URL, a.Image, a.CrawledAt, a.Site, a.Tags)
}

type ArticleImage struct {
	Image string
}

func NewArticleImage(image string) *ArticleImage {
	if image == "" {
		return nil
	}
	return &ArticleImage{Image: image}
}
