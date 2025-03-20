package model

import (
	"gorm.io/gorm"
)

type QuestionType string

const (
	MULTIPLECHOICE QuestionType = "Multiple_choice"
	SHORTANSWER    QuestionType = "Short_answer"
	TRUEFALSE      QuestionType = "True_false"
)

type Question struct {
	gorm.Model
	ID              uint `gorm:"primaryKey"`
	Examination     Examination
	Question_detail string
	Question_type   QuestionType
	Correct_answer  string
	Score           float64
}
