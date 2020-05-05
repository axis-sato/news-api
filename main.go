package main

import (
	"context"
	"github.com/c8112002/news-api/db"
	"github.com/c8112002/news-api/handler"
	"github.com/c8112002/news-api/store"
	"github.com/gofiber/fiber"
	"github.com/volatiletech/sqlboiler/boil"
)

//go:generate sqlboiler --wipe --no-tests mysql --no-auto-timestamps

func main() {

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

	app := fiber.New()
	h := handler.NewHandler(as)
	h.Register(app)

	_ = app.Listen(3001)
}
