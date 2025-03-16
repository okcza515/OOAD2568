package model

import (
	"github.com/google/uuid"
)

type Department struct {
	Name   		string
	Parent 		Faculty
	Students 	[]Student
	Instructors []Instructor
	CourseId 	[]uuid.UUID 		//seems like I have to use UUID as it would create circular dependency
	Budget 		int
}