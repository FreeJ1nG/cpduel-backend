package dto

type SubmitCodePayload struct {
	LanguageId string `json:"language_id"`
	ProblemId  string `json:"problem_id"`
	SourceCode string `json:"source_code"`
}
