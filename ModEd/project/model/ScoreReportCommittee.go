package model

import "gorm.io/gorm"

type ScoreReportCommittee struct {
	gorm.Model
	ReportID    uint    `gorm:"not null"`
	CommitteeID uint    `gorm:"not null"`
	Score       float64 `gorm:"not null"`
}
