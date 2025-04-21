package model

import (
	"gorm.io/gorm"
)

type ScoreAssessmentCommittee struct {
	gorm.Model
	ScoreAssessmentComitteeId uint    `gorm:"primaryKey"`
	AssessmentCriteriaLinkId  uint    `gorm:"not null;index"`
	ComitteeId                uint    `gorm:"not null;index"`
	Score                     float64 `gorm:"not null"`
}
