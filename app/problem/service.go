package problem

import (
	"context"
	"net/http"

	"github.com/FreeJ1nG/cpduel-backend/app/dto"
	"github.com/FreeJ1nG/cpduel-backend/app/interfaces"
	"github.com/FreeJ1nG/cpduel-backend/app/webscrapper"
)

type service struct {
	ctx                context.Context
	problemRepo        interfaces.ProblemRepository
	webscrapperService webscrapper.Service
}

func NewService(ctx context.Context, problemRepo interfaces.ProblemRepository, webscrapperService webscrapper.Service) *service {
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
		status = http.StatusNotFound
		return
	}
	languages, err := s.problemRepo.GetLanguagesOfProblemById(problemId)
	if err != nil {
		status = http.StatusInternalServerError
		return
	}
	res = dto.GetProblemResponse{
		Problem:   problem,
		Languages: languages,
	}
	return
}
