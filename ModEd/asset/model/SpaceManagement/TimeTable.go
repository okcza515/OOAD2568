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
	Room        Room      `gorm:"foreignKey:ID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	IsAvailable bool      `gorm:"type:boolean"`
	BookingID   *uint     `gorm:"type:integer"`
	ScheduleID  *uint     `gorm:"type:integer"`
}
