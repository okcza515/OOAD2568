package model

import (
	commonModel "ModEd/common/model"

	"gorm.io/gorm"
)

type StudentAdvisor struct {
	gorm.Model
	commonModel.Instructor
	Students []commonModel.Student `gorm:"many2many:student_advisor_students;"`
}
