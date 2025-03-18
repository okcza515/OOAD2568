package model

import (
	"gorm.io/gorm"
)

type ScorePresentationAdvisor struct {
	gorm.Model
	PresentationId uint    `gorm:"not null;index"`
	AdvisorId      uint    `gorm:"not null;index"`
	Score          float64 `gorm:"not null"`
}
