package model

import (
	commonModel "ModEd/common/model"

	"gorm.io/gorm"
)

type StudentAdvisor struct {
	gorm.Model
	AdvisorId string
	Students  []commonModel.Student
}
