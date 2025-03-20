package model

import (
	"gorm.io/gorm"
)

type Advisor struct {
	gorm.Model
	SeniorProjectId int  `gorm:"type:text;not null;index"`
	InstructorId    int  `gorm:"type:text;not null;index"`
	IsPrimary       bool `gorm:"not null"`
}
