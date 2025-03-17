package model

import (
	"ModEd/common/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Advisor struct {
	gorm.Model
	model.Instructor
		AdvisorID uuid.UUID
}