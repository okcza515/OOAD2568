package model

import "ModEd/common/model"

type Curriculum struct {
	Name       string
	StartYear  int
	EndYear    int
	Program    model.ProgramType
	CourseList []Course
}
