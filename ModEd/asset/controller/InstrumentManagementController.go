// MEP-1013

package controller

import (
	model "ModEd/asset/model"
	"ModEd/core"
	"ModEd/core/migration"
	"errors"
	"gorm.io/gorm"
)

type InstrumentManagementInterface interface {
    List(condition map[string]interface{}, preloads ...string)([]model.InstrumentManagement, error) //Getall
    RetrieveByID(id uint, preloads ...string) (model.InstrumentManagement, error)
    RetrieveByRoomId(roomID uint) (*[]model.InstrumentManagement, error)
    Insert(payload *model.InstrumentManagement) error
    UpdateById(payload *model.InstrumentManagement) error
    DeleteByID(id uint) error
	InsertMany(data []model.InstrumentManagement) error
}

type InstrumentManagementController struct {
	db *gorm.DB
	*core.BaseController[model.InstrumentManagement]
}

func NewInstrumentManagementController() *InstrumentManagementController {
	db := migration.GetInstance().DB
	return &InstrumentManagementController{
		db : db,
		BaseController: core.NewBaseController[model.InstrumentManagement](db),
	}
}

func (c *InstrumentManagementController) List(condition map[string]interface{}, preloads ...string) ([]model.InstrumentManagement, error) {
    records, err := c.BaseController.List(condition, preloads...)
    if err != nil {
		//SOMETHING
        return nil, err
    }
    return records, err
}

func (c *InstrumentManagementController) RetrieveByID(id uint, preloads ...string) (model.InstrumentManagement, error) {
    record, err := c.BaseController.RetrieveByID(id, preloads...)
    if err != nil {
        return model.InstrumentManagement{}, err
    }
    return record, nil
}

func (c *InstrumentManagementController) RetrieveByRoomId(roomID uint) (*[]model.InstrumentManagement, error) {
	if roomID == 0 {
		return nil, errors.New("no RoomID provided")
	}

	assetList := new([]model.InstrumentManagement)
	result := c.db.Where("room_id = ?", roomID).Find(&assetList)

	return assetList, result.Error
}


func (c *InstrumentManagementController) Insert(payload *model.InstrumentManagement) error {
	if payload == nil {
		return errors.New("invalid instrument management data")
	}
	err := c.BaseController.Insert(*payload)
	return err
}

func (c *InstrumentManagementController) UpdateById(payload *model.InstrumentManagement) error {
    if payload == nil {
        return errors.New("payload cannot be nil")
    }
    if payload.GetID() == 0 {
        return errors.New("invalid ID: ID cannot be zero")
    }
    err := c.BaseController.UpdateByID(*payload)
    return err
}

func (c *InstrumentManagementController) DeleteByID(Id uint) error {
	if Id == 0 {
		return errors.New("no Id provide")
	 }
	err := c.BaseController.DeleteByID(Id)
	return err
}

//seed data
func (c *InstrumentManagementController) InsertMany(data []model.InstrumentManagement) error {
    if len(data) == 0 {
		return errors.New("no Instrument Management to insert")
	}

	err := c.BaseController.InsertMany(data)

	return err
}