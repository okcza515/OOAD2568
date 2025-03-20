package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SupervisorReview struct {
	gorm.Model
	SupervisorReviewId int           `gorm:"primaryKey autoIncrement"`
	Student            InternStudent `gorm:"foreignKey:StudentID"`
	InstructorScore    int           `gorm:"type:int"`
	MentorScore        int           `gorm:"type:int"`
}
