package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ScorePresentationAdvisor struct {
	gorm.Model
	PresentationId uuid.UUID `gorm:"type:text;not null;index"`
	AdvisorID      uuid.UUID `gorm:"type:text;not null;index"`
	Score          float64   `gorm:"not null"`
}
