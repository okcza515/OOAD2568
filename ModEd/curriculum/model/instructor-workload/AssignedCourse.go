package model

import (
	commonModel "ModEd/common/model"
	"ModEd/curriculum/model"
)

type AssignedCourse struct {
	AssignedCourseId int                      `gorm:"primaryKey;autoIncrement"`
	Instructors      []commonModel.Instructor `gorm:"many2many:assigned_course_instructors;"`
	Course           model.Course
}
