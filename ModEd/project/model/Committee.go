package model

import (
	"gorm.io/gorm"
)

type Committee struct {
	gorm.Model
	CommitteeId     uint `gorm:"primaryKey"`
	SeniorProjectId uint `gorm:"not null;index"`
	InstructorId    uint `gorm:"not null;index"`
}
