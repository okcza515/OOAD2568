package model

import (
	"time"
)

// Submission represents a student's assignment submission
type Submission struct {
	ID           uint      `gorm:"primaryKey"`
	Content      string    `gorm:"not null"`
	SubmittedAt  time.Time `gorm:"not null"`
	AssignmentID uint      `gorm:"not null"`
	StudentID    uint      `gorm:"not null"`
	Status       string    `gorm:"default:'submitted'"` // submitted, received, evaluated
	Evaluation   *Evaluation
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
