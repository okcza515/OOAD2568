package controller

import (
	model "ModEd/asset/model"
	"ModEd/core"
	"ModEd/utils/deserializer"
	"errors"

	"time"

	"gorm.io/gorm"
)

type RoomControllerInterface interface {
	SeedRoomsDatabase(path string) ([]*model.Room, error)
	GetAll() (*[]model.Room, error)
	GetById(Id uint) (*model.Room, error)
	CreateRoom(payload *model.Room) error
	UpdateRoom(Id uint, payload *model.Room) error
	DeleteRoom(Id uint) error
	DeleteAllRooms() error
}

type RoomController struct {
	db *gorm.DB
	*core.BaseController[model.Room]
}

func (c *RoomController) SeedRoomsDatabase(path string) (rooms []*model.Room, err error) {
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
			return nil, errors.New("failed to seed Room DB")
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
		return nil, errors.New("no Id provide")
	}
	roomInfo := new(model.Room)
	result := c.db.First(&roomInfo, "ID = ?", Id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("room not found, please check the ID")
		}
		return nil, result.Error
	}
	if roomInfo.IsRoomOutOfService {
		return nil, errors.New("room is out of service")
	}
	return roomInfo, nil
}

func (c *RoomController) CreateRoom(payload *model.Room) error {
	if payload == nil {
		return errors.New("invalid room data")
	}
	result := c.db.Create(payload)
	return result.Error
}

func (c *RoomController) UpdateRoom(Id uint, payload *model.Room) error {
	if payload == nil || Id == 0 {
		return errors.New("invalid info")
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
		return errors.New("no ID provided")
	}
	existingRoom := new(model.Room)
	if err := c.db.Unscoped().First(existingRoom, Id).Error; err != nil {
		return errors.New("room not found, please check the ID")
	}
	result := c.db.Model(&existingRoom).UpdateColumn("DeletedAt", time.Now())
	return result.Error
}

func (c *RoomController) DeleteAllRooms() error {
	allRooms := new([]model.Room)
	result := c.db.Find(&allRooms)
	if result.Error != nil {
		return result.Error
	}
	for _, room := range *allRooms {
		if err := c.db.Model(&room).UpdateColumn("DeletedAt", time.Now()).Error; err != nil {
			return err
		}
	}
	return nil
}
