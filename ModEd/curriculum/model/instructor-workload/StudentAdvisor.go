package model

import (
	commonModel "ModEd/common/model"

	"gorm.io/gorm"
)

type StudentAdvisor struct {
	gorm.Model
	AdvisorId uint                   `gorm:"not null;index"`
	Advisor   commonModel.Instructor `gorm:"foreignKey:AdvisorID"`
	Students  []commonModel.Student  `gorm:"many2many:student_advisor_students;"`
}
