package model

// MEP-1012 Asset

import (
	"ModEd/core"
	"time"
)

type BorrowInstrument struct {
	core.BaseModel
	StaffUserID        uint       `gorm:"not null"`
	BorrowUserID       uint       `gorm:"not null"`
	ExpectedReturnDate *time.Time `gorm:"type:timestamp"`
	BorrowDate         time.Time  `gorm:"type:timestamp;not null"`
	ReturnDate         time.Time  `gorm:"type:timestamp;not null"`
	InstrumentID       uint       `gorm:"not null"`
	Description        *string    `gorm:"type:text"`
	BorrowObjective    string     `gorm:"type:text;not null"`
	Instrument         Instrument `gorm:"foreignKey:InstrumentID;constraint:OnUpdate:CASCADE;"`
}
