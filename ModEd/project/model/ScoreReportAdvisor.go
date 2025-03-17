package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ScoreReportAdvisor struct {
	gorm.Model
	ReportID  uuid.UUID `gorm:"type:uuid;not null"`
	AdvisorID uuid.UUID `gorm:"type:uuid;not null"`
	Score     float64   `gorm:"not null"`
	Report    *Report   `gorm:"foreignKey:ReportId"`
}

