package model

type ExternalMeeting struct {
	Meeting
	CompanyName string `gorm:"type:varchar(100);not null"`
}

type ExternalMeetingFactory struct {
	CompanyName string
}

func (f ExternalMeetingFactory) CreateMeeting(meeting Meeting) MeetingProductInterface {
	return &ExternalMeeting{
		Meeting: Meeting{
			Title:       meeting.Title,
			Description: meeting.Description,
			Location:    meeting.Location,
			Date:        meeting.Date,
			StartTime:   meeting.StartTime,
			EndTime:     meeting.EndTime,
			Attendees:   meeting.Attendees,
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
