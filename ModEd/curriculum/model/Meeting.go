// MEP-1008
package model

import (
	commonModel "ModEd/common/model"
	"ModEd/core"
	"fmt"
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
	*core.SerializableRecord
}

type OnlineMeeting struct {
	Meeting
	ZoomLink string `gorm:"type:varchar(255);not null"`
}

type ExternalMeeting struct {
	Meeting
	CompanyName string `gorm:"type:varchar(100);not null"`
}

func (m *Meeting) GetID() uint {
	return m.ID
}

func (m *Meeting) ToString() string {
	return fmt.Sprintf("%+v", m)
}

func (m *Meeting) Validate() error {
	if m.Title == "" {
		return fmt.Errorf("Title cannot be empty")
	}
	if m.Date.IsZero() {
		return fmt.Errorf("Date cannot be empty")
	}
	if m.StartTime.IsZero() {
		return fmt.Errorf("Start time cannot be empty")
	}
	if m.EndTime.IsZero() {
		return fmt.Errorf("End time cannot be empty")
	}
	if m.StartTime.After(m.EndTime) {
		return fmt.Errorf("Start time must be before end time")
	}
	return nil
}
