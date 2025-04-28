package model

import (
	commonModel "ModEd/common/model"
	"time"
)

type OnlineMeeting struct {
	Meeting
	ZoomLink string `gorm:"type:varchar(255);not null"`
}

type OnlineMeetingFactory struct {
	ZoomLink string
}

func (f OnlineMeetingFactory) CreateMeeting(title, description, location string, date, startTime, endTime time.Time, attendees []commonModel.Instructor) MeetingProductInterface {
	return &OnlineMeeting{
		Meeting: Meeting{
			Title:       title,
			Description: description,
			Location:    location,
			Date:        date,
			StartTime:   startTime,
			EndTime:     endTime,
			Attendees:   attendees,
		},
		ZoomLink: f.ZoomLink,
	}
}

func (m OnlineMeeting) GetTitle() string {
	return m.Title
}

func (m OnlineMeeting) GetDetails() string {
	return m.Description
}
