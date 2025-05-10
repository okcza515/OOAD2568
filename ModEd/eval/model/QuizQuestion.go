// MEP-1006
package model

import (
	"gorm.io/gorm"
)

type QuizQuestion struct {
	gorm.Model
	QuizID        uint         `json:"quiz_id" gorm:"not null"`
	QuestionText  string       `json:"question_text" gorm:"not null"`
	QuestionType  string       `json:"question_type" gorm:"not null"`
	Score         float64      `json:"score"`
	Options       []QuizOption `json:"options" gorm:"foreignKey:QuestionID"`
	CorrectAnswer string       `json:"correct_answer"`
}
