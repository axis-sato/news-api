package handler

import (
	"time"

	"github.com/c8112002/news-api/entities"
)

const (
	internalServerErrorCode    = "INTERNAL_SERVER_ERROR"
	internalServerErrorMessage = "予期せぬエラーが発生しました。"
)

type articleResponse struct {
	Id        uint           `json:"id"`
	Title     string         `json:"title"`
	URL       string         `json:"url"`
	Image     string         `json:"image,omitempty"`
	CrawlerAt time.Time      `json:"crawledAt"`
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
	ar := make([]*articleResponse, 0)
	for _, a := range articles {
		ar = append(ar, newArticleResponse(a))
	}
	return &articlesResponse{Articles: ar}
}

type tagResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func newTagResponse(t *entities.Tag) *tagResponse {
	return &tagResponse{
		ID:   t.ID,
		Name: t.Name,
	}
}

func newTagSliceResponse(tags []*entities.Tag) []*tagResponse {
	tr := make([]*tagResponse, 0)
	for _, t := range tags {
		tr = append(tr, newTagResponse(t))
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
	ID   uint   `json:"id"`
	Name string `json:"name"`
	URL  string `json:"url"`
}

func newSiteResponse(site *entities.Site) *siteResponse {
	return &siteResponse{
		ID:   site.ID,
		Name: site.Name,
		URL:  site.URL,
	}
}

type errorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func newErrorResponse(code string, message string) *errorResponse {
	return &errorResponse{
		Code:    code,
		Message: message,
	}
}

func newInternalServerErrorResponse() *errorsResponse {
	resp := &errorResponse{
		Code:    internalServerErrorCode,
		Message: internalServerErrorMessage,
	}
	return newErrorsResponse([]*errorResponse{resp})
}

type errorsResponse struct {
	Errors []*errorResponse `json:"errors"`
}

func newErrorsResponse(resp []*errorResponse) *errorsResponse {
	return &errorsResponse{Errors: resp}
}
