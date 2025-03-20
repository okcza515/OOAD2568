package model

import (
	"gorm.io/gorm"
)

type ScoreAssignmentCommittee struct {
	gorm.Model
	ScoreAssignmentCommitteeId uint    `gorm:"primaryKey"`
	AssignmentId               uint    `gorm:"not null;index"`
	CommitteeId                uint    `gorm:"not null;index"`
	Score                      float64 `gorm:"not null"`
}
