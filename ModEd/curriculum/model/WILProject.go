package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type WILProject struct {
	gorm.Model
	ClassId          uint               `gorm:"not null"`
	SeniorProjectId  uuid.UUID          `gorm:"not null"`
	Company          uuid.UUID          `gorm:"not null"`
	Mentor           string             `gorm:"not null"`
	IndependentStudy []IndependentStudy `gorm:"foreignKey:IndependentStudyId"`
}
