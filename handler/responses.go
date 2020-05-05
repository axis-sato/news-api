package handler

import (
	"github.com/c8112002/news-api/entities"
	"time"
)

type articleResponse struct {
	Id        uint           `json:"id"`
	Title     string         `json:"title"`
	URL       string         `json:"url"`
	Image     string         `json:"image,omitempty"`
	CrawlerAt time.Time      `json:"crawled_at"`
	Tags      []*tagResponse `json:"tags"`
	Site      *siteResponse  `json:"site"`
}

func newArticleResponse(a *entities.Article) *articleResponse {
	sr := newSiteResponse(a.Site)
	tr := newTagSliceResponse(a.Tags)
	resp := &articleResponse{
		Id:        a.ID,
		Title:     a.Title,
		URL:       a.URL,
		Image:     a.Image.Image,
		CrawlerAt: a.CrawledAt,
		Tags:      tr,
		Site:      sr,
	}

	return resp
}

type articlesResponse struct {
	Articles []*articleResponse `json:"articles"`
}

func newArticlesResponse(articles []*entities.Article) *articlesResponse {
	var ar []*articleResponse
	for _, a := range  articles {
		ar = append(ar, newArticleResponse(a))
	}
	return &articlesResponse{Articles: ar}
}


type tagResponse struct {
	ID uint `json:"id"`
	Name string `json:"name"`
}

func newTagResponse(t *entities.Tag) *tagResponse {
	return &tagResponse{
		ID:   t.ID,
		Name: t.Name,
	}
}

func newTagSliceResponse(tags []*entities.Tag) []*tagResponse {
	var tr []*tagResponse
	for _, t := range tags {
		tr= append(tr, newTagResponse(t))
	}
	return tr
}

type tagsResponse struct {
	Tags []*tagResponse `json:"tags"`
}

func newTagsResponse(tags []*entities.Tag) *tagsResponse {
	return &tagsResponse{Tags: newTagSliceResponse(tags)}
}

type siteResponse struct {
	ID uint `json:"id"`
	Name string `json:"name"`
}

func newSiteResponse(site *entities.Site) *siteResponse {
	return &siteResponse{
		ID:   site.ID,
		Name: site.Name,
	}
}
