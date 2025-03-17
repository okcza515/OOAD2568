package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type WILProject struct {
	gorm.Model
	WilProjectId     uuid.UUID
	SeniorProjectId  uuid.UUID
	Company          uuid.UUID
	Mentor           string
	IndependentStudy []IndependentStudy
}
