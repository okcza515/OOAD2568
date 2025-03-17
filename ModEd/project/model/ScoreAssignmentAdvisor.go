package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ScoreAssignmentAdvisor struct {
	gorm.Model
	ScoreAssignmentAdvisorId uint `gorm:"primaryKey;default:gen_random_uuid()"`
	AssignmentId uuid.UUID    `gorm:"not null;index"`
	AdvisorId    uuid.UUID    `gorm:"not null;index"`
	Score        float64 `gorm:"not null"`
	Assignment   Assignment `gorm:"foreignKey:AssignmentId"`
}
