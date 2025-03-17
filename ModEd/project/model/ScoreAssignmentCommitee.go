package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ScoreAssignmentCommittee struct {
	gorm.Model
	ScoreAssignmentCommitteeId uint `gorm:"primaryKey;default:gen_random_uuid()"`
	AssignmentId uuid.UUID    `gorm:"not null;index"`
	CommitteeId  uuid.UUID    `gorm:"not null;index"`
	Score        float64 `gorm:"not null"`
	Assignment   Assignment `gorm:"foreignKey:AssignmentId"`
}
