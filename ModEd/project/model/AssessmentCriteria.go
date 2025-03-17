package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AssessmentCriteria struct {
	gorm.Model
	AssessmentCriteriaId uuid.UUID `gorm:"type:uuid;primaryKey"`
	CriteriaName         string    `gorm:"not null"`
}
