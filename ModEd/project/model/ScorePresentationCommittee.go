package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ScorePresentationCommittee struct {
	gorm.Model
	PresentationId uuid.UUID     `gorm:"not null;index"`
	CommitteeId    uuid.UUID     `gorm:"not null;index"`
	Score          float64       `gorm:"not null"`
	Presentation   *Presentation `gorm:"foreignKey:PresentationId"`
}
