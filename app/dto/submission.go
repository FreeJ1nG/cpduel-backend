package dto

import "github.com/FreeJ1nG/cpduel-backend/app/models"

type GetSubmissionsOfUserResponse struct {
	Submissions []models.PublicSubmission `json:"submissions"`
}
