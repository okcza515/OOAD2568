package model

import (
	"gorm.io/gorm"
)

type Advisor struct {
	gorm.Model
	SeniorProjectId uint `gorm:"not null;index"`
	InstructorId    uint `gorm:"not null;index"`
	IsPrimary       bool `gorm:"not null"`
}
