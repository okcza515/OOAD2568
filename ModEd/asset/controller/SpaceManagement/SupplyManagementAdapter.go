// MEP-1013
package spacemanagement

import (
	model "ModEd/asset/model/spacemanagement"
	"errors"
	"gorm.io/gorm"
)

type SupplyManagementAdapter struct {
	db *gorm.DB
}

func NewSupplyManagementAdapter(db *gorm.DB) *SupplyManagementAdapter {
    return &SupplyManagementAdapter{db: db}
}

func (a *SupplyManagementAdapter) getAllSupplyManagement() (*[]model.SupplyManagement, error) {
	assetInfo := new([]model.SupplyManagement)
	result := a.db.Find(&assetInfo)
	return assetInfo, result.Error
}

func (a *SupplyManagementAdapter) getSupplyManagementById(Id uint) (*model.SupplyManagement, error) {
	if Id == 0 {
		return nil, errors.New("no Id provide")
	}
	assetInfo := new(model.SupplyManagement)
	result := a.db.First(&assetInfo, "ID = ?", Id)
	return assetInfo, result.Error
}

func (a *SupplyManagementAdapter) GetAssetManagementByRoomId(roomID uint) (*[]model.SupplyManagement, error) {
	if roomID == 0 {
		return nil, errors.New("no RoomID provided")
	}

	assetList := new([]model.SupplyManagement)
	result := a.db.Where("room_id = ?", roomID).Find(&assetList)

	return assetList, result.Error
}

func (a *SupplyManagementAdapter) CreateAssetManagement(payload *model.SupplyManagement) error {
	if payload == nil {
		return errors.New("invalid supply data")
	}
	result := a.db.Create(payload)
	return result.Error
}

func (a *SupplyManagementAdapter) UpdateSupplyManagement(Id uint, payload *model.SupplyManagement) error {
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

func (a *SupplyManagementAdapter) DeleteSupplyManagement(Id uint) error {
	if Id == 0 {
		return errors.New("no Id provide")
	}
	assetInfo := new(model.SupplyManagement)
	result := a.db.Delete(&assetInfo, Id)
	return result.Error
}
