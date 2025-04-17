package model

import (
	"time"

	commonModel "ModEd/common/model"

	curriculumModel "ModEd/curriculum/model"

	"gorm.io/gorm"
)

type Assignment struct {
	gorm.Model
	FirstName    commonModel.Instructor
	LastName     commonModel.Instructor
	CourseId     curriculumModel.Course
	AssignmentId uint
	Title        string
	Description  string
	StartDate    time.Time
	DueDate      time.Time
	Submission   []AssignmentSubmission
}

type AssignmentSubmission struct {
	gorm.Model
	StudentCode commonModel.Student
	Content     string
	SubmittedAt time.Time
}
