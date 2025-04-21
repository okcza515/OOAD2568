package model

import (
	"time"

	commonModel "ModEd/common/model"

	curriculumModel "ModEd/curriculum/model"

	"gorm.io/gorm"
)

type Quiz struct {
	gorm.Model
	InstructorCode commonModel.Instructor
	FirstName      commonModel.Instructor
	LastName       commonModel.Instructor
	CourseId       curriculumModel.Course
	QuizId         uint
	Title          string
	Description    string
	Released       bool
	QuizStart      time.Time
	QuizEnd        time.Time
	Status         string
	Submission     []QuizSubmission
}

type QuizSubmission struct {
	gorm.Model
	StudentCode commonModel.Student
	FirstName   commonModel.Student
	LastName    commonModel.Student
	Email       commonModel.Student
	Answers     string
	Submitted   bool
	SubmittedAt time.Time
}
