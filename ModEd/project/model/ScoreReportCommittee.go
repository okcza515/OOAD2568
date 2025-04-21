package model

import (
	"fmt"
	"gorm.io/gorm"
)

type ScoreReportCommittee struct {
	gorm.Model
	ReportId    uint    `gorm:"not null"`
	CommitteeId uint    `gorm:"not null"`
	Score       float64 `gorm:"not null"`
}

func (s *ScoreReportCommittee) GetID() uint {
	return s.ID
}

func (s *ScoreReportCommittee) ToString() string {
	return fmt.Sprintf("ID: %d, ReportID: %d, CommitteeID: %d, Score: %.2f", s.ID, s.ReportId, s.CommitteeId, s.Score)
}

func (s *ScoreReportCommittee) Validate() error {
	if s.Score < 0 || s.Score > 100 {
		return fmt.Errorf("score must be between 0 and 100")
	}
	return nil
}

func (s *ScoreReportCommittee) ToCSVRow() string {
	return fmt.Sprintf("%d,%d,%d,%.2f", s.ID, s.ReportId, s.CommitteeId, s.Score)
}

func (s *ScoreReportCommittee) FromCSV(raw string) error {
	// Implement CSV parsing logic if needed
	return nil
}

func (s *ScoreReportCommittee) ToJSON() string {
	// Implement JSON serialization logic if needed
	return ""
}

func (s *ScoreReportCommittee) FromJSON(raw string) error {
	// Implement JSON deserialization logic if needed
	return nil
}
