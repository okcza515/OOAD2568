package model

import (
	"gorm.io/gorm"
)

type Committee struct {
	gorm.Model
	InstructorId    int `gorm:"type:text;not null;index"`
	SeniorProjectId uint
	SeniorProject   SeniorProject `gorm:"foreignKey:SeniorProjectId"`
}
