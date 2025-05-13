package model

import (
	"encoding/json"
	"fmt"

	"gorm.io/gorm"
)

type Assessment struct {
	gorm.Model
	SeniorProjectId        uint                     `gorm:"not null;index"`
	SeniorProject          SeniorProject            `gorm:"foreignKey:SeniorProjectId"`
	AssessmentCriteriaLink []AssessmentCriteriaLink `gorm:"foreignKey:AssessmentId"`
}

func (p Assessment) GetID() uint {
	return p.ID
}

func (p Assessment) ToString() string {
	return fmt.Sprintf("Assessment ID: %v, SeniorProject ID: %v", p.ID, p.SeniorProjectId)
}

func (p Assessment) Validate() error {
	if p.SeniorProjectId == 0 {
		return fmt.Errorf("senior project id is required")
	}
	return nil
}

func (p Assessment) ToCSVRow() string {
	return fmt.Sprintf("%d,%d\n", p.ID, p.SeniorProjectId)
}

func (p Assessment) FromCSV(raw string) error {
	_, err := fmt.Sscanf(raw, "%d,%d\n", &p.ID, &p.SeniorProjectId)
	if err != nil {
		return err
	}
	return nil
}

func (p Assessment) ToJSON() string {
	return fmt.Sprintf(`{"id":%d,"senior_project_id":%d}`, p.ID, p.SeniorProjectId)
}

func (p Assessment) FromJSON(raw string) error {
	var data struct {
		ID              uint `json:"id"`
		SeniorProjectID uint `json:"senior_project_id"`
	}
	err := json.Unmarshal([]byte(raw), &data)
	if err != nil {
		return err
	}
	p.ID = data.ID
	p.SeniorProjectId = data.SeniorProjectID
	return nil
}
