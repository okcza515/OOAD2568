package model

import (
	"fmt"
	"gorm.io/gorm"
)

type ScoreAssessmentAdvisor struct {
	gorm.Model
	AssessmentCriteriaLinkId uint    `gorm:"not null;index"`
	AdvisorId                uint    `gorm:"not null;index"`
	Score                    float64 `gorm:"not null"`
}

func (s *ScoreAssessmentAdvisor) GetID() uint {
	return s.ID
}

func (s *ScoreAssessmentAdvisor) ToString() string {
	return fmt.Sprintf("ID: %d, CriteriaLinkID: %d, AdvisorID: %d, Score: %.2f", s.ID, s.AssessmentCriteriaLinkId, s.AdvisorId, s.Score)
}

func (s *ScoreAssessmentAdvisor) Validate() error {
	if s.Score < 0 || s.Score > 100 {
		return fmt.Errorf("score must be between 0 and 100")
	}
	return nil
}

func (s *ScoreAssessmentAdvisor) ToCSVRow() string {
	return fmt.Sprintf("%d,%d,%d,%.2f", s.ID, s.AssessmentCriteriaLinkId, s.AdvisorId, s.Score)
}

func (s *ScoreAssessmentAdvisor) FromCSV(raw string) error {
	// Implement CSV parsing logic if needed
	return nil
}

func (s *ScoreAssessmentAdvisor) ToJSON() string {
	// Implement JSON serialization logic if needed
	return ""
}

func (s *ScoreAssessmentAdvisor) FromJSON(raw string) error {
	// Implement JSON deserialization logic if needed
	return nil
}
