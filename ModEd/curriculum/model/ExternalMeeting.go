package model

import (
	commonModel "ModEd/common/model"
	"time"
)

type ExternalMeeting struct {
	Meeting
	CompanyName string `gorm:"type:varchar(100);not null"`
}

type ExternalMeetingFactory struct {
	CompanyName string
}

func (f ExternalMeetingFactory) CreateMeeting(title, description, location string, date, startTime, endTime time.Time, attendees []commonModel.Instructor) MeetingProductInterface {
	return &ExternalMeeting{
		Meeting: Meeting{
			Title:       title,
			Description: description,
			Location:    location,
			Date:        date,
			StartTime:   startTime,
			EndTime:     endTime,
			Attendees:   attendees,
		},
		CompanyName: f.CompanyName,
	}
}

func (em *ExternalMeeting) GetTitle() string {
	return em.Title
}

func (em *ExternalMeeting) GetID() uint {
	return em.ID
}
