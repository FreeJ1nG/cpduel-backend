package app

import (
	"context"

	"github.com/FreeJ1nG/cpduel-backend/app/problem"
	"github.com/FreeJ1nG/cpduel-backend/app/webscrapper"
	"github.com/FreeJ1nG/cpduel-backend/util"
	"github.com/gofiber/fiber/v2"
)

func (s *Server) InjectDependencies(ctx context.Context) {
	s.router.Use(util.LoggerMiddleware())

	// Repositories
	problemRepo := problem.NewRepository(s.db)

	// Services
	webscrapperService := webscrapper.NewService(ctx)
	problemService := problem.NewService(ctx, problemRepo, webscrapperService)

	// Controllers
	problemHandler := problem.NewHandler(problemService)

	s.router.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "ok",
		})
	})

	problemRouter := s.router.Group("/problems")
	problemRouter.Get("/:id", problemHandler.GetProblem)
}
