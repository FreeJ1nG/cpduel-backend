package webscrapper

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/FreeJ1nG/cpduel-backend/app/models"
	"github.com/chromedp/chromedp"
	"github.com/gofiber/fiber/v2"
)

type service struct {
	ctx context.Context
}

type Service interface {
	ScrapProblem(ctx context.Context, problemId string) (problem models.Problem, status int, err error)
	ScrapLanguagesOfProblem(ctx context.Context, problemId string) (languages []models.Language, status int, err error)
}

func NewService(ctx context.Context) *service {
	return &service{
		ctx: ctx,
	}
}

var CODEFORCES_STRING string = "https://codeforces.com"

func parseProblemId(problemId string) (problemNumber string, problemCode string) {
	problemNumber = problemId[:len(problemId)-1]
	problemCode = string(problemId[len(problemId)-1])
	return
}

func (s *service) loginToBotAccount(ctx context.Context) (err error) {
	if err = chromedp.Run(
		ctx,
		chromedp.WaitVisible(`a[href^="/enter?back"]`, chromedp.ByQuery),
		chromedp.Click(`a[href^="/enter?back"]`, chromedp.NodeVisible),
	); err != nil {
		err = fmt.Errorf("unable to locate login button: %s", err.Error())
		return
	}

	if err = chromedp.Run(
		ctx,
		chromedp.Sleep(time.Second),
		chromedp.WaitVisible("input[name=handleOrEmail]"),
		chromedp.SendKeys("input[name=handleOrEmail]", "cleopatracpduel@gmail.com", chromedp.NodeVisible),
	); err != nil {
		err = fmt.Errorf("unable to insert email/handle: %s", err.Error())
		return
	}

	if err = chromedp.Run(
		ctx,
		chromedp.WaitVisible("input[name=password]"),
		chromedp.SendKeys("input[name=password]", "penyuiscute", chromedp.NodeVisible),
	); err != nil {
		err = fmt.Errorf("unable to insert password: %s", err.Error())
		return
	}

	if err = chromedp.Run(
		ctx,
		chromedp.WaitVisible("input[type=submit]"),
		chromedp.Submit("input[type=submit]"),
	); err != nil {
		err = fmt.Errorf("unable to login: %s", err.Error())
	}
	return
}

func (s *service) ScrapLanguagesOfProblem(ctx context.Context, problemId string) (languages []models.Language, status int, err error) {
	status = fiber.StatusOK
	problemNumber, problemCode := parseProblemId(problemId)

	url := fmt.Sprintf("%s/contest/%s/problem/%s", CODEFORCES_STRING, problemNumber, problemCode)
	if err = chromedp.Run(ctx, chromedp.Navigate(url)); err != nil {
		err = fmt.Errorf("unable to open problem: %s", err.Error())
		status = fiber.StatusBadRequest
		return
	}

	err = s.loginToBotAccount(ctx)
	if err != nil {
		status = fiber.StatusInternalServerError
		return
	}

	if err = chromedp.Run(
		ctx,
		chromedp.WaitVisible("select[name=programTypeId]"),
		chromedp.Sleep(time.Second),
		chromedp.Evaluate(
			`Array.from(document.querySelectorAll('select[name=programTypeId] > *')).map(element => ({
				id: element.getAttribute('value'), 
				value: element.innerHTML,
			}))`,
			&languages,
		),
	); err != nil {
		err = fmt.Errorf("unable to scrap language: %s", err.Error())
		status = fiber.StatusInternalServerError
		return
	}

	return
}

func (s *service) ScrapProblem(ctx context.Context, problemId string) (problem models.Problem, status int, err error) {
	status = fiber.StatusOK
	problemNumber, problemCode := parseProblemId(problemId)

	url := fmt.Sprintf("%s/contest/%s/problem/%s", CODEFORCES_STRING, problemNumber, problemCode)

	var title string
	var timeLimit string
	var memoryLimit string
	var inputType string
	var outputType string
	var difficulty string
	var body string

	if err = chromedp.Run(
		ctx,
		chromedp.Navigate(url),
	); err != nil {
		err = fmt.Errorf("unable to open problem: %s", err.Error())
		status = fiber.StatusBadRequest
		return
	}

	var currentUrl string
	if err = chromedp.Run(
		ctx,
		chromedp.Sleep(time.Second),
		chromedp.Location(&currentUrl),
	); err != nil {
		err = fmt.Errorf("unable to get page location: %s", err.Error())
		status = fiber.StatusInternalServerError
		return
	}

	if currentUrl != url {
		err = fmt.Errorf("problem page doesn't exist")
		status = fiber.StatusBadRequest
		return
	}

	if err = chromedp.Run(
		ctx,
		chromedp.Text(".header > .title", &title, chromedp.ByQuery),
		chromedp.Text(".header > .time-limit", &timeLimit, chromedp.ByQuery),
		chromedp.Text(".header > .memory-limit", &memoryLimit, chromedp.ByQuery),
		chromedp.Text(".header > .input-file", &inputType, chromedp.ByQuery),
		chromedp.Text(".header > .output-file", &outputType, chromedp.ByQuery),
		chromedp.Text("span.tag-box[title=Difficulty]", &difficulty, chromedp.ByQuery),
		chromedp.Evaluate(
			`Array.from(document.querySelectorAll('.problem-statement > *:not(.header)')).map(element => element.outerHTML).reduce((acc, cur) => acc + cur, "");`,
			&body,
		),
	); err != nil {
		err = fmt.Errorf("unable to scrap problem: %s", err.Error())
		status = fiber.StatusInternalServerError
		return
	}

	title = title[3:]
	timeLimit = strings.TrimSpace(strings.Replace(timeLimit, "time limit per test", "", 1))
	memoryLimit = strings.TrimSpace(strings.Replace(memoryLimit, "memory limit per test", "", 1))
	inputType = strings.TrimSpace(strings.Replace(inputType, "input", "", 1))
	outputType = strings.TrimSpace(strings.Replace(outputType, "output", "", 1))
	difficulty = strings.TrimSpace(difficulty)[1:]
	parsedDifficulty, err := strconv.ParseInt(difficulty, 10, 32)
	if err != nil {
		err = fmt.Errorf("unable to parse difficulty: %s", err.Error())
		status = fiber.StatusInternalServerError
		return
	}

	problem = models.Problem{
		Id:          problemId,
		Title:       title,
		TimeLimit:   timeLimit,
		MemoryLimit: memoryLimit,
		InputType:   inputType,
		OutputType:  outputType,
		Difficulty:  int(parsedDifficulty),
		Body:        body,
	}
	return
}
