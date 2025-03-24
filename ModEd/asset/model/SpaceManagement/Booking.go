package space

import (
	"time"

	"gorm.io/gorm"
)

type Booking struct {
	gorm.Model
	BookingID     uint      `gorm:"type:integer;primaryKey" json:"booking_id" csv:"booking_id"`
	RoomID        uint      `gorm:"type:integer;not null;index" json:"room_id" csv:"room_id"`
	Room          Room      `gorm:"foreignKey:RoomID;references:RoomID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"room"`
	UserID        uint      `gorm:"type:integer;not null" json:"user_id" csv:"user_id"`
	UserRole      Role      `gorm:"type:text;not null" json:"user_role" csv:"user_role"`
	StartDate     time.Time `gorm:"type:timestamp;not null" json:"start_date" csv:"start_date"`
	EndDate       time.Time `gorm:"type:timestamp;not null" json:"end_date" csv:"end_date"`
	IsAvailable   bool      `gorm:"type:boolean;not null" json:"is_available" csv:"is_available"`
	EventName     string    `gorm:"type:text;not null" json:"event_name" csv:"event_name"`
	CapacityLimit int       `gorm:"type:integer;not null" json:"capacity_limit" csv:"capacity_limit"`
}
