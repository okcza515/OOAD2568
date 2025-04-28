// MEP 1013
package model

import (
	"time"

	"ModEd/core"
)

type TimeTable struct {
	core.BaseModel
	StartDate   time.Time `gorm:"type:timestamp" json:"start_date" csv:"start_date"`
	EndDate     time.Time `gorm:"type:timestamp" json:"end_date" csv:"end_date"`
	RoomID      uint      `gorm:"type:integer" json:"room_id" csv:"room_id"`
	Room        Room      `gorm:"foreignKey:RoomID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	IsAvailable bool      `gorm:"type:boolean" json:"is_available" csv:"is_available"`
	BookingID   *uint     `gorm:"type:integer"`
	ScheduleID  *uint     `gorm:"type:integer"`
}
