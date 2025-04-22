// MEP-1013

package controller

import (
	model "ModEd/asset/model/spacemanagement"
	"errors"

	"gorm.io/gorm"
)

type InstrumentManagementAdapter struct {
	db *gorm.DB
}

func NewInstrumentManagementAdapter(db *gorm.DB) *InstrumentManagementAdapter {
	return &InstrumentManagementAdapter{db: db}
}

func (a *InstrumentManagementAdapter) getAllInstrumentManagement() (*[]model.InstrumentManagement, error) {
	assetInfo := new([]model.InstrumentManagement)
	result := a.db.Find(&assetInfo)
	return assetInfo, result.Error
}

func (a *InstrumentManagementAdapter) getInstrumentManagementById(Id uint) (*model.InstrumentManagement, error) {
	if Id == 0 {
		return nil, errors.New("no ID provided")
	}
	assetInfo := new(model.InstrumentManagement)
	result := a.db.First(&assetInfo, "ID = ?", Id)
	return assetInfo, result.Error
}

func (a *InstrumentManagementAdapter) GetInstrumentManagementByRoomId(roomID uint) (*[]model.InstrumentManagement, error) {
	if roomID == 0 {
		return nil, errors.New("no RoomID provided")
	}

	assetList := new([]model.InstrumentManagement)
	result := a.db.Where("room_id = ?", roomID).Find(&assetList)

	return assetList, result.Error
}

func (a *InstrumentManagementAdapter) CreateInstrumentManagement(payload *model.InstrumentManagement) error {
	if payload == nil {
		return errors.New("invalid instrument data")
	}
	result := a.db.Create(payload)
	return result.Error
}

func (a *InstrumentManagementAdapter) UpdateInstrumentManagement(Id uint, payload *model.InstrumentManagement) error {
	if payload == nil || Id == 0 {
		return errors.New("invalid info")
	}
	existingAsset := new(model.InstrumentManagement)
	if err := a.db.First(existingAsset, Id).Error; err != nil {
		return err
	}
	result := a.db.Model(existingAsset).Updates(payload)
	return result.Error
}

func (a *InstrumentManagementAdapter) DeleteInstrumentManagement(Id uint) error {
	if Id == 0 {
		return errors.New("no Id provide")
	}
	assetInfo := new(model.InstrumentManagement)
	result := a.db.Delete(&assetInfo, Id)
	return result.Error
}
