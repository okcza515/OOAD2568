package controller

import (
	model "ModEd/asset/model"
	"ModEd/core"
	"ModEd/core/migration"
	"ModEd/utils/deserializer"
	"errors"

	"time"

	"gorm.io/gorm"
)

type RoomControllerInterface interface {
	SeedRoomsDatabase(path string) ([]*model.Room, error)
	ListAll() ([]string, error)
	List(condition map[string]interface{}, preloads ...string) ([]model.Room, error)
	RetrieveByID(id uint, preloads ...string) (model.Room, error)
	Insert(data model.Room) error
	UpdateByID(data model.Room) error
	DeleteByID(id uint) error
	DeleteAll() error
	InsertMany(data []model.Room) error

	// addObserver(observer SpaceManagementObserverInterface[model.Room])
	// removeObserver(observer SpaceManagementObserverInterface[model.Room])
}

type RoomController struct {
	db *gorm.DB
	*core.BaseController[model.Room]
	// observers map[string]SpaceManagementObserverInterface[model.Room]
}

func NewRoomController() *RoomController {
	db := migration.GetInstance().DB
	return &RoomController{
		db:             db,
		BaseController: core.NewBaseController[model.Room](db),
		// observers:      make(map[string]SpaceManagementObserverInterface[model.Room]),
	}
}

// func (c *RoomController) addObserver(observer SpaceManagementObserverInterface[model.Room]) {
// 	c.observers[observer.GetObserverID()] = observer
// }

// func (c *RoomController) removeObserver(observer SpaceManagementObserverInterface[model.Room]) {
// 	delete(c.observers, observer.GetObserverID())
// }

func (c *RoomController) SeedRoomsDatabase(path string) (rooms []*model.Room, err error) {
	deserializer, err := deserializer.NewFileDeserializer(path)
	if err != nil {
		return nil, errors.New("failed to create file deserializer")
	}
	if err := deserializer.Deserialize(&rooms); err != nil {
		return nil, errors.New("failed to deserialize curriculums")
	}
	for _, room := range rooms {
		err := c.Insert(*room)
		if err != nil {
			return nil, errors.New("failed to seed Room DB")
		}
	}
	return rooms, nil
}

func (c *RoomController) ListAll() ([]string, error) {
	result, err := c.List(nil)
	if err != nil {
		return nil, err
	}

	roomNames := make([]string, len(result))
	for i, room := range result {
		roomNames[i] = room.RoomName
	}

	return roomNames, nil
}

func (c *RoomController) List(condition map[string]interface{}, preloads ...string) ([]model.Room, error) {
	records, err := c.BaseController.List(condition, preloads...)
	return records, err
}

func (c *RoomController) RetrieveByID(id uint, preloads ...string) (model.Room, error) {
	if id == 0 {
		return model.Room{}, errors.New("no id provided")
	}
	result, err := c.BaseController.RetrieveByID(id, preloads...)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.Room{}, errors.New("room not found, please check the ID")
		}
		return model.Room{}, err
	}
	if result.IsRoomOutOfService {
		return model.Room{}, errors.New("room is out of service")
	}
	return result, nil
}

func (c *RoomController) Insert(data model.Room) error {
	return c.BaseController.Insert(data)
}

func (c *RoomController) UpdateByID(data model.Room) error {
	if data.GetID() == 0 {
		return errors.New("invalid info")
	}
	return c.BaseController.UpdateByID(data)
}

func (c *RoomController) DeleteByID(id uint) error {
	if id == 0 {
		return errors.New("no ID provided")
	}
	existingRoom := new(model.Room)
	if err := c.db.Unscoped().First(existingRoom, id).Error; err != nil {
		return errors.New("room not found, please check the ID")
	}
	result := c.db.Model(&existingRoom).UpdateColumn("DeletedAt", time.Now())
	return result.Error
}

func (c *RoomController) DeleteAll() error {
	rooms, err := c.List(nil)
	if err != nil {
		return err
	}
	for _, room := range rooms {
		if err := c.DeleteByID(room.ID); err != nil {
			return err
		}
	}
	return nil
}

func (c *RoomController) InsertMany(data []model.Room) error {
	if len(data) == 0 {
		return errors.New("no rooms to insert")
	}
	for _, room := range data {
		if err := c.Insert(room); err != nil {
			return err
		}
	}
	return nil
}
