package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ScorePresentationAdvisor struct {
	gorm.Model
	PresentationId uuid.UUID     `gorm:"not null;index"`
	AdvisorID      uuid.UUID     `gorm:"not null;index"`
	Score          float64       `gorm:"not null"`
	Presentation   *Presentation `gorm:"foreignKey:PresentationId"`
}
