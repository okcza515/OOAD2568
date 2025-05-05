// MEP-1007
package model

import (
	"time"

	"gorm.io/gorm"
)

type AnswerProductInterface interface {
	GetQuestionID() uint
	GetSubmissionID() uint
}

type AnswerFactory interface {
	NewAnswer(questionID, submissionID uint) AnswerProductInterface
}

type Answer struct {
	gorm.Model
	StudentID   uint      `json:"student_id"`
	ExamID      uint      `json:"exam_id"`
	SubmittedAt time.Time `json:"submitted_at"`
	IsLocked    bool      `json:"is_locked"`
}

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
