package model

import "gorm.io/gorm"

type AssessmentCriteria struct {
	gorm.Model
	ID           uint64 `gorm:"primaryKey"`
	CriteriaName string `gorm:"not null"`
}
