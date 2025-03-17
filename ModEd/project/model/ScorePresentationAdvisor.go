package model

import "gorm.io/gorm"

type ScorePresentationAdvisor struct {
	gorm.Model
	PresentationID uint    `gorm:"not null;index"`
	AdvisorID      uint    `gorm:"not null;index"`
	Score          float64 `gorm:"not null"`
}
