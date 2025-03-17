package model

import (
	"ModEd/common/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Advisor struct {
	gorm.Model
	AdvisorID uuid.UUID
	Advisor   model.Instructor `json:"Advisor" gorm:"foreignKey:AdvisorID;references:instructor_id"`
	Student   []model.Student  `json:"Student" gorm:"foreignKey:StudentID;references:SID"`
}
