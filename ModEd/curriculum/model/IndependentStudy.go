package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IndependentStudy struct {
	gorm.Model
	IndependentStudyId      uuid.UUID  `gorm:"primaryKey;unique"`
	IndependentStudyTopic   string     `gorm:"not null"`
	IndependentStudyContent string     `gorm:"not null"`
	EvaluationFromAdvisor   int        `gorm:"not null;default:0"`
	AssignmentDate          time.Time  `gorm:"not null"`
	TurnInDate              *time.Time `gorm:"default:null"`
	DueDate                 time.Time  `gorm:"not null"`
}
