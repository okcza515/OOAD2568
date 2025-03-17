package model

import "gorm.io/gorm"

type ScoreAssignmentAdvisor struct {
	gorm.Model
	AssignmentID uuid.UUID    `gorm:"not null;index"`
	AdvisorID    uuid.UUID    `gorm:"not null;index"`
	Score        float64 `gorm:"not null"`
}
