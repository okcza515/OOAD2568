package spacemanagement

import (
	model "ModEd/asset/model/SpaceManagement"
	"ModEd/utils/deserializer"
	"errors"

	"gorm.io/gorm"
)

type RoomController struct {
	db *gorm.DB
}

func (c *RoomController) CreateSeedRooms(path string) (rooms []*model.Room, err error) {
	deserializer, err := deserializer.NewFileDeserializer(path)
	if err != nil {
		return nil, errors.New("failed to create file deserializer")
	}
	if err := deserializer.Deserialize(&rooms); err != nil {
		return nil, errors.New("failed to deserialize curriculums")
	}
	for _, room := range rooms {
		err := c.CreateRoom(room)
		if err != nil {
			return nil, errors.New("Failed to seed Room DB")
		}
	}
	return rooms, nil
}

func (c *RoomController) GetAll() (*[]model.Room, error) {
	roomInfo := new([]model.Room)
	result := c.db.Find(&roomInfo)
	return roomInfo, result.Error
}

func (c *RoomController) GetById(Id uint) (*model.Room, error) {
	if Id == 0 {
		return nil, errors.New("No Id provide")
	}
	roomInfo := new(model.Room)
	result := c.db.First(&roomInfo, "ID = ?", Id)
	return roomInfo, result.Error
}

func (c *RoomController) CreateRoom(payload *model.Room) error {
	if payload == nil {
		return errors.New("Invalid room data")
	}
	result := c.db.Create(payload)
	return result.Error
}

func (c *RoomController) UpdateRoom(Id uint, payload *model.Room) error {
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

func (c *RoomController) DeleteRoom(Id uint) error {
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
