package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ScoreReportCommittee struct {
	gorm.Model
	ReportId    uuid.UUID `gorm:"type:uuid;not null"`
	CommitteeId uuid.UUID `gorm:"type:uuid;not null"`
	Score       float64   `gorm:"not null"`
	Report      *Report   `gorm:"foreignKey:ReportId"`
}
