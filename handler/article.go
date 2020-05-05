package handler

import (
	"github.com/gofiber/fiber"
	"strconv"
)

func (h *Handler) GetArticles(c *fiber.Ctx) {
	limit, _ := getIntQuery(c.Query("limit"), 20)
	after, _ := getIntQuery(c.Query("after"), 0)
	articles, _ := h.as.GetAllArticles(limit, uint(after))
	resp := newArticlesResponse(articles)
	_ = c.JSON(resp)
}

func getIntQuery(query string, defaultLimit int) (int, error) {
	if query == "" {
		return defaultLimit, nil
	}
	return strconv.Atoi(query)
}

