// MEP-1013
package model

import (
	"ModEd/core"
	"fmt"
	"github.com/go-playground/validator/v10"
)


type SupplyManagement struct {
	core.BaseModel
	RoomID            uint   `gorm:"type:integer;not null;index" json:"room_id" csv:"room_id" validate:"required"`
	Room              Room   `gorm:"foreignKey:RoomID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"room" validate:"-"`
	SupplyID          uint   `gorm:"type:integer;not null;index" json:"supply_id" csv:"supply_id" validate:"required"`
	Supply            Supply `gorm:"foreignKey:SupplyID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"supply" validate:"-"`
	SupplyLabel       string `gorm:"type:text;not null" json:"supply_label" csv:"supply_label" validate:"-"`
	Quantity          int    `gorm:"not null" json:"quantity" csv:"quantity" validate:"-"`
}


func (sm SupplyManagement) ToString() string {
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
│                 Supply Management Info                      │
├─────────────────────────────────────────────────────────────┤
│ ID             : %-42s │
│ Supply ID      : %-42s │
│ Label          : %-42s │
│ Quantity       : %-42s │
│ Room ID        : %-42s │
│ Room Name      : %-42s │
└─────────────────────────────────────────────────────────────┘`,
		truncate(fmt.Sprintf("%d", sm.ID), 42),
		truncate(fmt.Sprintf("%d", sm.SupplyID), 42),
		truncate(sm.Supply.SupplyLabel, 42),
		truncate(fmt.Sprintf("%d", sm.Supply.Quantity), 42),
		truncate(fmt.Sprintf("%d", sm.RoomID), 42),
		truncate(sm.Room.RoomName, 42),
	)
}

func (sm SupplyManagement) Validate() error {
	validate := validator.New()
	if err := validate.Struct(sm); err != nil {
		return err
	}
	return nil
}