package model

import "gorm.io/gorm"

type ScoreReportAdvisor struct {
	gorm.Model
	ReportID  uint    `gorm:"not null"`
	AdvisorID uint    `gorm:"not null"`
	Score     float64 `gorm:"not null"`
}
