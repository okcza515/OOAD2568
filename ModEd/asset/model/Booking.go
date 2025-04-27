// MEP-1013
package model

import (
	"gorm.io/gorm"
)

type Booking struct {
	gorm.Model
	TimeTableID uint      `gorm:"type:integer" json:"time_table_id" csv:"time_table_id"`
	TimeTable   TimeTable `gorm:"foreignKey:ID;references:ID" json:"time_table"`
	UserID      uint      `gorm:"type:integer;not null" json:"user_id" csv:"user_id"`
	UserRole    Role      `gorm:"type:text;not null" json:"user_role" csv:"user_role"`
	EventName   string    `gorm:"type:text;not null" json:"event_name" csv:"event_name"`
}
