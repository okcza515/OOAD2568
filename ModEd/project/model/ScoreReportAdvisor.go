package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ScoreReportAdvisor struct {
	gorm.Model
	ReportId  uuid.UUID `gorm:"type:text;not null"`
	AdvisorId uuid.UUID `gorm:"type:text;not null"`
	Score     float64   `gorm:"not null"`
	Report    *Report   `gorm:"foreignKey:ReportId"`
}

