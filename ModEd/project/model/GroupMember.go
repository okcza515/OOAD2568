package model

import (
	"gorm.io/gorm"
)

type GroupMember struct {
	gorm.Model
	StudentId       uint `gorm:"not null:index"`
	SeniorProjectId uint
	SeniorProject   SeniorProject `gorm:"foreignKey:SeniorProjectId"`
}
