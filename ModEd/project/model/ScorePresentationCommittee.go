package model

import "gorm.io/gorm"

type ScorePresentationCommittee struct {
	gorm.Model
	PresentationID uint    `gorm:"not null;index"`
	CommitteeID    uint    `gorm:"not null;index"`
	Score          float64 `gorm:"not null"`
}
