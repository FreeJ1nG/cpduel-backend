package dto

import "github.com/FreeJ1nG/cpduel-backend/app/models"

type WebsocketMessageResponse struct {
	Type  string        `json:"type"`
	Data  *ResponseData `json:"data,omitempty"`
	Error string        `json:"error,omitempty"`
}

type ResponseData struct {
	models.Submission
}

const (
	SubmissionStateMessageType string = "submission-state"
	ErrorMessageType           string = "error"
)
