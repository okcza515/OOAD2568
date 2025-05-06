package model

import "gorm.io/gorm"

type BaseAnswer struct {
	gorm.Model
	SubmissionID uint `json:"submission_id"`
	QuestionID   uint `json:"question_id"`
}

func (a *BaseAnswer) GetQuestionID() uint {
	return a.QuestionID
}

func (a *BaseAnswer) GetSubmissionID() uint {
	return a.SubmissionID
}
