package model

import (
	"gorm.io/gorm"
)

type Assessment struct {
	gorm.Model
	AssessmentId    uint `gorm:"primaryKey"`
	SeniorProjectId uint `gorm:"not null;index"`
}
