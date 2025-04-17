package model

import (
	"gorm.io/gorm"
)

type QuestionType string

const (
	MULTIPLECHOICE QuestionType = "Multiple_choice"
	SHORTANSWER    QuestionType = "Short_answer"
	TRUEFALSE      QuestionType = "True_false"
	SUBJECTIVE     QuestionType = "Subjective"
)

type Question struct {
	gorm.Model
	ID              	uint 			`gorm:"primaryKey" csv:"id" json:"id"`
	Exam_id				uint			`gorm:"not null" csv:"exam_id" json:"exam_id"`
	Examination     	Examination		`gorm:"foreignKey:Exam_id;references:ID" csv:"-" json:"-"`
	Question_detail 	string			`gorm:"not null" csv:"question_detail" json:"question_detail"`
	Question_type   	QuestionType 	`gorm:"not null" csv:"question_type" json:"question_type"`
	Correct_answer  	string			`gorm:"not null" csv:"correct_answer" json:"correct_answer"`
	Score           	float64			`gorm:"not null" csv:"score" json:"score"`
}
