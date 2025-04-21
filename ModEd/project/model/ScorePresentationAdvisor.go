package model

import (
	"fmt"
	"gorm.io/gorm"
)

type ScorePresentationAdvisor struct {
	gorm.Model
	PresentationId uint    `gorm:"not null;index"`
	AdvisorId      uint    `gorm:"not null;index"`
	Score          float64 `gorm:"not null"`
}

func (s *ScorePresentationAdvisor) GetID() uint {
	return s.ID
}

func (s *ScorePresentationAdvisor) ToString() string {
	return fmt.Sprintf("ID: %d, PresentationID: %d, AdvisorID: %d, Score: %.2f", s.ID, s.PresentationId, s.AdvisorId, s.Score)
}

func (s *ScorePresentationAdvisor) Validate() error {
	if s.Score < 0 || s.Score > 100 {
		return fmt.Errorf("score must be between 0 and 100")
	}
	return nil
}

func (s *ScorePresentationAdvisor) ToCSVRow() string {
	return fmt.Sprintf("%d,%d,%d,%.2f", s.ID, s.PresentationId, s.AdvisorId, s.Score)
}

func (s *ScorePresentationAdvisor) FromCSV(raw string) error {
	// Implement CSV parsing logic if needed
	return nil
}

func (s *ScorePresentationAdvisor) ToJSON() string {
	// Implement JSON serialization logic if needed
	return ""
}

func (s *ScorePresentationAdvisor) FromJSON(raw string) error {
	// Implement JSON deserialization logic if needed
	return nil
}
