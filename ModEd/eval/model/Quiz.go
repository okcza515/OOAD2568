package model

import (
	"time"

	"ModEd/common/model"

	"ModEd/curriculum/model"

	"gorm.io/gorm"
)

type Quiz struct {
	gorm.Model
	ID           uint
	Title        string
	Description  string
	CourseId     model.Course
	InstructorId model.Instructor
	Released     bool
	QuizStart    time.Time
	QuizEnd      time.Time
	Status       string
	Submission   []QuizSubmission
}

type QuizSubmission struct {
	gorm.Model
	SID        model.Student
	FirstName  model.Student
	LastName   model.Student
	Email      model.Student
	Answers    string
	Submitted  bool
	SubmitTime time.Time
}
