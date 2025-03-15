package model

import (
	"gorm.io/gorm"
)

type Instrument struct {
	InstrumentID       string               `gorm:"type:text;primaryKey;not null;unique" json:"instrument_id" csv:"instrument_id"`
	InstrumentLabel    string               `gorm:"not null" json:"instrument_label" csv:"instrument_label"`
	InstrumentCode     string               `gorm:"not null;unique" json:"instrument_code" csv:"instrument_code"`
	InstrumentImageURL string               `gorm:"type:text" json:"instrument_image_url" csv:"instrument_image_url"`
	Description        string               `gorm:"type:text" json:"description" csv:"description"`
	InstrumentStatus   InstrumentStatusEnum `gorm:"type:text;not null" json:"instrument_status" csv:"instrument_status"`
	RoomID             string               `gorm:"type:text;not null" json:"room_id" csv:"room_id"`
	Location           string               `gorm:"type:text" json:"location" csv:"location"`
	CategoryID         *string              `gorm:"type:text" json:"category_id,omitempty" csv:"category_id"`
	Cost               float64              `gorm:"type:real" json:"cost" csv:"cost"`
	InstrumentSerialID string               `gorm:"type:text" json:"instrument_serial_id" csv:"instrument_serial_id"`
	AddYear            int                  `gorm:"type:integer" json:"add_year" csv:"add_year"`
	BudgetSource       string               `gorm:"type:text" json:"budget_source" csv:"budget_source"`
	Brand              string               `gorm:"type:text" json:"brand" csv:"brand"`
	Model              string               `gorm:"type:text" json:"model" csv:"model"`
	DeletedAt          gorm.DeletedAt       `gorm:"index" json:"deleted_at,omitempty" csv:"deleted_at"`
}
