package model

import "gorm.io/gorm"

type Progress struct {
	gorm.Model
	AssignmentID uuid.UUID   `gorm:"not null;index"`
	Name         string `gorm:"not null"`
	IsCompleted  bool   `gorm:"not null"`
}
