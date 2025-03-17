package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SupervisorReview struct {
	gorm.Model
	ReviewId        uuid.UUID
	Student         InternStudent `gorm:"foreignKey:StudentID"`
	InstructorScore int
	MentorScore     int
}
