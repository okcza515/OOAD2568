package model

import (
	"fmt"
	"gorm.io/gorm"
)

type ScorePresentationCommittee struct {
	gorm.Model
	PresentationId uint
	Presentation   Presentation `gorm:"foreignKey:PresentationId"`
	CommitteeId    uint
	Committee      Committee `gorm:"foreignKey:CommitteeId"`
	Score          float64
}

func (s *ScorePresentationCommittee) GetID() uint {
	return s.ID
}

func (s *ScorePresentationCommittee) ToString() string {
	return fmt.Sprintf("ID: %d, PresentationID: %d, CommitteeID: %d, Score: %.2f", s.ID, s.PresentationId, s.CommitteeId, s.Score)
}

func (s *ScorePresentationCommittee) Validate() error {
	if s.Score < 0 || s.Score > 100 {
		return fmt.Errorf("score must be between 0 and 100")
	}
	return nil
}

func (s *ScorePresentationCommittee) ToCSVRow() string {
	return fmt.Sprintf("%d,%d,%d,%.2f", s.ID, s.PresentationId, s.CommitteeId, s.Score)
}

func (s *ScorePresentationCommittee) FromCSV(raw string) error {
	// Implement CSV parsing logic if needed
	return nil
}

func (s *ScorePresentationCommittee) ToJSON() string {
	// Implement JSON serialization logic if needed
	return ""
}

func (s *ScorePresentationCommittee) FromJSON(raw string) error {
	// Implement JSON deserialization logic if needed
	return nil
}
