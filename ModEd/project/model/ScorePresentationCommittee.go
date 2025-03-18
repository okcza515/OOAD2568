package model

import (
	"gorm.io/gorm"
)

type ScorePresentationCommittee struct {
	gorm.Model
	PresentationId uint    `gorm:"not null;index"`
	CommitteeId    uint    `gorm:"not null;index"`
	Score          float64 `gorm:"not null"`
}
