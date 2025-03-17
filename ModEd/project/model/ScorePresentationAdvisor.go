package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ScorePresentationAdvisor struct {
	gorm.Model
	PresentationId uuid.UUID     `gorm:"type:text;not null;index"`
	AdvisorId      uuid.UUID     `gorm:"type:text;not null;index"`
	Score          float64       `gorm:"not null"`
	Presentation   *Presentation `gorm:"foreignKey:PresentationId"`
}
