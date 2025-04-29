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
	Title       string                   `csv:"title" gorm:"type:varchar(100);not null"`
	Description string                   `csv:"description" gorm:"type:text"`
	Date        time.Time                `csv:"date" gorm:"not null"`
	Location    string                   `csv:"location" gorm:"type:varchar(255)"`
	StartTime   time.Time                `csv:"start_time" gorm:"not null"`
	EndTime     time.Time                `csv:"end_time"`
	Attendees   []commonModel.Instructor `gorm:"many2many:meeting_attendees;constraint:OnDelete:CASCADE"`
	*core.SerializableRecord
}

type MeetingProductInterface interface {
	GetTitle() string
	GetDate() time.Time
	GetLocation() string
}

type MeetingFactory interface {
	CreateMeeting(meeting Meeting) MeetingProductInterface
}

type RegularMeetingFactory struct{}

func (f RegularMeetingFactory) CreateMeeting(meeting Meeting) MeetingProductInterface {
	return &Meeting{
		Title:       meeting.Title,
		Description: meeting.Description,
		Location:    meeting.Location,
		Date:        meeting.Date,
		StartTime:   meeting.StartTime,
		EndTime:     meeting.EndTime,
		Attendees:   meeting.Attendees,
	}
}

func (m *Meeting) GetTitle() string {
	return m.Title
}

func (m *Meeting) GetDate() time.Time {
	return m.Date
}

func (m *Meeting) GetID() uint {
	return m.ID
}

func (m *Meeting) GetLocation() string {
	return m.Location
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
