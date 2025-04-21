package model

import (
	"fmt"
	"gorm.io/gorm"
)

type ScoreAssignmentCommittee struct {
	gorm.Model
	AssignmentId uint    `gorm:"not null;index"`
	CommitteeId  uint    `gorm:"not null;index"`
	Score        float64 `gorm:"not null"`
}

func (s *ScoreAssignmentCommittee) GetID() uint {
	return s.ID
}

func (s *ScoreAssignmentCommittee) ToString() string {
	return fmt.Sprintf("ID: %d, AssignmentID: %d, CommitteeID: %d, Score: %.2f", s.ID, s.AssignmentId, s.CommitteeId, s.Score)
}

func (s *ScoreAssignmentCommittee) Validate() error {
	if s.Score < 0 || s.Score > 100 {
		return fmt.Errorf("score must be between 0 and 100")
	}
	return nil
}

func (s *ScoreAssignmentCommittee) ToCSVRow() string {
	return fmt.Sprintf("%d,%d,%d,%.2f", s.ID, s.AssignmentId, s.CommitteeId, s.Score)
}

func (s *ScoreAssignmentCommittee) FromCSV(raw string) error {
	// Implement CSV parsing logic if needed
	return nil
}

func (s *ScoreAssignmentCommittee) ToJSON() string {
	// Implement JSON serialization logic if needed
	return ""
}

func (s *ScoreAssignmentCommittee) FromJSON(raw string) error {
	// Implement JSON deserialization logic if needed
	return nil
}
