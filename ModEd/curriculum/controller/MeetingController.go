//MEP-1008
package controller

import (
	commonModel "ModEd/common/model"
	"ModEd/core"
	model "ModEd/curriculum/model"

	"gorm.io/gorm"
)

type MeetingControllerService interface {
	GetAll() (*[]model.Meeting, error)
	GetByID(meetingID uint) (*model.Meeting, error)
	CreateMeeting(body *model.Meeting) error
	UpdateMeeting(meetingID uint, body *model.Meeting) error
	DeleteMeeting(meetingID uint) error
	AddAttendee(meetingID uint, instructorID uint) error
}

type MeetingController struct {
	*core.BaseController
	Connector *gorm.DB
}

func CreateMeetingController(db *gorm.DB) MeetingControllerService {
	return &MeetingController{
		BaseController: core.NewBaseController("Meeting", db),
		Connector:      db,
	}
}

func (c *MeetingController) GetAll() (*[]model.Meeting, error) {
	meetings := new([]model.Meeting)
	result := c.Connector.Find(&meetings)
	return meetings, result.Error
}

func (c *MeetingController) GetByID(meetingID uint) (*model.Meeting, error) {
	meetings := new(model.Meeting)
	result := c.Connector.First(&meetings, "ID = ?", meetingID)
	return meetings, result.Error
}

func (c *MeetingController) CreateMeeting(body *model.Meeting) error {
	result := c.Connector.Create(body)
	return result.Error
}

func (c *MeetingController) UpdateMeeting(meetingID uint, body *model.Meeting) error {
	body.ID = meetingID
	result := c.Connector.Updates(body)
	return result.Error
}

func (c *MeetingController) DeleteMeeting(meetingID uint) error {
	result := c.Connector.Model(&model.Meeting{}).Where("ID = ?", meetingID).Update("deleted_at", nil)
	return result.Error
}

func (c *MeetingController) AddAttendee(meetingID uint, instructorID uint) error {
	var meeting model.Meeting
	if err := c.Connector.First(&meeting, meetingID).Error; err != nil {
		return err
	}

	var instructor commonModel.Instructor
	if err := c.Connector.First(&instructor, instructorID).Error; err != nil {
		return err
	}

	if err := c.Connector.Model(&meeting).Association("Attendees").Append(&instructor); err != nil {
		return err
	}

	return nil
}
