package model

// MEP-1012 Asset

import (
	"ModEd/core"
	"fmt"
	"time"
)

type BorrowInstrument struct {
	core.BaseModel
	StaffUserID        uint       `gorm:"not null"`
	BorrowUserID       uint       `gorm:"not null"`
	ExpectedReturnDate time.Time  `gorm:"type:timestamp;not null"`
	BorrowDate         time.Time  `gorm:"type:timestamp;not null"`
	ReturnDate         *time.Time `gorm:"type:timestamp;`
	InstrumentID       uint       `gorm:"not null"`
	Description        *string    `gorm:"type:text"`
	BorrowObjective    string     `gorm:"type:text;not null"`
	Instrument         Instrument `gorm:"foreignKey:InstrumentID;constraint:OnUpdate:CASCADE;"`
}

func (inst BorrowInstrument) ToString() string {
	return fmt.Sprintf("[%v]\t%v\t%v", inst.Instrument, inst.BorrowDate, inst.ReturnDate)
}
