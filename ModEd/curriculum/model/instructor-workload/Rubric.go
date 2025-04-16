package model

import "gorm.io/gorm"

type Rubric struct {
	gorm.Model
	AssignmentId  uint    `gorm:"not null;index"`
	CriterionName string  `gorm:"type:varchar(100);not null"`
	Description   string  `gorm:"type:text"`
	Weight        float64 `gorm:"not null"`
	MaxScore      float64 `gorm:"default:10"`
}
