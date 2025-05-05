// MEP-1013
package model

import (
	"ModEd/core"
	"fmt"
)

type SupplyManagement struct {
	core.BaseModel
	RoomID            uint   `gorm:"type:integer;not null;index" json:"room_id" csv:"room_id"`
	Room              Room   `gorm:"foreignKey:ID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"room"`
	SupplyID          uint   `gorm:"type:integer;not null;index" json:"supply_id" csv:"supply_id"`
	Supply            Supply `gorm:"foreignKey:ID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"supply"`
	SupplyLabel       string `gorm:"type:text;not null" json:"supply_label" csv:"supply_label"`
	Quantity          int    `gorm:"not null" json:"quantity" csv:"quantity"`
}

func (im SupplyManagement) ToString() string {
    return fmt.Sprintf("ID: %d | Label: %s | Room ID: %d | Supply ID: %d | Quantity: %d",
        im.ID,
        im.SupplyLabel,
        im.RoomID,
        im.SupplyID,
        im.Quantity)
}