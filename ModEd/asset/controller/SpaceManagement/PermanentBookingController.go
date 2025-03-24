// MEP-1013
package spacemanagement

import (
	model "ModEd/asset/model/SpaceManagement"
	"errors"

	"gorm.io/gorm"
)

type PermanentScheduleController struct {
	DB *gorm.DB
}

func (c *PermanentScheduleController) CheckRoomInService(roomID uint) (*bool, error) {
	var room model.Room

	err := c.DB.First(&room, roomID).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("Room not found")
	} else if err != nil {
		return nil, err
	}

	if room.IsRoomOutOfService {
		return nil, errors.New("Room is out of service")
	}

	isInService := !room.IsRoomOutOfService
	return &isInService, nil
}

func (c *PermanentScheduleController) CreateSubjectSchedule(body *model.PermanentSchedule) error {
	if body.StartDate.IsZero() || body.EndDate.IsZero() {
		return errors.New("Start date and end date are required")
	}
	result := c.DB.Create(body)
	return result.Error
}

func (c *PermanentScheduleController) UpdateSubjectSchedule()
