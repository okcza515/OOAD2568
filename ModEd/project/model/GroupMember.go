package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GroupMember struct {
	gorm.Model
	GroupMemberId   uuid.UUID `gorm:"type:uuid;primaryKey"`
	SeniorProjectID uint64    `gorm:"not null"`
	StudentID       uint64    `gorm:"not null"`
}
