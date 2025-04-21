package model

import (
	"gorm.io/gorm"
)

type ScoreAssessmentAdvisor struct {
	gorm.Model
	ScoreAssessmentAdvisorId uint    `gorm:"primaryKey"`
	AssessmentCriteriaLinkId uint    `gorm:"not null;index"`
	AdvisorId                uint    `gorm:"not null;index"`
	Score                    float64 `gorm:"not null"`
}
