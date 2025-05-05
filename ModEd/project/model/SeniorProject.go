package model

import (
	"gorm.io/gorm"
)

type SeniorProject struct {
	gorm.Model
	GroupName     string         `gorm:"not null"`
	Members       []GroupMember  `gorm:"foreignKey:SeniorProjectId"`
	Advisors      []Advisor      `gorm:"foreignKey:SeniorProjectId"`
	Committees    []Committee    `gorm:"foreignKey:SeniorProjectId"`
	Assignments   []Assignment   `gorm:"foreignKey:SeniorProjectId"`
	Presentations []Presentation `gorm:"foreignKey:SeniorProjectId"`
	Reports       []Report       `gorm:"foreignKey:SeniorProjectId"`
	Assessments   []Assessment   `gorm:"foreignKey:SeniorProjectId"`
}

func (p SeniorProject) GetID() uint {
	return p.ID
}

func (p SeniorProject) ToString() string {
	return ""
}

func (p SeniorProject) Validate() error {
	return nil
}

func (p SeniorProject) ToCSVRow() string {
	return ""
}

func (p SeniorProject) FromCSV(raw string) error {
	return nil
}

func (p SeniorProject) ToJSON() string {
	return ""
}

func (p SeniorProject) FromJSON(raw string) error {
	return nil
}
