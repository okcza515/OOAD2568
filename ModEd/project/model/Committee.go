package model

import (
	"gorm.io/gorm"
)

type Committee struct {
	gorm.Model
	SeniorProjectId uint `gorm:"not null;index"`
	InstructorId    uint `gorm:"not null;index"`
}
