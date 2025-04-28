// MEP-1013
package controller

import (
	model "ModEd/asset/model"
	"errors"

	"gorm.io/gorm"
)

type SupplyManagementController struct {
	db *gorm.DB
}

func NewSupplyManagementController(db *gorm.DB) *SupplyManagementController {
	return &SupplyManagementController{db: db}
}

func (a *SupplyManagementController) GetAll() (*[]model.SupplyManagement, error) {
	assetInfo := new([]model.SupplyManagement)
	result := a.db.Find(&assetInfo)
	return assetInfo, result.Error
}

func (a *SupplyManagementController) GetById(Id uint) (*model.SupplyManagement, error) {
	if Id == 0 {
		return nil, errors.New("no Id provide")
	}
	assetInfo := new(model.SupplyManagement)
	result := a.db.First(&assetInfo, "ID = ?", Id)
	return assetInfo, result.Error
}

func (a *SupplyManagementController) GetByRoomId(roomID uint) (*[]model.SupplyManagement, error) {
	if roomID == 0 {
		return nil, errors.New("no RoomID provided")
	}

	assetList := new([]model.SupplyManagement)
	result := a.db.Where("room_id = ?", roomID).Find(&assetList)

	return assetList, result.Error
}

func (a *SupplyManagementController) Create(payload *model.SupplyManagement) error {
	if payload == nil {
		return errors.New("invalid supply data")
	}
	result := a.db.Create(payload)
	return result.Error
}

func (a *SupplyManagementController) Update(Id uint, payload *model.SupplyManagement) error {
	if payload == nil || Id == 0 {
		return errors.New("invalid info")
	}
	existingAsset := new(model.SupplyManagement)
	if err := a.db.First(existingAsset, Id).Error; err != nil {
		return err
	}
	result := a.db.Model(existingAsset).Updates(payload)
	return result.Error
}

func (a *SupplyManagementController) Delete(Id uint) error {
	if Id == 0 {
		return errors.New("no Id provide")
	}
	assetInfo := new(model.SupplyManagement)
	result := a.db.Delete(&assetInfo, Id)
	return result.Error
}
