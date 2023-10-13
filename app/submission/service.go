package submission

import "github.com/FreeJ1nG/cpduel-backend/app/models"

type service struct {
}

type Service interface {
}

func NewService() *service {
	return &service{}
}

func (s *service) SubmitCode() (submission models.Submission, err error) {
	return
}
