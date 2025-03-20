package model

import (
	"time"

	"gorm.io/gorm"
)

type Assignment struct {
	gorm.Model
	SeniorProjectId int        `gorm:"type:text;not null;index"`
	Name            string     `gorm:"not null"`
	Description     string     `gorm:"not null"`
	SubmissionDate  *time.Time `gorm:"type:date"`
	DueDate         time.Time  `gorm:"type:date;not null"`
}
