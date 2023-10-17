package interfaces

import (
	"context"

	"github.com/FreeJ1nG/cpduel-backend/app/models"
)

type WebscrapperService interface {
	ScrapProblem(ctx context.Context, problemId string) (problem models.Problem, status int, err error)
	ScrapLanguagesOfProblem(ctx context.Context, problemId string) (languages []models.Language, status int, err error)
	SubmitCode(
		ctx context.Context,
		submission *models.Submission,
		problemId string,
		sourceCode string,
		languageId string,
	) (err error)
}
