package model

import (
	"fmt"

	"gorm.io/gorm"
)

type ScoreAssignmentAdvisor struct {
	gorm.Model
	AssignmentId uint       `gorm:"not null;index"`
	AdvisorId    uint       `gorm:"not null;index"`
	Score        float64    `gorm:"not null"`
	Assignment   Assignment `gorm:"foreignKey:AssignmentId"`
	Advisor      Advisor    `gorm:"foreignKey:AdvisorId"`
}

func (s *ScoreAssignmentAdvisor) GetID() uint {
	return s.ID
}

func (s *ScoreAssignmentAdvisor) ToString() string {
	return fmt.Sprintf("ID: %d, AssignmentID: %d, AdvisorID: %d, Score: %.2f", s.ID, s.AssignmentId, s.AdvisorId, s.Score)
}

func (s *ScoreAssignmentAdvisor) Validate() error {
	if s.Score < 0 || s.Score > 100 {
		return fmt.Errorf("score must be between 0 and 100")
	}
	return nil
}

func (s *ScoreAssignmentAdvisor) ToCSVRow() string {
	return fmt.Sprintf("%d,%d,%d,%.2f", s.ID, s.AssignmentId, s.AdvisorId, s.Score)
}

func (s *ScoreAssignmentAdvisor) FromCSV(raw string) error {
	// Implement CSV parsing logic if needed
	return nil
}

func (s *ScoreAssignmentAdvisor) ToJSON() string {
	// Implement JSON serialization logic if needed
	return ""
}

func (s *ScoreAssignmentAdvisor) FromJSON(raw string) error {
	// Implement JSON deserialization logic if needed
	return nil
}
