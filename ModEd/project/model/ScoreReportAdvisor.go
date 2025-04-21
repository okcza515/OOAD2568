package model

import (
	"fmt"

	"gorm.io/gorm"
)

type ScoreReportAdvisor struct {
	gorm.Model
	ReportId             uint    `gorm:"not null"`
	AdvisorId            uint    `gorm:"not null"`
	Score                float64 `gorm:"not null"`
}

func (s *ScoreReportAdvisor) GetID() uint {
	return s.ID
}

func (s *ScoreReportAdvisor) ToString() string {
	return fmt.Sprintf("ID: %d, ReportID: %d, AdvisorID: %d, Score: %.2f", s.ID, s.ReportId, s.AdvisorId, s.Score)
}

func (s *ScoreReportAdvisor) Validate() error {
	if s.Score < 0 || s.Score > 100 {
		return fmt.Errorf("score must be between 0 and 100")
	}
	return nil
}

func (s *ScoreReportAdvisor) ToCSVRow() string {
	return fmt.Sprintf("%d,%d,%d,%.2f", s.ID, s.ReportId, s.AdvisorId, s.Score)
}

func (s *ScoreReportAdvisor) FromCSV(raw string) error {
	// Implement CSV parsing logic if needed
	return nil
}

func (s *ScoreReportAdvisor) ToJSON() string {
	// Implement JSON serialization logic if needed
	return ""
}

func (s *ScoreReportAdvisor) FromJSON(raw string) error {
	// Implement JSON deserialization logic if needed
	return nil
}
