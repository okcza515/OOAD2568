package model

import (
	"gorm.io/gorm"
)

type GroupMember struct {
	gorm.Model
	SeniorProjectId uint `gorm:"not null;index"`
	StudentId       uint `gorm:"not null:index"`
}
