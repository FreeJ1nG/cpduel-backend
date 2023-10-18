package main

import (
	"context"
	"log"

	"github.com/FreeJ1nG/cpduel-backend/app"
	"github.com/FreeJ1nG/cpduel-backend/db"
	"github.com/FreeJ1nG/cpduel-backend/util"
	"github.com/chromedp/chromedp"
)

func main() {
	config, err := util.SetConfig()
	if err != nil {
		log.Fatal("Failed to load config", err.Error())
		return
	}

	var headlessFlag func(*chromedp.ExecAllocator)
	if config.Headless || config.Env != "local" {
		headlessFlag = chromedp.Flag("headless", true)
	} else {
		headlessFlag = chromedp.Flag("headless", false)
	}

	opts := append(
		chromedp.DefaultExecAllocatorOptions[:],
		headlessFlag,
	)

	ctx := context.Background()

	ctx, cancel := chromedp.NewExecAllocator(ctx, opts...)
	defer cancel()

	mainDB := db.CreatePool(config.DBDsn)
	db.TestConnection(mainDB)
	defer func() {
		mainDB.Close()
	}()

	s := app.MakeServer(config, mainDB)

	s.InjectDependencies(ctx)
	s.RunServer()
}
