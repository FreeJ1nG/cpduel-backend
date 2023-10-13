package problem

import (
	"github.com/gofiber/fiber/v2"
)

type handler struct {
	problemService Service
}

type Handler interface {
}

func NewHandler(problemService Service) *handler {
	return &handler{
		problemService: problemService,
	}
}

func (h *handler) GetProblem(c *fiber.Ctx) error {
	id := c.Params("id")

	res, status, err := h.problemService.GetProblem(id)
	if err != nil {
		return c.Status(status).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(status).JSON(res)
}
