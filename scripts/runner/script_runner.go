package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/FreeJ1nG/cpduel-backend/app/problem"
	"github.com/FreeJ1nG/cpduel-backend/app/webscrapper"
	"github.com/FreeJ1nG/cpduel-backend/db"
	"github.com/FreeJ1nG/cpduel-backend/scripts"
	"github.com/FreeJ1nG/cpduel-backend/util"
	"github.com/chromedp/chromedp"
)

const (
	ScrapProblem      = "scrap_problems"
	SetLanguageConfig = "set_language_config"
)

var scriptOptions = []string{
	ScrapProblem,
	SetLanguageConfig,
}

func main() {
	config, err := util.SetConfig()
	if err != nil {
		log.Fatal("failed to load config: ", err.Error())
		return
	}

	opts := append(
		chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", true),
	)

	ctx := context.Background()

	ctx, cancel := chromedp.NewExecAllocator(ctx, opts...)
	defer cancel()

	mainDB := db.CreatePool(config.DBDsn)
	db.TestConnection(mainDB)
	defer mainDB.Close()

	problemRepo := problem.NewRepository(mainDB)

	webscrapperService := webscrapper.NewService(ctx)

	inputReader := bufio.NewReader(os.Stdin)

	fmt.Println("Script Options:")
	for _, option := range scriptOptions {
		fmt.Println(" -", option)
	}
	fmt.Println("Enter the script you want to run")
	scriptName, _ := inputReader.ReadString('\n')
	scriptName = strings.TrimSuffix(scriptName, "\n")

	if scriptName == "scrap_problems" {
		scripts.ScrapProblem(ctx, webscrapperService, problemRepo)
	} else if scriptName == "set_language_config" {
		scripts.SetLanguageConfig(ctx, mainDB)
	}
}
