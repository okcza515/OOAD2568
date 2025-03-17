package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ScorePresentationCommittee struct {
	gorm.Model
	PresentationId uuid.UUID     `gorm:"type:text;not null;index"`
	CommitteeId    uuid.UUID     `gorm:"type:text;not null;index"`
	Score          float64       `gorm:"not null"`
	Presentation   *Presentation `gorm:"foreignKey:PresentationId"`
	Committee      *Committee    `gorm:"foreignKey:CommitteeId"`
}
