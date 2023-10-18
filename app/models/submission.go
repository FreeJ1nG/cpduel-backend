package models

type Submission struct {
	Id             int    `json:"id"`
	ProblemId      string `json:"problemId"`
	Owner          string `json:"owner"`
	Content        string `json:"content"`
	LanguageId     string `json:"languageId"`
	SubmittedAt    int64  `json:"submittedAt"`
	Verdict        string `json:"verdict"`
	OJSubmissionId string `json:"ojSubmissionId"`
}
