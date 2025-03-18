package model

import (
	"gorm.io/gorm"
)

type SeniorProject struct {
	gorm.Model
	SeniorProjectId uint   `gorm:"primaryKey"`
	GroupName       string `gorm:"not null"`
}
