package model

import (
	"encoding/json"
	"fmt"

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
	return p.CriteriaName
}

func (p AssessmentCriteria) Validate() error {
	if p.CriteriaName == "" {
		return fmt.Errorf("criteria name cannot be empty")
	}
	return nil
}

func (p AssessmentCriteria) ToCSVRow() string {
	return fmt.Sprintf("%d,%s\n", p.ID, p.CriteriaName)
}

func (p AssessmentCriteria) FromCSV(raw string) error {
	_, err := fmt.Sscanf(raw, "%d,%s\n", &p.ID, &p.CriteriaName)
	if err != nil {
		return err
	}
	return nil
}

func (p AssessmentCriteria) ToJSON() string {
	return fmt.Sprintf(`{"id":%d,"criteria name":"%s"}`, p.ID, p.CriteriaName)
}

func (p AssessmentCriteria) FromJSON(raw string) error {
	var data struct {
		ID           uint   `json:"id"`
		CriteriaName string `json:"criteria_name"`
	}
	err := json.Unmarshal([]byte(raw), &data)
	if err != nil {
		return err
	}
	p.ID = data.ID
	p.CriteriaName = data.CriteriaName
	return nil
}
