// MEP-1013
package model

import (
	"ModEd/core"
	"fmt"
)

type InstrumentManagement struct {
	core.BaseModel
	BorrowUserID       uint             `gorm:"type:integer" json:"borrow_id" csv:"borrow_id"`
	BorrowInstrument   BorrowInstrument `gorm:"foreignKey:ID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"borrow_instrument"`
	RoomID             uint             `gorm:"type:integer;not null;index" json:"room_id" csv:"room_id"`
	Room               Room             `gorm:"foreignKey:ID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"room"`
	InstrumentSerialID uint             `gorm:"type:integer;not null;index" json:"instrument_id" csv:"instrument_id"`
	Instrument         Instrument       `gorm:"foreignKey:ID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"instrument"`
	InstrumentLabel    string           `gorm:"type:text;not null" json:"instrument_label" csv:"instrument_label"`
}

func (im InstrumentManagement) ToString() string {
    return fmt.Sprintf("ID: %d | Label: %s | Room ID: %d | Instrument ID: %d | Borrow ID: %d",
        im.ID,
        im.InstrumentLabel,
        im.RoomID,
        im.InstrumentSerialID,
        im.BorrowUserID)
}