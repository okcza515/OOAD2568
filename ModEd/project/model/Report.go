package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Report struct {
	gorm.Model
	ReportId        uuid.UUID  `gorm:"type:text;primaryKey"`
	SeniorProjectId uuid.UUID  `gorm:"type:text;not null"`
	ReportType      ReportType `gorm:"type:varchar(50);not null"`
	Date            time.Time  `gorm:"not null"`

	ScoresReportAdvisor   []ScoreReportAdvisor   `gorm:"foreignKey:ReportId"`
	ScoresReportCommittee []ScoreReportCommittee `gorm:"foreignKey:ReportId"`
}
