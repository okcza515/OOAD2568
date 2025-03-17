package model

import "gorm.io/gorm"

type ScoreAssignmentCommittee struct {
	gorm.Model
	AssignmentID uuid.UUID    `gorm:"not null;index"`
	CommitteeID  uuid.UUID    `gorm:"not null;index"`
	Score        float64 `gorm:"not null"`
}
