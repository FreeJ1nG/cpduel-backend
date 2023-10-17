package submission

import (
	"fmt"

	"github.com/FreeJ1nG/cpduel-backend/app/interfaces"
	"github.com/FreeJ1nG/cpduel-backend/app/models"
)

type service struct {
	submissionRepo     interfaces.SubmissionRepository
	webscrapperService interfaces.WebscrapperService
}

func NewService(
	submissionRepo interfaces.SubmissionRepository,
	webscrapperService interfaces.WebscrapperService,
) *service {
	return &service{
		submissionRepo:     submissionRepo,
		webscrapperService: webscrapperService,
	}
}

func (s *service) MakeSubmission(owner string, problemId string, content string, languageId string) (submission models.Submission, err error) {
	submission, err = s.submissionRepo.CreateSubmission(owner, problemId, content, languageId)
	submission.Verdict = "Pending"
	if err != nil {
		err = fmt.Errorf("unable to create submission: %s", err.Error())
		return
	}
	return
}

func (s *service) GetSubmissionInPool(id int) (submission *models.Submission, found bool) {
	return s.submissionRepo.GetSubmissionInPool(id)
}
