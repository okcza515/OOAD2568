package model

import (
	commonModel "ModEd/common/model"
	"ModEd/curriculum/model"
)

type AssignedCourse struct {
	AssignedCourseId string                   `gorm:"type:string;default:uuid_generate_v4();primaryKey"`
	Instructors      []commonModel.Instructor `gorm:"many2many:assigned_course_instructors;"`
	Course           model.Course
}
