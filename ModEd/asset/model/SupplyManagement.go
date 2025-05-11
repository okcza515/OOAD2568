// MEP-1013
package model

import (
	"ModEd/core"
	"fmt"
)

type SupplyManagement struct {
	core.BaseModel
	RoomID            uint   `gorm:"type:integer;not null;index" json:"room_id" csv:"room_id"`
	Room              Room   `gorm:"foreignKey:RoomID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"room"`
	SupplyID          uint   `gorm:"type:integer;not null;index" json:"supply_id" csv:"supply_id"`
	Supply            Supply `gorm:"foreignKey:SupplyID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"supply"`
	SupplyLabel       string `gorm:"type:text;not null" json:"supply_label" csv:"supply_label"`
	Quantity          int    `gorm:"not null" json:"quantity" csv:"quantity"`
}

func (sm SupplyManagement) ToString() string {
	
    return fmt.Sprintf("ID: %d | Supply ID: %d | Label: %s | Quantity: %d | Room ID: %d | Room Name: %s",
        sm.ID,
        sm.SupplyID,
        sm.Supply.SupplyLabel,
		sm.Supply.Quantity,
        sm.RoomID,
        sm.Room.RoomName)
}