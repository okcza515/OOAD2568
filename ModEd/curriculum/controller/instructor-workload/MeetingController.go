package controller

import (
	commonModel "ModEd/common/model"
	model "ModEd/curriculum/model/instructor-workload"

	"gorm.io/gorm"
)

type MeetingController struct {
	db *gorm.DB
}

func (c *MeetingController) GetAll() (*[]model.Meeting, error) {
	meetings := new([]model.Meeting)
	result := c.db.Find(&meetings)
	return meetings, result.Error
}

func (c *MeetingController) GetByID(meetingID uint) (*model.Meeting, error) {
	meetings := new(model.Meeting)
	result := c.db.First(&meetings, "ID = ?", meetingID)
	return meetings, result.Error
}

func (c *MeetingController) Create(body *model.Meeting) error {
	result := c.db.Create(body)
	return result.Error
}

func (c *MeetingController) Update(meetingID uint, body *model.Meeting) error {
	body.ID = meetingID
	result := c.db.Updates(body)
	return result.Error
}

func (c *MeetingController) Delete(meetingID uint) error {
	result := c.db.Model(&model.Meeting{}).Where("ID = ?", meetingID).Update("deleted_at", nil)
	return result.Error
}

func (c *MeetingController) AddAttendee(meetingID uint, instructorID uint) error {
	var meeting model.Meeting
	if err := c.db.First(&meeting, meetingID).Error; err != nil {
		return err
	}

	var instructor commonModel.Instructor
	if err := c.db.First(&instructor, instructorID).Error; err != nil {
		return err
	}

	if err := c.db.Model(&meeting).Association("Attendees").Append(&instructor); err != nil {
		return err
	}

	return nil
}
