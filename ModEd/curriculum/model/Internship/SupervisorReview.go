package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SupervisorReview struct {
	gorm.Model
	ReviewId        uuid.UUID     `gorm:"primaryKey"`
	Student         InternStudent `gorm:"foreignKey:StudentID"`
	InstructorScore int           `gorm:"type:int"`
	MentorScore     int           `gorm:"type:int"`
}
