package model

import (
	"gorm.io/gorm"
)

type WILProject struct {
	gorm.Model
	ClassId         uint   `gorm:"not null"`
	SeniorProjectId uint   `gorm:"not null"`
	Company         uint   `gorm:"not null"`
	Mentor          string `gorm:"not null"`
}
