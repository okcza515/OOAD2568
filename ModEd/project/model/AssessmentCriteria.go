package model

import (
	"gorm.io/gorm"
)

type AssessmentCriteria struct {
	gorm.Model
	AssessmentCriteriaId uint   `gorm:"primaryKey"`
	CriteriaName         string `gorm:"not null"`
}
