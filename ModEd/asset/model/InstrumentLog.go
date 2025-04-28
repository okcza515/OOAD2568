package model

// MEP-1012 Asset

import (
	"ModEd/core"
	"fmt"
)

type InstrumentLog struct {
	core.BaseModel
	RefUserID        *uint
	StaffUserID      uint                    `gorm:"type:integer;not null"`
	Action           InstrumentLogActionEnum `gorm:"not null"`
	InstrumentID     uint                    `gorm:"type:integer;not null"`
	Description      string                  `gorm:"not null"`
	RefBorrowID      *uint                   `gorm:"type:integer"`
	BorrowInstrument BorrowInstrument        `gorm:"foreignKey:RefBorrowID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Instrument       Instrument              `gorm:"foreignKey:InstrumentID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (i InstrumentLog) ToString() string {
	return fmt.Sprintf("Log ID:\t[%v]\n"+
		"Instrument:\t[%v] %v\n"+
		"Action:\t\t%v\n"+
		"Description: \t\t%v\n", i.ID, i.Instrument.InstrumentCode, i.Instrument.InstrumentLabel, i.Action, i.Description)
}
