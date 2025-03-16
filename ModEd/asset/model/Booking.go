package model

import (
	"time"
)

type Booking struct {
	BookingID		string
	ClassroomID		string
	UserID          string
	UserRole        Role
	StartDate       time.Time
	EndDate         time.Time
	IsAvailable     bool
	EventName       string
	RoomStatus      bool
	CapacityLimit   int
}

// func (b *Booking) IsRoomAvailable(startTime, endTime time.Time) bool {
//     if (startTime.After(b.StartDate) && startTime.Before(b.EndDate)) ||
// 	(endTime.After(b.StartDate) && endTime.Before(b.EndDate)) ||
// 	(startTime.Equal(b.StartDate) || endTime.Equal(b.EndDate)) {
//         return false
//     }
//     return true
// }