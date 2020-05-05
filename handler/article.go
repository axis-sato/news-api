package handler

import "github.com/gofiber/fiber"

func (h *Handler) GetArticles(c *fiber.Ctx) {
	articles, _ := h.as.GetAllArticles()
	resp := newArticlesResponse(articles)
	_ = c.JSON(resp)
}