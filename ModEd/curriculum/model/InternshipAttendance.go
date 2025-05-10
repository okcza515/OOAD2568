// MEP-1009 Student Internship
package model

import (
	"ModEd/core"
	"time"
)

type Attendance struct {
	core.BaseModel

	Date          time.Time             `gorm:"type:date"`
	CheckInTime   time.Time             `gorm:"type:time"`
	CheckOutTime  time.Time             `gorm:"type:time"`
	CheckInStatus bool                  `gorm:"type:bool"`
	AssingWork    string                `gorm:"type:varchar(255)"`
	StudentInfoID uint                  `gorm:"not null"`
	Student       InternshipInformation `gorm:"foreignKey:StudentCode;references:StudentCode"`
}
