// MEP-1013
package model

import (
	"ModEd/core"
	"fmt"
)

type InstrumentManagement struct {
	core.BaseModel
	BorrowUserID     uint              `gorm:"type:integer" json:"borrow_id" csv:"borrow_id"`
	BorrowInstrument BorrowInstrument  `gorm:"foreignKey:BorrowUserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"borrow_instrument"`
	RoomID           uint              `gorm:"type:integer;not null;index" json:"room_id" csv:"room_id"`
	Room             Room              `gorm:"foreignKey:RoomID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"room"`
	InstrumentID     uint              `gorm:"type:integer;not null;index" json:"instrument_id" csv:"instrument_id"`
	Instrument       Instrument        `gorm:"foreignKey:InstrumentID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"instrument"`
	InstrumentLabel  string            `gorm:"type:text;not null" json:"instrument_label" csv:"instrument_label"`
}

func (im InstrumentManagement) ToString() string {
	// Local truncate function inside ToString
	truncate := func(s string, width int) string {
		r := []rune(s)
		if len(r) > width {
			return string(r[:width-3]) + "..."
		}
		// pad with spaces to the right
		for len(r) < width {
			r = append(r, ' ')
		}
		return string(r)
	}

	return fmt.Sprintf(`
┌─────────────────────────────────────────────────────────────┐
│                Instrument Management Info                   │
├─────────────────────────────────────────────────────────────┤
│ ID             : %-42s │
│ Instrument ID  : %-42s │
│ Label          : %-42s │
│ Room ID        : %-42s │
│ Room Name      : %-42s │
│ Borrow User ID : %-42s │
└─────────────────────────────────────────────────────────────┘`,
		truncate(fmt.Sprintf("%d", im.ID), 42),
		truncate(fmt.Sprintf("%d", im.InstrumentID), 42),
		truncate(im.Instrument.InstrumentLabel, 42),
		truncate(fmt.Sprintf("%d", im.RoomID), 42),
		truncate(im.Room.RoomName, 42),
		truncate(fmt.Sprintf("%d", im.BorrowUserID), 42),
	)
}



