package model

import (
	"gorm.io/gorm"
)

type Progress struct {
	gorm.Model
	ProgressId   uint   `gorm:"primaryKey"`
	AssignmentId uint   `gorm:"not null;index"`
	Name         string `gorm:"not null"`
	IsCompleted  bool   `gorm:"not null"`
}
