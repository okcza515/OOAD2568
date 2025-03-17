package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Progress struct {
	gorm.Model
	ProgressId     uuid.UUID       `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	AssignmentId uuid.UUID   `gorm:"not null;index"`
	Name         string `gorm:"not null"`
	IsCompleted  bool   `gorm:"not null"`
	Assignment   Assignment `gorm:"foreignKey:AssignmentId"`
}
