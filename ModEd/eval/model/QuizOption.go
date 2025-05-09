// MEP-1006
package model

import (
	"gorm.io/gorm"
)

type QuizOption struct {
	gorm.Model
	QuestionID uint   `json:"question_id" gorm:"not null"`
	OptionText string `json:"option_text" gorm:"not null"`
	IsCorrect  bool   `json:"is_correct"`
}
