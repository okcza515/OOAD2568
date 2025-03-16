package model

import (
	"gorm.io/gorm"
)

type Committee struct {
	gorm.Model
	SeniorProjectID 	uint           `gorm:"not null;index"`
	InstructorID    	uint           `gorm:"not null;index"`
}