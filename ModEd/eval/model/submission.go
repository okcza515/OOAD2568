package model

import (
	"time"

	"gorm.io/gorm"

	"ModEd/common/model"
)

// Submission represents a student's assignment submission
type Submission struct {
	gorm.Model
	ID          uint          `gorm:"primaryKey"`
	Content     string        `gorm:"not null"`
	SubmittedAt time.Time     `gorm:"not null"`
	SID         model.Student `gorm:"foreignKey:SID"`
	Status      string        `gorm:"default:'submitted'"` // submitted, received, evaluated
	Evaluation  *Evaluation
	CreatedAt   time.Time
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}
