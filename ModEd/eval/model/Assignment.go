package model

import (
	"time"

	"gorm.io/gorm"
)

type Assignment struct {
	gorm.Model
	AssignmentId uint
	Title        string
	Description  string
	Released     bool
	StartDate    time.Time
	DueDate      time.Time
	Status       string
	Submission   []AssignmentSubmission
}

type AssignmentSubmission struct {
	gorm.Model
	Answers     string
	Submitted   bool
	SubmittedAt time.Time
}
