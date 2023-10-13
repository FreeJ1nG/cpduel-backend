package models

type Problem struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	TimeLimit   string `json:"time_limit"`
	MemoryLimit string `json:"memory_limit"`
	InputType   string `json:"input_type"`
	OutputType  string `json:"output_type"`
	Body        string `json:"body"`
	Difficulty  int    `json:"difficulty"`
}
