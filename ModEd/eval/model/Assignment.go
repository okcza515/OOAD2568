package model

import (
	"time"

	"ModEd/common/model"

	"gorm.io/gorm"
)

type Assignment struct {
	gorm.Model
	InstructorId model.Instructor
	Title        string
	Description  string
	StartDate    time.Time
	DueDate      time.Time
	Submission   []AssignmentSubmission
}

type AssignmentSubmission struct {
	gorm.Model
	StudentCode model.Student
	Content     string
	SubmittedAt string
}
