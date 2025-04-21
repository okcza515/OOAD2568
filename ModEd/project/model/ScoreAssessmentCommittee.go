package model

import (
	"fmt"
	"gorm.io/gorm"
)

type ScoreAssessmentCommittee struct {
	gorm.Model
	AssessmentCriteriaLinkId uint    `gorm:"not null;index"`
	CommitteeId              uint    `gorm:"not null;index"`
	Score                    float64 `gorm:"not null"`
}

func (s *ScoreAssessmentCommittee) GetID() uint {
	return s.ID
}

func (s *ScoreAssessmentCommittee) ToString() string {
	return fmt.Sprintf("ID: %d, CriteriaLinkID: %d, CommitteeID: %d, Score: %.2f", s.ID, s.AssessmentCriteriaLinkId, s.CommitteeId, s.Score)
}

func (s *ScoreAssessmentCommittee) Validate() error {
	if s.Score < 0 || s.Score > 100 {
		return fmt.Errorf("score must be between 0 and 100")
	}
	return nil
}

func (s *ScoreAssessmentCommittee) ToCSVRow() string {
	return fmt.Sprintf("%d,%d,%d,%.2f", s.ID, s.AssessmentCriteriaLinkId, s.CommitteeId, s.Score)
}

func (s *ScoreAssessmentCommittee) FromCSV(raw string) error {
	// Implement CSV parsing logic if needed
	return nil
}

func (s *ScoreAssessmentCommittee) ToJSON() string {
	// Implement JSON serialization logic if needed
	return ""
}

func (s *ScoreAssessmentCommittee) FromJSON(raw string) error {
	// Implement JSON deserialization logic if needed
	return nil
}
