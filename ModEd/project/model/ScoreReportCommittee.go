package model

import (
	"gorm.io/gorm"
)

type ScoreReportCommittee struct {
	gorm.Model
	ScoreReportCommitteeId uint    `gorm:"primaryKey"`
	ReportId               uint    `gorm:"not null"`
	CommitteeId            uint    `gorm:"not null"`
	Score                  float64 `gorm:"not null"`
}
