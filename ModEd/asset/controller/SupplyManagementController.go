// MEP-1013
package controller

import (
	model "ModEd/asset/model"
	"ModEd/core"
	"ModEd/core/migration"
	"errors"
	"gorm.io/gorm"
)

type SupplyManagementInterface interface {
    List(condition map[string]interface{}, preloads ...string)([]model.SupplyManagement, error) //Getall
    RetrieveByID(id uint, preloads ...string) (model.SupplyManagement, error)
    RetrieveByRoomId(roomID uint) (*[]model.SupplyManagement, error)
    Insert(payload *model.SupplyManagement) error
    UpdateById(payload *model.SupplyManagement) error
    DeleteByID(id uint) error
    InsertMany(data []model.SupplyManagement) error
}

type SupplyManagementController struct {
	db *gorm.DB
	*core.BaseController[model.SupplyManagement]
}

func NewSupplyManagementController() *SupplyManagementController {
	db := migration.GetInstance().DB
	return &SupplyManagementController{
		db : db,
		BaseController: core.NewBaseController[model.SupplyManagement](db),
	}
}

func (c *SupplyManagementController) List(condition map[string]interface{}, preloads ...string) ([]model.SupplyManagement, error) {
	 // Add default preloads if none provided - only Room and Supply
     if len(preloads) == 0 {
        preloads = []string{"Room", "Supply"}
    }
    
    records, err := c.BaseController.List(condition, preloads...)
    if err != nil {
        return nil, err
    }
    return records, err
}

func (c *SupplyManagementController) RetrieveByID(id uint, preloads ...string) (model.SupplyManagement, error) {
     // Add default preloads if none provided - only Room and Supply
     if len(preloads) == 0 {
        preloads = []string{"Room", "Supply"}
    }
    
    record, err := c.BaseController.RetrieveByID(id, preloads...)
    if err != nil {
        return model.SupplyManagement{}, err
    }
    return record, nil
}

func (c *SupplyManagementController) RetrieveByRoomId(roomID uint) (*[]model.SupplyManagement, error) {
    if roomID == 0 {
		return nil, errors.New("no RoomID provided")
	}

	assetList := new([]model.SupplyManagement)
	// Add preloading to this query - only Room and Supply
	result := c.db.Preload("Room").Preload("Supply").Where("room_id = ?", roomID).Find(&assetList)

	return assetList, result.Error
}

func (c *SupplyManagementController) Insert(payload *model.SupplyManagement) error {
	if payload == nil {
		return errors.New("invalid supply management data")
	}
	err := c.BaseController.Insert(*payload)
	return err
}

func (c *SupplyManagementController) UpdateById(payload *model.SupplyManagement) error {
    if payload == nil {
        return errors.New("payload cannot be nil")
    }
    
    if payload.GetID() == 0 {
        return errors.New("invalid ID: ID cannot be zero")
    }
    
    err := c.BaseController.UpdateByID(*payload)
    return err
}

func (c *SupplyManagementController) DeleteByID(Id uint) error {
	if Id == 0 {
		return errors.New("no Id provide")
	 }
	err := c.BaseController.DeleteByID(Id)
	return err
}

//seed data
func (c *SupplyManagementController) InsertMany(data []model.SupplyManagement) error {
    if len(data) == 0 {
		return errors.New("no Supply Management to insert")
	}
	
	err := c.BaseController.InsertMany(data)

	return err
}
