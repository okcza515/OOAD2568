package model

import (
	commonModel "ModEd/common/model"
	"ModEd/curriculum/model"

	"gorm.io/gorm"
)

type AssignedCourse struct {
	gorm.Model
	ActionTracker
	CourseId    uint                     `gorm:"index"`
	Course      model.Course             `gorm:"foreignKey:CourseId;references:CourseId"`
	Instructors []commonModel.Instructor `gorm:"many2many:assigned_course_instructors;"`
}
