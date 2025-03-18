package model

import (
	"gorm.io/gorm"
)

type ScoreReportAdvisor struct {
	gorm.Model
	ScoreReportAdvisorId uint    `gorm:"primaryKey"`
	ReportId             uint    `gorm:"not null"`
	AdvisorId            uint    `gorm:"not null"`
	Score                float64 `gorm:"not null"`
}
