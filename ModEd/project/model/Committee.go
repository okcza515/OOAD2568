package model

import (
	"gorm.io/gorm"
)

type Committee struct {
	gorm.Model
	SeniorProjectId int `gorm:"type:text;not null;index"`
	InstructorId    int `gorm:"type:text;not null;index"`
}
