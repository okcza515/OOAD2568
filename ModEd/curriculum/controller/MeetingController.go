// MEP-1008
package controller

import (
	commonModel "ModEd/common/model"
	"ModEd/core"
	model "ModEd/curriculum/model"
	"errors"

	"gorm.io/gorm"
)

type MeetingControllerService interface {
	GetAll() (*[]model.Meeting, error)
	//GetByID(meetingID uint) (*model.Meeting, error)
	RetrieveByID(id uint) (*model.Meeting, error)
	CreateMeeting(body *model.Meeting) error
	CreateMeetingByFactory(factory model.MeetingFactory, meeting model.Meeting) error
	//UpdateMeeting(meetingID uint, body *model.Meeting) error
	UpdateByID(data model.Meeting) error
	//DeleteMeeting(meetingID uint) error
	DeleteByID(id uint) error
	AddAttendee(meetingID uint, instructorID uint) error
}

type MeetingController struct {
	*core.BaseController[*model.Meeting]
	Connector *gorm.DB
}

func CreateMeetingController(db *gorm.DB) *MeetingController {
	return &MeetingController{
		BaseController: core.NewBaseController[*model.Meeting](db),
		Connector:      db,
	}
}

func (c *MeetingController) GetAll() (*[]model.Meeting, error) {
	meetings := new([]model.Meeting)
	result := c.Connector.Find(&meetings)
	return meetings, result.Error
}

// func (c *MeetingController) GetByID(meetingID uint) (*model.Meeting, error) {
// 	meetings := new(model.Meeting)
// 	result := c.Connector.First(&meetings, "ID = ?", meetingID)
// 	return meetings, result.Error
// }

func (c *MeetingController) CreateMeeting(body *model.Meeting) error {
	result := c.Connector.Create(body)
	return result.Error
}

func (c *MeetingController) CreateMeetingByFactory(factory model.MeetingFactory, meeting model.Meeting) error {
	meetingProduct := factory.CreateMeeting(
		meeting.Title,
		meeting.Description,
		meeting.Location,
		meeting.Date,
		meeting.StartTime,
		meeting.EndTime,
		meeting.Attendees,
	)

	// Store the created meeting in the database
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

// func (c *MeetingController) UpdateMeeting(meetingID uint, body *model.Meeting) error {
// 	body.ID = meetingID
// 	result := c.Connector.Updates(body)
// 	return result.Error
// }

// func (c *MeetingController) DeleteMeeting(meetingID uint) error {
// 	result := c.Connector.Model(&model.Meeting{}).Where("ID = ?", meetingID).Update("deleted_at", nil)
// 	return result.Error
// }

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
