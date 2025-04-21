package model

import (
	commonModel "ModEd/common/model"
	"time"

	"gorm.io/gorm"
)

type Meeting struct {
	gorm.Model
	Title       string                   `gorm:"type:varchar(100);not null"`
	Description string                   `gorm:"type:text"`
	Date        time.Time                `gorm:"not null"`
	Location    string                   `gorm:"type:varchar(255)"`
	Attendees   []commonModel.Instructor `gorm:"many2many:meeting_attendees;constraint:OnDelete:CASCADE"`
	StartTime   time.Time                `gorm:"not null"`
	EndTime     time.Time                `gorm:""`
}

type OnlineMeeting struct {
	Meeting
	ZoomLink string `gorm:"type:varchar(255);not null"`
}

type ExternalMeeting struct {
	Meeting
	CompanyName string `gorm:"type:varchar(100);not null"`
}
