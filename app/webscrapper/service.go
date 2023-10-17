package webscrapper

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/FreeJ1nG/cpduel-backend/app/models"
	"github.com/FreeJ1nG/cpduel-backend/util"
	"github.com/chromedp/chromedp"
)

type service struct {
	ctx context.Context
}

func NewService(ctx context.Context) *service {
	return &service{
		ctx: ctx,
	}
}

var CODEFORCES_STRING string = "https://codeforces.com"

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
		chromedp.Sleep(1500*time.Millisecond),
	); err != nil {
		err = fmt.Errorf("unable to login: %s", err.Error())
	}
	return
}

func (s *service) ScrapLanguagesOfProblem(ctx context.Context, problemId string) (languages []models.Language, status int, err error) {
	status = http.StatusOK
	problemNumber, problemCode := util.ParseProblemId(problemId)

	url := fmt.Sprintf("%s/contest/%s/problem/%s", CODEFORCES_STRING, problemNumber, problemCode)
	if err = chromedp.Run(ctx, chromedp.Navigate(url)); err != nil {
		err = fmt.Errorf("unable to open problem: %s", err.Error())
		status = http.StatusBadRequest
		return
	}

	err = s.loginToBotAccount(ctx)
	if err != nil {
		status = http.StatusInternalServerError
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
		status = http.StatusInternalServerError
		return
	}

	return
}

func (s *service) ScrapProblem(ctx context.Context, problemId string) (problem models.Problem, status int, err error) {
	status = http.StatusOK
	problemNumber, problemCode := util.ParseProblemId(problemId)

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
		status = http.StatusBadRequest
		return
	}

	var currentUrl string
	if err = chromedp.Run(
		ctx,
		chromedp.Sleep(time.Second),
		chromedp.Location(&currentUrl),
	); err != nil {
		err = fmt.Errorf("unable to get page location: %s", err.Error())
		status = http.StatusInternalServerError
		return
	}

	if currentUrl != url {
		err = fmt.Errorf("problem page doesn't exist")
		status = http.StatusBadRequest
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
		status = http.StatusInternalServerError
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
		status = http.StatusInternalServerError
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

func (s *service) SubmitCode(
	ctx context.Context,
	submission *models.Submission,
	problemId string,
	sourceCode string,
	languageId string,
) (err error) {
	problemNumber, problemCode := util.ParseProblemId(problemId)

	url := fmt.Sprintf("%s/contest/%s", CODEFORCES_STRING, problemNumber)
	if err = chromedp.Run(
		ctx,
		chromedp.Navigate(url),
	); err != nil {
		err = fmt.Errorf("contest page does not exist: %s", err.Error())
		return
	}

	err = s.loginToBotAccount(ctx)
	if err != nil {
		return
	}

	url = fmt.Sprintf("%s/contest/%s/submit", CODEFORCES_STRING, problemNumber)
	if err = chromedp.Run(
		ctx,
		chromedp.Navigate(url),
	); err != nil {
		err = fmt.Errorf("unable to navigate to submission page: %s", err.Error())
		return
	}

	if err = chromedp.Run(
		ctx,
		chromedp.WaitVisible(`select[name="submittedProblemIndex"]`, chromedp.ByQuery),
		chromedp.Sleep(100*time.Millisecond),
		chromedp.SetValue(`select[name="submittedProblemIndex"]`, problemCode, chromedp.BySearch),
		chromedp.Sleep(100*time.Millisecond),
		chromedp.SetValue(`select[name="programTypeId"]`, languageId, chromedp.BySearch),
		chromedp.Sleep(100*time.Millisecond),
		chromedp.SetValue(`#sourceCodeTextarea`, sourceCode, chromedp.ByID),
		chromedp.Sleep(100*time.Millisecond),
		chromedp.Submit(`#singlePageSubmitButton`, chromedp.ByID),
	); err != nil {
		err = fmt.Errorf("unable to submit code: %s", err.Error())
		return
	}

	var duplicate bool = false

	chromedp.Run(
		ctx,
		chromedp.Sleep(300*time.Millisecond),
		chromedp.Text(`.for__source`, new(string), chromedp.ByQuery, chromedp.AtLeast(0)),
		chromedp.ActionFunc(func(ctx context.Context) error {
			duplicate = true
			return nil
		}),
	)

	if duplicate {
		err = fmt.Errorf("unable to submit code: duplicate submission")
		return
	}

	url = fmt.Sprintf("%s/contest/%s/my", CODEFORCES_STRING, problemNumber)

	if err = chromedp.Run(
		ctx,
		chromedp.Sleep(150*time.Millisecond),
		chromedp.Text(`tbody > tr:nth-child(2) > td.id-cell`, &submission.OJSubmissionId, chromedp.ByQuery),
	); err != nil {
		err = fmt.Errorf("unable to locate submission id: %s", err.Error())
		return
	}

	for {
		if err = chromedp.Run(
			ctx,
			chromedp.Sleep(150*time.Millisecond),
			chromedp.Navigate(url),
			chromedp.Text(
				fmt.Sprintf(`tbody > tr[data-submission-id="%s"] > td.status-verdict-cell`, submission.OJSubmissionId),
				&submission.Verdict,
				chromedp.ByQuery,
			),
		); err != nil {
			err = fmt.Errorf("unable to get submission verdict: %s", err.Error())
			return
		}

		var waiting string
		var exists bool

		if err = chromedp.Run(
			ctx,
			chromedp.AttributeValue(fmt.Sprintf(`tbody > tr[data-submission-id="%s"] > td.status-verdict-cell`, submission.OJSubmissionId), "waiting", &waiting, &exists, chromedp.ByQuery),
		); err != nil {
			err = fmt.Errorf("unable to locate waiting status: %s", err.Error())
			return
		}

		if exists && waiting == "false" {
			return
		}
	}
}
