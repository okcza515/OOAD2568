// MEP-1007
package model

import (
	"time"

	"gorm.io/gorm"
)

type Answer struct {
	gorm.Model
	StudentID   uint      `json:"student_id"`
	ExamID      uint      `json:"exam_id"`
	StartTime   time.Time `json:"started_at"`
	SubmittedAt time.Time `json:"submitted_at"`
}
