package submission

import (
	"fmt"
	"net/http"

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

func (s *service) GetSubmissionsOfUser(username string) (submissions []models.PublicSubmission, status int, err error) {
	status = http.StatusOK
	submissions, err = s.submissionRepo.GetSubmissionsOfUser(username)
	if err != nil {
		err = fmt.Errorf("unable to get submission of user: %s", err.Error())
		status = http.StatusInternalServerError
		return
	}
	return
}

func (s *service) UpdateSubmissionVerdict(submissionId int, verdict string) (err error) {
	return s.submissionRepo.UpdateSubmissionVerdict(submissionId, verdict)
}
