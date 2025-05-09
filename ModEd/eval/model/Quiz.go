// MEP-1006
package model

import (
	"time"

	"gorm.io/gorm"
)

type Quiz struct {
	gorm.Model
	Title        string         `json:"title" gorm:"not null"`
	Description  string         `json:"description"`
	Status       string         `json:"status" gorm:"default:'draft'"`
	StartDate    time.Time      `json:"start_date"`
	EndDate      time.Time      `json:"end_date"`
	CourseID     uint           `json:"course_id" gorm:"not null"`
	InstructorID uint           `json:"instructor_id" gorm:"not null"`
	Attempts     int            `json:"attempts" gorm:"default:1"`
	TotalScore   float64        `json:"total_score"`
	Questions    []QuizQuestion `json:"questions" gorm:"foreignKey:QuizID"`
}

type QuizQuestion struct {
	gorm.Model
	QuizID        uint         `json:"quiz_id" gorm:"not null"`
	QuestionText  string       `json:"question_text" gorm:"not null"`
	QuestionType  string       `json:"question_type" gorm:"not null"`
	Score         float64      `json:"score"`
	Options       []QuizOption `json:"options" gorm:"foreignKey:QuestionID"`
	CorrectAnswer string       `json:"correct_answer"`
}

type QuizOption struct {
	gorm.Model
	QuestionID uint   `json:"question_id" gorm:"not null"`
	OptionText string `json:"option_text" gorm:"not null"`
	IsCorrect  bool   `json:"is_correct"`
}

type QuizStatus string

const (
	QuizDraft     QuizStatus = "draft"
	QuizPublished QuizStatus = "published"
	QuizHidden    QuizStatus = "hidden"
)
