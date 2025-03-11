package model

import (
	"ModEd/common/model"
	"time"
)

type Class struct {
	Course      Course
	Section     int
	Schedule    time.Time
	StudentList []model.Student
	// TODO: Instructor is not defined in the model package, wait for Aj Bo
	// Instructor   model.Instructor
}
