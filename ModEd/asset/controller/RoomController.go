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
	result, err := c.BaseController.List(nil)
	return &result, err
}

func (c *RoomController) GetById(Id uint) (*model.Room, error) {
	if Id == 0 {
		return nil, errors.New("no Id provide")
	}
	result, err := c.BaseController.RetrieveByID(Id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("room not found, please check the ID")
		}
		return nil, err
	}
	if result.IsRoomOutOfService {
		return nil, errors.New("room is out of service")
	}
	return &result, nil
}

func (c *RoomController) CreateRoom(payload *model.Room) error {
	if payload == nil {
		return errors.New("invalid room data")
	}
	err := c.BaseController.Insert(*payload)
	return err
}

func (c *RoomController) UpdateRoom(id uint, payload *model.Room) error {
	if payload == nil || id == 0 {
		return errors.New("invalid info")
	}
	payload.ID = id
	err := c.BaseController.UpdateByID(*payload)
	return err
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
	result, err := c.BaseController.List(nil)
	if err != nil {
		return err
	}
	for _, room := range result {
		if err := c.BaseController.DeleteByID(room.ID); err != nil {
			return err
		}
	}
	return nil
}
