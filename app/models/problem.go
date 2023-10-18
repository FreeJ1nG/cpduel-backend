package models

type Problem struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	TimeLimit   string `json:"timeLimit"`
	MemoryLimit string `json:"memoryLimit"`
	InputType   string `json:"inputType"`
	OutputType  string `json:"outputType"`
	Body        string `json:"body"`
	Difficulty  int    `json:"difficulty"`
}
