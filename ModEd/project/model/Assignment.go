package model

import (
	"time"

	"gorm.io/gorm"
)

type Assignment struct {
	gorm.Model
	Name            string     `gorm:"not null"`
	Description     string     `gorm:"not null"`
	SubmissionDate  *time.Time `gorm:"type:date"`
	DueDate         time.Time  `gorm:"type:date;not null"`
	SeniorProjectId uint
	SeniorProject   SeniorProject `gorm:"foreignKey:SeniorProjectId"`
}
