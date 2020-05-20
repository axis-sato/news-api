package main

import (
	"context"
	"os"
	"time"

	"github.com/c8112002/news-api/db"
	"github.com/c8112002/news-api/handler"
	"github.com/c8112002/news-api/store"
	"github.com/gofiber/fiber"
	"github.com/gofiber/recover"
	log "github.com/sirupsen/logrus"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

//go:generate sqlboiler --wipe --no-tests mysql --no-auto-timestamps

func init() {
	log.SetLevel(log.DebugLevel)
	log.SetReportCaller(true)
	log.SetOutput(os.Stdout)
	time.Local = time.FixedZone("Asia/Tokyo", 9*60*60)
}

func main() {

	boil.DebugMode = true
	boil.SetLocation(time.Local)

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
	app.Use(recover.New(recover.Config{
		Handler: handler.ErrorHandler,
	}))

	h := handler.NewHandler(as)
	h.Register(app)

	_ = app.Listen(3001)
}
