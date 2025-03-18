package model

import (
	"gorm.io/gorm"
)

type ScoreAssessmentCommittee struct {
	gorm.Model
	ScoreAssessmentAdvisorId uint    `gorm:"primaryKey"`
	AssessmentId             uint    `gorm:"not null;index"`
	ComitteeId               uint    `gorm:"not null;index"`
	Score                    float64 `gorm:"not null"`
}
