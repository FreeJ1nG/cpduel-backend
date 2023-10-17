package interfaces

import "github.com/FreeJ1nG/cpduel-backend/app/models"

type SubmissionRepository interface {
	CreateSubmission(
		owner string,
		problemId string,
		content string,
		languageId string,
	) (submission models.Submission, err error)
	UpdateSubmissionVerdictInPool(id int, verdict string) (submission *models.Submission, found bool)
	GetSubmissionInPool(id int) (submission *models.Submission, found bool)
}

type SubmissionService interface {
	MakeSubmission(owner string, problemId string, content string, languageId string) (submission models.Submission, err error)
	GetSubmissionInPool(id int) (submission *models.Submission, found bool)
}
