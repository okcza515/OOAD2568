package model

import (
	commonModel "ModEd/common/model"

	"gorm.io/gorm"
)

type AttendeeAdapter interface {
	GetCode() string
	GetType() string
}

type MeetingAttendee struct {
	gorm.Model
	MeetingID    uint
	AttendeeCode string
	AttendeeType string
}

type InstructorAdapter struct {
	Instructor commonModel.Instructor
}

func (a InstructorAdapter) GetCode() string {
	return a.Instructor.InstructorCode
}

func (a InstructorAdapter) GetType() string {
	return "instructor"
}

type StudentAdapter struct {
	Student commonModel.Student
}

func (a StudentAdapter) GetCode() string {
	return a.Student.StudentCode
}

func (a StudentAdapter) GetType() string {
	return "student"
}
