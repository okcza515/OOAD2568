package model

import (
	"gorm.io/gorm"
)

type AssessmentCriteriaLink struct {
	gorm.Model
	AssessmentId         uint `gorm:"primaryKey"`
	AssessmentCriteriaId uint `gorm:"primaryKey"`

	ScoreAssessmentAdvisor   ScoreAssessmentAdvisor     `gorm:"foreignKey:AssessmentCriteriaLinkId"`
	ScoreAssessmentCommittee []ScoreAssessmentCommittee `gorm:"foreignKey:AssessmentCriteriaLinkId"`
}

func (p AssessmentCriteriaLink) GetID() uint {
	return p.ID
}

func (p AssessmentCriteriaLink) ToString() string {
	return ""
}

func (p AssessmentCriteriaLink) Validate() error {
	return nil
}

func (p AssessmentCriteriaLink) ToCSVRow() string {
	return ""
}

func (p AssessmentCriteriaLink) FromCSV(raw string) error {
	return nil
}

func (p AssessmentCriteriaLink) ToJSON() string {
	return ""
}

func (p AssessmentCriteriaLink) FromJSON(raw string) error {
	return nil
}
