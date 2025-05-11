// MEP-1013
package model

import (
	"ModEd/core"
	"fmt"
)

type InstrumentManagement struct {
	core.BaseModel
	BorrowUserID     uint       	  `gorm:"type:integer" json:"borrow_id" csv:"borrow_id"`
	BorrowInstrument BorrowInstrument `gorm:"foreignKey:BorrowUserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"borrow_instrument"`
	RoomID           uint       	  `gorm:"type:integer;not null;index" json:"room_id" csv:"room_id"`
	Room             Room       	  `gorm:"foreignKey:RoomID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"room"`
	InstrumentID     uint       	  `gorm:"type:integer;not null;index" json:"instrument_id" csv:"instrument_id"`
	Instrument       Instrument 	  `gorm:"foreignKey:InstrumentID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"instrument"`
	InstrumentLabel  string     	  `gorm:"type:text;not null" json:"instrument_label" csv:"instrument_label"`
}

func (im InstrumentManagement) ToString() string {
    
    return fmt.Sprintf("ID: %d | Instrument ID: %d | Label: %s | Room ID: %d | Room Name: %s | Borrow ID: %d",
        im.ID,
        im.InstrumentID,
        im.Instrument.InstrumentLabel,
        im.RoomID,
        im.Room.RoomName,
        im.BorrowUserID)
}