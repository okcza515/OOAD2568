package model

import (
	"ModEd/core"
	"fmt"
	"time"
	"gorm.io/gorm"
)

type Question struct {
	gorm.Model
	ID             uint         `gorm:"primaryKey" csv:"id" json:"id"`
	ExamID         uint         `gorm:"not null" csv:"exam_id" json:"exam_id"`
	QuestionDetail string       `gorm:"not null" csv:"question_detail" json:"question_detail"`
	QuestionType   QuestionType `gorm:"not null" csv:"question_type" json:"question_type"`
	CorrectAnswer string		`gorm:"not null" csv:"correct_answer" json:"correct_answer"`
	Score          float64      `gorm:"not null" csv:"score" json:"score"`
	*core.SerializableRecord
}

type QuestionProductInterface interface {
	GetID() uint
	GetDetail() string
	GetType() QuestionType
	GetCorrectAnswer() string
	GetScore() float64
	Validate() error
}

type QuestionFactory interface {
	CreateQuestion(base Question) QuestionProductInterface
}

type RegularQuestionFactory struct{}

func (f RegularQuestionFactory) CreateQuestion(base Question) QuestionProductInterface {
	return &Question{
		ExamID:         base.ExamID,
		QuestionDetail: base.QuestionDetail,
		QuestionType:   base.QuestionType,
		CorrectAnswer:  base.CorrectAnswer,
		Score:          base.Score,
	}
}

func (q *Question) GetID() uint {
	return q.ID
}

func (q *Question) GetDetail() string {
	return q.QuestionDetail
}

func (q *Question) GetType() QuestionType {
	return q.QuestionType
}

func (q *Question) GetCorrectAnswer() string {
	return q.CorrectAnswer
}

func (q *Question) GetScore() float64 {
	return q.Score
}

func (q *Question) Validate() error {
	if q.QuestionDetail == "" {
		return fmt.Errorf("question detail cannot be empty")
	}
	if q.QuestionType == "" {
		return fmt.Errorf("question type is required")
	}
	if q.Score <= 0 {
		return fmt.Errorf("score must be greater than 0")
	}
	return nil
}