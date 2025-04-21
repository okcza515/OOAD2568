package model

import (
	"ModEd/core"
	"time"
)

type IndependentStudy struct {
	core.BaseModel
	WILProjectId            uint       `gorm:"default:0"`
	WILProject              WILProject `json:"WILProject" gorm:"foreignKey:WILProjectId;references:ID"`
	IndependentStudyTopic   string     `gorm:"not null"`
	IndependentStudyContent string     `gorm:"not null"`
	EvaluationFromAdvisor   int        `gorm:"not null;default:0"`
	AssignmentDate          time.Time  `gorm:"not null"`
	TurnInDate              *time.Time `gorm:"default:null"`
	DueDate                 time.Time  `gorm:"not null"`
}
