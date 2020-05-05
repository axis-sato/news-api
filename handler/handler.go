package handler

import "github.com/c8112002/news-api/store"

type Handler struct {
	as *store.ArticleStore
}

func NewHandler(as *store.ArticleStore) *Handler {
	return &Handler{as: as}
}