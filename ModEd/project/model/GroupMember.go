package model

import (
	"gorm.io/gorm"
)

type GroupMember struct {
	gorm.Model
	StudentId       uint `gorm:"not null:index"`
	SeniorProjectId uint
	SeniorProject   SeniorProject `gorm:"foreignKey:SeniorProjectId"`
}

func (p GroupMember) GetID() uint {
	return p.ID
}

func (p GroupMember) ToString() string {
	return ""
}

func (p GroupMember) Validate() error {
	return nil
}

func (p GroupMember) ToCSVRow() string {
	return ""
}

func (p GroupMember) FromCSV(raw string) error {
	return nil
}

func (p GroupMember) ToJSON() string {
	return ""
}

func (p GroupMember) FromJSON(raw string) error {
	return nil
}
