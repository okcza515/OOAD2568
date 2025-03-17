package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Report struct {
	gorm.Model
	ReportId        uuid.UUID  `gorm:"type:uuid;primaryKey"`
	SeniorProjectId uuid.UUID  `gorm:"type:uuid;not null"`
	ReportType      ReportType `gorm:"type:varchar(255);not null"`
	Date            time.Time  `gorm:"not null"`

	ScoreReportAdvisors   []ScoreReportAdvisor   `gorm:"foreignKey:ReportId"`
	ScoreReportCommittees []ScoreReportCommittee `gorm:"foreignKey:ReportId"`
}
