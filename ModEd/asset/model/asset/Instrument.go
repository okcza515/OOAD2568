package asset

import (
	"gorm.io/gorm"
)

type Instrument struct {
	gorm.Model
	InstrumentLabel    string               `gorm:"not null"`
	InstrumentCode     string               `gorm:"type:text; not null; uniqueIndex"`
	Description        *string              `gorm:"type:text"`
	InstrumentStatus   InstrumentStatusEnum `gorm:"type:text;not null"`
	RoomID             string               `gorm:"type:text;not null"`
	Location           string               `gorm:"type:text;not null"`
	CategoryID         uint                 `gorm:"type:text;not null"`
	Cost               float64              `gorm:"type:real;not null"`
	InstrumentSerialID *string              `gorm:"type:text"`
	BudgetYear         int                  `gorm:"type:integer;not null"`
	BudgetSource       *string              `gorm:"type:text"`
	InstrumentBrand    *string              `gorm:"type:text"`
	InstrumentModel    *string              `gorm:"type:text"`
	Category           Category             `gorm:"foreignKey:CategoryID;constraint:OnUpdate:CASCADE;"`
}
