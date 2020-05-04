package main

import (
	"context"
	"github.com/c8112002/news-api/db"
	"github.com/c8112002/news-api/store"
	"github.com/gofiber/fiber"
	"github.com/volatiletech/sqlboiler/boil"
)

//go:generate sqlboiler --wipe --no-tests mysql --no-auto-timestamps

func main() {
	app := fiber.New()

	boil.DebugMode = true

	d, err := db.NewDB()
	if err != nil {
		panic(err.Error())
	}

	defer func() {
		if err = d.Close(); err != nil {
			panic(err.Error())
		}
	}()

	ctx := context.Background()
	as := store.NewArticleStore(d, ctx)

	api := app.Group("/api")
	v1 := api.Group("/v1")
	v1.Get("/", func(c *fiber.Ctx) {
		articles, _ := as.GetAllArticles()
		_ = c.JSON(articles)
	})

	_ = app.Listen(3001)
}
