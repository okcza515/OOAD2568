package model

import (
	"gorm.io/gorm"
)

type Advisor struct {
	gorm.Model
	IsPrimary       bool `gorm:"not null"`
	SeniorProjectId uint
	InstructorId    uint
	SeniorProject   SeniorProject `gorm:"foreignKey:SeniorProjectId"`
}
