package models

import "time"

type Submission struct {
	Id          int       `json:"id"`
	Owner       string    `json:"owner"`
	Content     string    `json:"content"`
	Language    string    `json:"language"`
	SubmittedAt time.Time `json:"submitted_at"`
	SubmittedBy string    `json:"submitted_by"`
	Verdict     string    `json:"verdict"`
}
