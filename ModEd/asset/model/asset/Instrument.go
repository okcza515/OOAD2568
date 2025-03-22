package asset

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Instrument struct {
	InstrumentID       uuid.UUID            `gorm:"type:text;primaryKey" json:"instrument_id" csv:"instrument_id"`
	InstrumentLabel    string               `gorm:"not null" json:"instrument_label" csv:"instrument_label"`
	InstrumentCode     string               `gorm:"type:text; not null; uniqueIndex" json:"instrument_code" csv:"instrument_code"`
	Description        *string              `gorm:"type:text" json:"description" csv:"description"`
	InstrumentStatus   InstrumentStatusEnum `gorm:"type:text;not null" json:"instrument_status" csv:"instrument_status"`
	RoomID             string               `gorm:"type:text;not null" json:"room_id" csv:"room_id"`
	Location           string               `gorm:"type:text;not null" json:"location" csv:"location"`
	CategoryID         uuid.UUID            `gorm:"type:text;not null" json:"category_id" csv:"category_id"`
	Cost               float64              `gorm:"type:real;not null" json:"cost" csv:"cost"`
	InstrumentSerialID *string              `gorm:"type:text" json:"instrument_serial_id" csv:"instrument_serial_id"`
	BudgetYear         int                  `gorm:"type:integer;not null" json:"add_year" csv:"add_year"`
	BudgetSource       *string              `gorm:"type:text" json:"budget_source" csv:"budget_source"`
	InstrumentBrand    *string              `gorm:"type:text" json:"brand" csv:"brand"`
	InstrumentModel    *string              `gorm:"type:text" json:"model" csv:"model"`
	DeletedAt          gorm.DeletedAt       `gorm:"index" json:"deleted_at" csv:"deleted_at"`
	InstrumentLog      []InstrumentLog      `gorm:"foreignKey:InstrumentID;references:InstrumentID;constraint:OnUpdate:CASCADE;"`
}
