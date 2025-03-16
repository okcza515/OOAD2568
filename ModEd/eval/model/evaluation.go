package model

import (
	"time"
)

// Evaluation represents an instructor's evaluation of a student's submission
type Evaluation struct {
	ID           uint    `gorm:"primaryKey"`
	Score        float64 `gorm:"not null"`
	Feedback     string
	EvaluatedAt  time.Time `gorm:"not null"`
	SubmissionID uint      `gorm:"uniqueIndex;not null"`
	InstructorID uint      `gorm:"not null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
