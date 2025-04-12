package model

import (
	"time"

	"ModEd/common/model"

	"ModEd/curriculum/model"

	"gorm.io/gorm"
)

type Assignment struct {
	gorm.Model
	FirstName    model.Instructor
	LastName     model.Instructor
	CourseId     model.Course
	AssignmentId uint
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
