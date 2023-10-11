package main

import (
	"context"
	"log"
	"time"

	"github.com/FreeJ1nG/cpduel-backend/app"
	"github.com/FreeJ1nG/cpduel-backend/db"
	"github.com/FreeJ1nG/cpduel-backend/util"
	"github.com/chromedp/chromedp"
)

func login(ctx context.Context) {
	url := "https://codeforces.com/enter"
	if err := chromedp.Run(
		ctx,
		chromedp.Navigate(url),
		chromedp.SendKeys("input[name=handleOrEmail]", "cleopatracpduel@gmail.com"),
		chromedp.SendKeys("input[name=password]", "penyuiscute"),
		chromedp.Submit("input[type=submit]"),
	); err != nil {
		log.Fatal(err)
	}
}

func main() {
	config, err := util.SetConfig()
	if err != nil {
		log.Fatal("Failed to load config", err.Error())
		return
	}

	opts := append(
		chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", config.Headless),
	)

	ctx := context.Background()

	ctx, cancel := chromedp.NewExecAllocator(ctx, opts...)
	defer cancel()

	ctx, cancel = chromedp.NewContext(ctx)
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	mainDB := db.CreatePool(ctx, config.DBDsn)
	db.TestConnection(ctx, mainDB)
	defer func() {
		mainDB.Close()
	}()

	s := app.MakeServer(config, mainDB)

	s.InjectDependencies()
	s.RunServer()
}
