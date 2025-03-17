package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Report struct {
	gorm.Model
	ReportId        uuid.UUID  `gorm:"type:text;primaryKey;default:gen_random_uuid()"`
	SeniorProjectId uuid.UUID  `gorm:"type:text;not null"`
	ReportType      ReportType `gorm:"type:varchar(50);not null"`
	SubmissionDate  *time.Time `gorm:"type:date"`
	DueDate         time.Time  `gorm:"type:date;not null"`

	ScoresReportAdvisor   []ScoreReportAdvisor   `gorm:"foreignKey:ReportId"`
	ScoresReportCommittee []ScoreReportCommittee `gorm:"foreignKey:ReportId"`
	SeniorProject         *SeniorProject         `gorm:"foreignKey:SeniorProjectId"`
}
