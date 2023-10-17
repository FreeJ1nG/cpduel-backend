package models

type Submission struct {
	Id             int    `json:"id"`
	ProblemId      string `json:"problem_id"`
	Owner          string `json:"owner"`
	Content        string `json:"content"`
	LanguageId     string `json:"language_id"`
	SubmittedAt    int64  `json:"submitted_at"`
	Verdict        string `json:"verdict"`
	OJSubmissionId string `json:"oj_submission_id"`
}
