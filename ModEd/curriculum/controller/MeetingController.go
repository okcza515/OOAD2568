// MEP-1008
package controller

import (
	"ModEd/core"
	model "ModEd/curriculum/model"
	"errors"

	"gorm.io/gorm"
)

type MeetingControllerInterface interface {
	List(condition map[string]interface{}) ([]*model.Meeting, error)

	RetrieveByID(id uint) (*model.Meeting, error)
	CreateMeeting(body *model.Meeting) error
	CreateMeetingByFactory(factory model.MeetingFactory, meeting model.Meeting) error
	UpdateByID(data model.Meeting) error
	DeleteByID(id uint) error
	AddAttendee(meetingID uint, instructorID uint) error
}

type MeetingController struct {
	*core.BaseController[*model.Meeting]
	Connector *gorm.DB
}

func NewMeetingController(db *gorm.DB) *MeetingController {
	return &MeetingController{
		BaseController: core.NewBaseController[*model.Meeting](db),
		Connector:      db,
	}
}

func (c *MeetingController) CreateMeeting(body *model.Meeting) error {
	result := c.Connector.Create(body)
	return result.Error
}

func (c *MeetingController) CreateMeetingByFactory(factory model.MeetingFactory, meeting model.Meeting) error {
	meetingProduct := factory.CreateMeeting(meeting)

	switch m := meetingProduct.(type) {
	case *model.Meeting:
		return c.Connector.Create(m).Error
	case *model.ExternalMeeting:
		return c.Connector.Create(m).Error
	case *model.OnlineMeeting:
		return c.Connector.Create(m).Error
	default:
		return errors.New("unsupported meeting type")
	}
}

func (c *MeetingController) AddAttendee(meetingID uint, attendee model.AttendeeAdapter) error {
	var meeting model.Meeting
	if err := c.Connector.First(&meeting, meetingID).Error; err != nil {
		return err
	}

	attendeeRecord := model.MeetingAttendee{
		MeetingID:    meetingID,
		AttendeeCode: attendee.GetCode(),
		AttendeeType: attendee.GetType(),
	}

	if err := c.Connector.First(&meeting, &meetingID).Error; err != nil {
		return err
	}

	if err := c.Connector.Model(&meeting).Association("Attendees").Append(&attendeeRecord); err != nil {
		return err
	}
	return nil
}
