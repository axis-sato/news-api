package handler

import "github.com/gofiber/fiber"

func (h *Handler) Register(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/v1")
	articles := v1.Group("/articles")

	articles.Get("/", h.GetArticles)
}