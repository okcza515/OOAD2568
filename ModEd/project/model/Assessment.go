package model

import (
	"gorm.io/gorm"
)

type Assessment struct {
	gorm.Model
	SeniorProjectId uint `gorm:"not null;index"`

	AssessmentCriteriaLink []AssessmentCriteriaLink `gorm:"foreignKey:AssessmentId"`
}

func (p Assessment) GetID() uint {
	return p.ID
}

func (p Assessment) ToString() string {
	return ""
}

func (p Assessment) Validate() error {
	return nil
}

func (p Assessment) ToCSVRow() string {
	return ""
}

func (p Assessment) FromCSV(raw string) error {
	return nil
}

func (p Assessment) ToJSON() string {
	return ""
}

func (p Assessment) FromJSON(raw string) error {
	return nil
}
