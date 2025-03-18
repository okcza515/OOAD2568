package model

import (
	"time"

	"gorm.io/gorm"
)

type Report struct {
	gorm.Model
	SeniorProjectId uint       `gorm:"type:text;not null"`
	ReportType      ReportType `gorm:"type:varchar(50);not null"`
	SubmissionDate  *time.Time `gorm:"type:date"`
	DueDate         time.Time  `gorm:"type:date;not null"`
}
