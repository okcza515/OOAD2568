package model

import "ModEd/common/model"

type Curriculum struct {
	Name        string
	StartYear   int
	EndYear     int
	Faculty     model.Faculty
	Department  model.Department
	ProgramType model.ProgramType
	CourseList  []Course
}
