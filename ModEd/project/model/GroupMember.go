package model

import "gorm.io/gorm"

type GroupMember struct {
	gorm.Model
	ID              uint64 `gorm:"primaryKey"`
	SeniorProjectID uint64 `gorm:"not null"`
	StudentID       uint64 `gorm:"not null"`
}
