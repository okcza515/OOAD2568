package model

import (
	"gorm.io/gorm"
)

type ScoreAssignmentAdvisor struct {
	gorm.Model
	ScoreAssignmentAdvisorId uint    `gorm:"primaryKey"`
	AssignmentId             uint    `gorm:"not null;index"`
	AdvisorId                uint    `gorm:"not null;index"`
	Score                    float64 `gorm:"not null"`
}
