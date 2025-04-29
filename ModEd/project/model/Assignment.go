package model

import (
	"time"

	"gorm.io/gorm"
)

type Assignment struct {
	gorm.Model
	Name            string     `gorm:"not null"`
	Description     string     `gorm:"not null"`
	SubmissionDate  *time.Time `gorm:"type:date"`
	DueDate         time.Time  `gorm:"type:date;not null"`
	SeniorProjectId uint
	SeniorProject   SeniorProject `gorm:"foreignKey:SeniorProjectId"`
}

func (p Assignment) GetID() uint {
	return p.ID
}

func (p Assignment) ToString() string {
	return ""
}

func (p Assignment) Validate() error {
	return nil
}

func (p Assignment) ToCSVRow() string {
	return ""
}

func (p Assignment) FromCSV(raw string) error {
	return nil
}

func (p Assignment) ToJSON() string {
	return ""
}

func (p Assignment) FromJSON(raw string) error {
	return nil
}
