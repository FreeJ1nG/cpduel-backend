package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/FreeJ1nG/cpduel-backend/app/interfaces"
	"github.com/FreeJ1nG/cpduel-backend/app/models"
	"github.com/FreeJ1nG/cpduel-backend/app/problem"
	"github.com/FreeJ1nG/cpduel-backend/app/webscrapper"
	"github.com/FreeJ1nG/cpduel-backend/db"
	"github.com/FreeJ1nG/cpduel-backend/util"
	"github.com/chromedp/chromedp"
)

func startScrap(ctx context.Context, problemId string, webscrapperService webscrapper.Service, problemRepo interfaces.ProblemRepository) (err error) {
	ctx, cancel := chromedp.NewContext(ctx)
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	problem, err := problemRepo.GetProblemById(problemId)
	if err != nil {
		fmt.Println("Problem does not exist in database, scraping will continue")
	} else {
		fmt.Println("Problem already exist in database, deleting ...")
		err = problemRepo.DeleteProblemById(problemId)
		if err != nil {
			return
		}
		fmt.Println("Problem deleted successfully!")
	}

	problem, _, err = webscrapperService.ScrapProblem(ctx, problemId)
	if err != nil {
		return
	}

	fmt.Println("Now scraping languages for problem ...")
	languages, _, err := webscrapperService.ScrapLanguagesOfProblem(ctx, problemId)
	if err != nil {
		return
	}

	problem, err = problemRepo.CreateProblem(
		problemId,
		problem.Title,
		problem.TimeLimit,
		problem.MemoryLimit,
		problem.InputType,
		problem.OutputType,
		problem.Difficulty,
		problem.Body,
	)
	if err != nil {
		return
	}

	fmt.Printf("Finished scraping problem %s\n", problemId)
	fmt.Println("Id:", problemId)
	fmt.Println("Title:", problem.Title)
	fmt.Println("Time Limit:", problem.TimeLimit)
	fmt.Println("Memory Limit:", problem.MemoryLimit)
	fmt.Println("Input Type:", problem.InputType)
	fmt.Println("Output Type:", problem.OutputType)
	fmt.Println("Difficulty:", problem.Difficulty)
	fmt.Println("Body:")
	fmt.Println(problem.Body)

	languageIds := []string{}
	for _, language := range languages {
		languageIds = append(languageIds, language.Id)
	}

	_, foundIds, err := problemRepo.GetLanguageWithIds(languageIds)
	if err != nil {
		return
	}

	fmt.Println("Languages scraping succeeded!")
	fmt.Println(languages)

	filteredLanguages := []models.Language{}
	for _, language := range languages {
		if !foundIds[language.Id] {
			filteredLanguages = append(filteredLanguages, language)
		}
	}

	err = problemRepo.CreateLanguages(filteredLanguages)
	if err != nil {
		return
	}

	err = problemRepo.CreateProblemLanguages(languages, problemId)
	if err != nil {
		return
	}

	fmt.Println("Language and it's relations saved to database successfully!")
	return
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

	fmt.Println("Enter the lower bound of problem number")
	lowerBound, _ := inputReader.ReadString('\n')
	lowerBound = strings.TrimSuffix(lowerBound, "\n")
	parsedLowerBound, err := strconv.ParseInt(lowerBound, 10, 64)
	if err != nil {
		log.Fatal("unable to parse lower bound: ", err.Error())
		return
	}

	fmt.Println("Enter the upper bound of problem number")
	upperBound, _ := inputReader.ReadString('\n')
	upperBound = strings.TrimSuffix(upperBound, "\n")
	parsedUpperBound, err := strconv.ParseInt(upperBound, 10, 64)
	if err != nil {
		log.Fatal("unable to parse upper bound: ", err.Error())
		return
	}

	allScrapedProblems := []string{}

	for i := parsedLowerBound; i <= parsedUpperBound; i++ {
		characters := []string{
			"A", "B", "C", "D", "E", "F", "G", "H", "I", "J",
		}
		for _, char := range characters {
			problemId := fmt.Sprint(i) + char
			fmt.Printf("Scraping for problem %s ...\n", problemId)
			err := startScrap(ctx, problemId, webscrapperService, problemRepo)
			if err != nil {
				fmt.Printf("unable to scrap problem: %s\n", err.Error())
			} else {
				allScrapedProblems = append(allScrapedProblems, problemId)
			}
			fmt.Println("===========================================================================")
		}
		fmt.Printf("Finished scraping for problems in contest %d\n\n\n", i)
	}

	fmt.Println("Finished scraping process :)")
	fmt.Println("Here are all the problems that has been scraped:")
	for _, problemId := range allScrapedProblems {
		fmt.Println(" >>", problemId)
	}
	os.Exit(0)
}
