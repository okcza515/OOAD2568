package model

import (
	"gorm.io/gorm"
)

type AssessmentCriteria struct {
	gorm.Model
	CriteriaName string `gorm:"not null"`
}

func (p AssessmentCriteria) GetID() uint {
	return p.ID
}

func (p AssessmentCriteria) ToString() string {
	return ""
}

func (p AssessmentCriteria) Validate() error {
	return nil
}

func (p AssessmentCriteria) ToCSVRow() string {
	return ""
}

func (p AssessmentCriteria) FromCSV(raw string) error {
	return nil
}

func (p AssessmentCriteria) ToJSON() string {
	return ""
}

func (p AssessmentCriteria) FromJSON(raw string) error {
	return nil
}
