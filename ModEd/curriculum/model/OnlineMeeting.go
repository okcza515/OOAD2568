package model

type OnlineMeeting struct {
	Meeting
	ZoomLink string `gorm:"type:varchar(255);not null"`
}

type OnlineMeetingFactory struct {
	ZoomLink string
}

func (f OnlineMeetingFactory) CreateMeeting(meeting Meeting) MeetingProductInterface {
	return &OnlineMeeting{
		Meeting: Meeting{
			Title:       meeting.Title,
			Description: meeting.Description,
			Location:    meeting.Location,
			Date:        meeting.Date,
			StartTime:   meeting.StartTime,
			EndTime:     meeting.EndTime,
			Attendees:   meeting.Attendees,
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
