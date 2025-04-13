// MEP 1013
package spacemanagement

import (
	"time"

	"gorm.io/gorm"
)

type TimeTable struct {
	gorm.Model
	StartDate   time.Time `gorm:"type:timestamp"`
	EndDate     time.Time `gorm:"type:timestamp"`
	RoomID      uint      `gorm:"type:integer"`
	IsAvailable bool      `gorm:"type:boolean"`
	Room        Room      `gorm:"foreignKey:RoomID;references:RoomID"`
	BookingID   *uint     `gorm:"type:integer"`
	ScheduleID  *uint     `gorm:"type:integer"`
}
