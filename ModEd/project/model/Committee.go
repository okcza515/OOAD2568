package model

import (
	"gorm.io/gorm"
)

type Committee struct {
	gorm.Model
	InstructorId    uint `gorm:"type:text;not null;index"`
	SeniorProjectId uint
	SeniorProject   SeniorProject `gorm:"foreignKey:SeniorProjectId"`
}

func (p Committee) GetID() uint {
	return p.ID
}

func (p Committee) ToString() string {
	return ""
}

func (p Committee) Validate() error {
	return nil
}

func (p Committee) ToCSVRow() string {
	return ""
}

func (p Committee) FromCSV(raw string) error {
	return nil
}

func (p Committee) ToJSON() string {
	return ""
}

func (p Committee) FromJSON(raw string) error {
	return nil
}
