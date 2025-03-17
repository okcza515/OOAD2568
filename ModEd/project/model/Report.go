package model

import (
	"time"

	"gorm.io/gorm"
)

type Report struct {
	gorm.Model
	SeniorProjectID uint       `gorm:"not null"`
	ReportType      ReportType `gorm:"type:varchar(255);not null"`
	Date            time.Time  `gorm:"not null"`

	ScoreReportAdvisors   []ScoreReportAdvisor   `gorm:"foreignKey:ReportID"`
	ScoreReportCommittees []ScoreReportCommittee `gorm:"foreignKey:ReportID"`
}
