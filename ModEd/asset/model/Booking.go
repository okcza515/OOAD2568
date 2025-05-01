// MEP-1013
package model

import (
	"ModEd/core"
)

type Booking struct {
	core.BaseModel
	TimeTableID uint      `gorm:"type:integer;not null;uniqueIndex" json:"time_table_id" csv:"time_table_id"`
	TimeTable   TimeTable `gorm:"foreignKey:TimeTableID;references:ID" json:"time_table" csv:"time_table"`
	UserID      uint      `gorm:"type:integer;not null" json:"user_id" csv:"user_id"`
	UserRole    Role      `gorm:"type:text;not null" json:"user_role" csv:"user_role"`
	EventName   string    `gorm:"type:text;not null" json:"event_name" csv:"event_name"`
}
