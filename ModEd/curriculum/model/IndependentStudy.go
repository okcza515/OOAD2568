package model

import (
	"time"
	"ModEd/core"
)

type IndependentStudy struct {
	core.BaseModel
	IndependentStudyTopic   string     `gorm:"not null"`
	IndependentStudyContent string     `gorm:"not null"`
	EvaluationFromAdvisor   int        `gorm:"not null;default:0"`
	AssignmentDate          time.Time  `gorm:"not null"`
	TurnInDate              *time.Time `gorm:"default:null"`
	DueDate                 time.Time  `gorm:"not null"`
}
