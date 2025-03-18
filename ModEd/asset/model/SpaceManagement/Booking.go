package model

import (
	"time"

	"github.com/google/uuid"
)

type Booking struct {
	BookingID     	uuid.UUID	`gorm:"type:text;primaryKey" json:"booking_id" csv:"booking_id"`
	RoomID        	uuid.UUID	`gorm:"type:text;not null;index" json:"room_id" csv:"room_id"`
	Room          	Room		`gorm:"foreignKey:RoomID;references:RoomID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"room"`
	UserID        	uuid.UUID	`gorm:"type:text;not null" json:"user_id" csv:"user_id"`
	UserRole      	Role		`gorm:"type:text;not null" json:"user_role" csv:"user_role"`
	StartDate     	time.Time 	`gorm:"type:timestamp;not null" json:"start_date" csv:"start_date"`
	EndDate       	time.Time 	`gorm:"type:timestamp;not null" json:"end_date" csv:"end_date"`
	IsAvailable   	bool      	`gorm:"type:boolean;not null" json:"is_available" csv:"is_available"`
	EventName     	string    	`gorm:"type:text;not null" json:"event_name" csv:"event_name"`
	RoomStatus    	bool      	`gorm:"type:boolean;not null" json:"room_status" csv:"room_status"`
	CapacityLimit	int       	`gorm:"type:integer;not null" json:"capacity_limit" csv:"capacity_limit"`
}

// func (b *Booking) IsRoomAvailable(startTime, endTime time.Time) bool {
// 	if (startTime.After(b.StartDate) && startTime.Before(b.EndDate)) ||
// 		(endTime.After(b.StartDate) && endTime.Before(b.EndDate)) ||
// 		startTime.Equal(b.StartDate) || endTime.Equal(b.EndDate) {
// 		return false
// 	}
// 	return true
// }