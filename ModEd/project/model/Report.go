package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Report struct {
	gorm.Model
	SeniorProjectID uuid.UUID  `gorm:"type:uuid;not null"`
	ReportType      ReportType `gorm:"type:varchar(255);not null"`
	Date            time.Time  `gorm:"not null"`

	ScoreReportAdvisors   []ScoreReportAdvisor   `gorm:"foreignKey:ReportID"`
	ScoreReportCommittees []ScoreReportCommittee `gorm:"foreignKey:ReportID"`
}
