package model

import (
	"time"

	"gorm.io/gorm"
)

type Assignment struct {
	gorm.Model
	SeniorProjectID 		uint      `gorm:"not null;index"`
	AssignmentName  		string    `gorm:"not null;size:255"`
	AssignmentDescription	string    `gorm:"not null;size:255"`
	SubmissionDate			time.Time `gorm:""`
	DueDate         		time.Time `gorm:"not null"`
}