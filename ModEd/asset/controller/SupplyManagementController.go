// MEP-1013
package controller

import (
	model "ModEd/asset/model"
	"errors"
	"gorm.io/gorm"
	"ModEd/core"
)

type SupplyManagementController struct {
	db *gorm.DB
	*core.BaseController[model.SupplyManagement]
}

func NewSupplyManagementController(db *gorm.DB) *SupplyManagementController {
	return &SupplyManagementController{BaseController: core.NewBaseController[model.SupplyManagement](db)}
}

func (c *SupplyManagementController) GetAll() (*[]model.SupplyManagement, error) {
	assetList := []model.SupplyManagement{}
    records, err := c.BaseController.List(nil)
    assetList = records
    return &assetList, err
}

func (c *SupplyManagementController) GetById(Id uint) (*model.SupplyManagement, error) {
    asset := model.SupplyManagement{}
    record, err := c.BaseController.RetrieveByID(Id)
    if err != nil {
        return nil, err
    }
    asset = record
    return &asset, nil
}

func (c *SupplyManagementController) GetByRoomId(roomID uint) (*[]model.SupplyManagement, error) {
	if roomID == 0 {
		return nil, errors.New("no RoomID provided")
	}

	assetList := new([]model.SupplyManagement)
	result := c.db.Where("room_id = ?", roomID).Find(&assetList)

	return assetList, result.Error
}

func (c *SupplyManagementController) Create(payload *model.SupplyManagement) error {
	if payload == nil {
		return errors.New("invalid supply management data")
	}
	err := c.BaseController.Insert(*payload)
	return err
}

func (c *SupplyManagementController) Update(Id uint, payload *model.SupplyManagement) error {
	if payload == nil || Id == 0 {
		return errors.New("invalid info")
	}
	err := c.BaseController.UpdateByID(*payload)
	return err
}

func (c *SupplyManagementController) Delete(Id uint) error {
	if Id == 0 {
		return errors.New("no Id provide")
	 }
	err := c.BaseController.DeleteByID(Id)
	return err
}

func (c* InstrumentManagementController) SeedSupplyManagementDatabase(path string)(){
	
}
