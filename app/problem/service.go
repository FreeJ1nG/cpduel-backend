package problem

import (
	"context"

	"github.com/FreeJ1nG/cpduel-backend/app/webscrapper"
	"github.com/gofiber/fiber/v2"
)

type service struct {
	ctx                context.Context
	problemRepo        Repository
	webscrapperService webscrapper.Service
}

type Service interface {
	GetProblem(problemId string) (res GetProblemResponse, status int, err error)
}

func NewService(ctx context.Context, problemRepo Repository, webscrapperService webscrapper.Service) *service {
	return &service{
		ctx:                ctx,
		problemRepo:        problemRepo,
		webscrapperService: webscrapperService,
	}
}

func (s *service) GetProblem(problemId string) (res GetProblemResponse, status int, err error) {
	status = fiber.StatusOK
	problem, err := s.problemRepo.GetProblemById(problemId)
	if err != nil {
		status = fiber.StatusNotFound
		return
	}
	languages, err := s.problemRepo.GetLanguagesOfProblemById(problemId)
	if err != nil {
		status = fiber.StatusInternalServerError
		return
	}
	res = GetProblemResponse{
		Problem:   problem,
		Languages: languages,
	}
	return
}
