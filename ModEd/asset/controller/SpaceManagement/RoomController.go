package spacemanagement

import (
	model "ModEd/asset/model/spacemanagement"
	"errors"
	"gorm.io/gorm"
)

type RoomController struct {
	db *gorm.DB
}

func (c *RoomController) getAll() (*[]model.Room, error) {
	roomInfo := new([]model.Room)
	result := c.db.Find(&roomInfo)
	return roomInfo, result.Error
}

func (c *RoomController) getById(Id uint) (*model.Room, error) {
	if Id == 0 {
		return nil, errors.New("No Id provide")
	}
	roomInfo := new(model.Room)
	result := c.db.First(&roomInfo, "ID = ?", Id)
	return roomInfo, result.Error
}

func (c *RoomController) Create(payload *model.Room) error {
	if payload == nil {
		return errors.New("Invalid room data")
	}
	result := c.db.Create(payload)
	return result.Error
}

func (c *RoomController) Update(Id uint, payload *model.Room) error {
	if payload == nil || Id == 0 {
		return errors.New("Invalid info")
	}
	existingRoom := new(model.Room)
	if err := c.db.First(existingRoom, Id).Error; err != nil {
		return err
	}
	result := c.db.Model(&existingRoom).Updates(payload)
	return result.Error
}

func (c *RoomController) Delete(Id uint) error {
	if Id == 0 {
		return errors.New("No ID provides")
	}
	existingRoom := new(model.Room)
	if err := c.db.First(existingRoom, Id).Error; err != nil {
		return err
	}
	result := c.db.Delete(existingRoom)
	return result.Error
}
