package model

import (
	common "ModEd/common/model"

	gorm "gorm.io/gorm"
)

type StudentAdvisor struct {
	gorm.Model
	InstructorId uint              `gorm:"not null" json:"instructor_id"`
	Instructor   common.Instructor `gorm:"foreignKey:InstructorId;references:ID" json:"-"`
	Students     []common.Student  `gorm:"many2many:student_advisor_students;"`
}
