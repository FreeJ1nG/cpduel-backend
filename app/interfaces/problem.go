package interfaces

import (
	"net/http"

	"github.com/FreeJ1nG/cpduel-backend/app/dto"
	"github.com/FreeJ1nG/cpduel-backend/app/models"
)

type ProblemHandler interface {
	GetProblem(w http.ResponseWriter, r *http.Request)
}

type ProblemService interface {
	GetProblem(problemId string) (res dto.GetProblemResponse, status int, err error)
}

type ProblemRepository interface {
	CreateProblem(
		problemId string,
		title string,
		timeLimit string,
		memoryLimit string,
		inputType string,
		outputType string,
		difficulty int,
		body string,
	) (problem models.Problem, err error)
	CreateLanguages(languages []models.Language) (err error)
	CreateProblemLanguages(languages []models.Language, problemId string) (err error)
	GetProblemById(problemId string) (problem models.Problem, err error)
	GetLanguageWithIds(languageIds []string) (languages []models.Language, foundIds map[string]bool, err error)
	GetLanguagesOfProblemById(problemId string) (languages []models.Language, err error)
	DeleteProblemById(problemId string) (err error)
}
