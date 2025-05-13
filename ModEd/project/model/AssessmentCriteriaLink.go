package model

import (
	"ModEd/core/validation"
	"encoding/json"
	"fmt"

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
	return fmt.Sprintf(
		"AssessmentCriteriaLink{ID: %v, AssessmentID: %v, AssessmentCriteriaID: %v}",
		p.ID, p.AssessmentId, p.AssessmentCriteriaId,
	)
}

func (p AssessmentCriteriaLink) Validate() error {
	v := validation.NewModelValidator()
	return v.ModelValidate(&p)
}

func (p AssessmentCriteriaLink) ToCSVRow() string {
	return fmt.Sprintf("%d,%d\n", p.ID, p.AssessmentCriteriaId)
}

func (p AssessmentCriteriaLink) FromCSV(raw string) error {
	_, err := fmt.Sscanf(raw, "%d,%s\n", &p.ID, &p.AssessmentCriteriaId)
	if err != nil {
		return err
	}
	return nil
}

func (p AssessmentCriteriaLink) ToJSON() string {
	return fmt.Sprintf(`{"id":%d,"assessment_id":%d,"assessment_criteria_id":%d}`, p.ID, p.AssessmentId, p.AssessmentCriteriaId)
}

func (p AssessmentCriteriaLink) FromJSON(raw string) error {
	var data struct {
		ID                   uint `json:"id"`
		AssessmentId         uint `json:"assessment_id"`
		AssessmentCriteriaId uint `json:"assessment_criteria_id"`
	}
	err := json.Unmarshal([]byte(raw), &data)
	if err != nil {
		return err
	}
	p.ID = data.ID
	p.AssessmentId = data.AssessmentId
	p.AssessmentCriteriaId = data.AssessmentCriteriaId
	return nil
}
