package model

type SubjectiveAnswer struct {
	BaseAnswer
	AnswerText string `json:"content"`
}
