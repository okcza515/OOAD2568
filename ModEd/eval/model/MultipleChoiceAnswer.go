package model

type MultipleChoiceAnswer struct {
	BaseAnswer
	Choices []string `json:"choices"`
}
