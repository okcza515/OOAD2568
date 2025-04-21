package model

import (
	"time"

	commonModel "ModEd/common/model"

	curriculumModel "ModEd/curriculum/model"

	"gorm.io/gorm"
)

type Quiz struct {
	gorm.Model
	ID           uint
	Title        string
	Description  string
	CourseId     curriculumModel.Course
	InstructorId commonModel.Instructor
	Released     bool
	QuizStart    time.Time
	QuizEnd      time.Time
	Status       string
	Submission   []QuizSubmission
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
