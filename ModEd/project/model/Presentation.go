package model

import (
	"time"

	"gorm.io/gorm"
)

type Presentation struct {
	gorm.Model
	PresentationType PresentationType `gorm:"type:varchar(50);not null"`
	Date             time.Time        `gorm:"type:date;not null"`
	SeniorProjectId  uint
	SeniorProject    SeniorProject `gorm:"foreignKey:SeniorProjectId"`
}

func (p Presentation) GetID() uint {
	return p.ID
}

func (p Presentation) ToString() string {
	return ""
}

func (p Presentation) Validate() error {
	return nil
}

func (p Presentation) ToCSVRow() string {
	return ""
}

func (p Presentation) FromCSV(raw string) error {
	return nil
}

func (p Presentation) ToJSON() string {
	return ""
}

func (p Presentation) FromJSON(raw string) error {
	return nil
}
