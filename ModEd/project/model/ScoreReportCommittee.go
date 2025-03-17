package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ScoreReportCommittee struct {
	gorm.Model
	ReportID    uuid.UUID    `gorm:"type:uuid;not null"`
	CommitteeID uuid.UUID    `gorm:"type:uuid;not null"`
	Score       float64 `gorm:"not null"`
}
