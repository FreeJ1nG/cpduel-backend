package dto

import "github.com/FreeJ1nG/cpduel-backend/app/models"

type GetProblemResponse struct {
	models.Problem
	Languages []models.Language `json:"languages"`
}
