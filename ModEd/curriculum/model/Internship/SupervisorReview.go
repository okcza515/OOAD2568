package model

import (
	"gorm.io/gorm"
)

type SupervisorReview struct {
	gorm.Model
	InstructorScore int `gorm:"type:int"`
	MentorScore     int `gorm:"type:int"`
}
