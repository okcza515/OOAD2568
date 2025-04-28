package model

import (
	"gorm.io/gorm"
)

type Advisor struct {
	gorm.Model
	IsPrimary       bool `gorm:"not null"`
	SeniorProjectId uint
	InstructorId    uint
	SeniorProject   SeniorProject `gorm:"foreignKey:SeniorProjectId"`
}

func (p Advisor) GetID() uint {
	return p.ID
}

func (p Advisor) ToString() string {
	return ""
}

func (p Advisor) Validate() error {
	return nil
}

func (p Advisor) ToCSVRow() string {
	return ""
}

func (p Advisor) FromCSV(raw string) error {
	return nil
}

func (p Advisor) ToJSON() string {
	return ""
}

func (p Advisor) FromJSON(raw string) error {
	return nil
}
