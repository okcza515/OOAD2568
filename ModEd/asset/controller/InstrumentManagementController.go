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
    GetAll() (*[]model.InstrumentManagement, error)
    GetById(id uint) (*model.InstrumentManagement, error)
    GetByRoomId(roomID uint) (*[]model.InstrumentManagement, error)
    Create(payload *model.InstrumentManagement) error
    Update(id uint, payload *model.InstrumentManagement) error
    Delete(id uint) error
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

func (c *InstrumentManagementController) GetAll() (*[]model.InstrumentManagement, error) {
    assetList := []model.InstrumentManagement{}
    records, err := c.BaseController.List(nil)
    assetList = records
    return &assetList, err
}

func (c *InstrumentManagementController) GetById(Id uint) (*model.InstrumentManagement, error) {
	asset := model.InstrumentManagement{}
	record, err := c.BaseController.RetrieveByID(Id)
	asset = record
	return &asset, err
}

func (c *InstrumentManagementController) GetByRoomId(roomID uint) (*[]model.InstrumentManagement, error) {
	if roomID == 0 {
		return nil, errors.New("no RoomID provided")
	}

	assetList := new([]model.InstrumentManagement)
	result := c.db.Where("room_id = ?", roomID).Find(&assetList)

	return assetList, result.Error
}


func (c *InstrumentManagementController) Create(payload *model.InstrumentManagement) error {
	if payload == nil {
		return errors.New("invalid instrument management data")
	}
	err := c.BaseController.Insert(*payload)
	return err
}

func (c *InstrumentManagementController) Update(Id uint, payload *model.InstrumentManagement) error {
	if payload == nil || Id == 0 {
		return errors.New("invalid info")
	}
	err := c.BaseController.UpdateByID(*payload)
	return err
}

func (c *InstrumentManagementController) Delete(Id uint) error {
	if Id == 0 {
		return errors.New("no Id provide")
	 }
	err := c.BaseController.DeleteByID(Id)
	return err
}