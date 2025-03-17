package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ScoreReportCommittee struct {
	gorm.Model
	ScoreReportCommitteeId uuid.UUID `gorm:"type:text;primaryKey;default:gen_random_uuid()"`
	ReportId               uuid.UUID `gorm:"type:text;not null"`
	CommitteeId            uuid.UUID `gorm:"type:text;not null"`
	Score                  float64   `gorm:"not null"`

	Report    *Report    `gorm:"foreignKey:ReportId"`
	Committee *Committee `gorm:"foreignKey:CommitteeId"`
}
