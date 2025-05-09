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
	Attempts     uint           `json:"attempts" gorm:"default:1"`
	TotalScore   float64        `json:"total_score"`
	Questions    []QuizQuestion `json:"questions" gorm:"foreignKey:QuizID"`
}
