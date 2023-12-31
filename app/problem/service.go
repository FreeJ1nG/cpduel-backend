package problem

import (
	"context"
	"fmt"
	"net/http"

	"github.com/FreeJ1nG/cpduel-backend/app/dto"
	"github.com/FreeJ1nG/cpduel-backend/app/interfaces"
)

type service struct {
	ctx                context.Context
	problemRepo        interfaces.ProblemRepository
	webscrapperService interfaces.WebscrapperService
}

func NewService(ctx context.Context, problemRepo interfaces.ProblemRepository, webscrapperService interfaces.WebscrapperService) *service {
	return &service{
		ctx:                ctx,
		problemRepo:        problemRepo,
		webscrapperService: webscrapperService,
	}
}

func (s *service) GetProblem(problemId string) (res dto.GetProblemResponse, status int, err error) {
	status = http.StatusOK
	problem, err := s.problemRepo.GetProblemById(problemId)
	if err != nil {
		err = fmt.Errorf("unable to get problem from database: %s", err.Error())
		status = http.StatusNotFound
		return
	}
	languages, err := s.problemRepo.GetLanguagesOfProblemById(problemId)
	if err != nil {
		err = fmt.Errorf("unable to get language from database: %s", err.Error())
		status = http.StatusInternalServerError
		return
	}
	res = dto.GetProblemResponse{
		Problem:   problem,
		Languages: languages,
	}
	return
}
