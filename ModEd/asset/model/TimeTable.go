// MEP 1013
package model

import (
	"time"

	"ModEd/core"
)

type TimeTable struct {
	core.BaseModel
	StartDate   time.Time   `gorm:"type:timestamp;not null" json:"start_date" csv:"start_date"`
	EndDate     time.Time   `gorm:"type:timestamp;not null" json:"end_date" csv:"end_date"`
	RoomID      uint        `gorm:"type:integer;not null" json:"room_id" csv:"room_id"`
	Room        Room        `gorm:"foreignKey:RoomID;references:ID" json:"room" csv:"room"`
	IsAvailable bool        `gorm:"type:boolean;not null;default:true" json:"is_available" csv:"is_available"`
	BookingType BookingType `gorm:"type:text" json:"booking_type" csv:"booking_type"`
}