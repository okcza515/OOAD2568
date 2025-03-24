//MEP-1013
package spacemanagement

import(
	model "ModEd/asset/model/SpaceManagement"
	"gorm.io/gorm"
	"errors"
)

type InstrumentManagementController struct {
	db *gorm.DB
}

func (c* InstrumentManagementController) getAllInstrumentManagement() (*[]model.InstrumentManagement, error) {
	assetInfo := new([]model.InstrumentManagement)
	result := c.db.Find(&assetInfo)
	return assetInfo, result.Error
}

func (c* InstrumentManagementController) getInstrumentManagementById(Id uint) (*model.InstrumentManagement, error) {
	if Id == 0 {
		return nil, errors.New("No Id provide")
	}
	assetInfo := new(model.InstrumentManagement)
	result := c.db.First(&assetInfo, "ID = ?", Id)
	return assetInfo, result.Error
}

func (c *InstrumentManagementController) GetInstrumentManagementByRoomId(roomID uint) (*[]model.InstrumentManagement, error) {
    if roomID == 0 {
        return nil, errors.New("No RoomID provided")
    }
    
    assetList := new([]model.InstrumentManagement)
    result := c.db.Where("room_id = ?", roomID).Find(&assetList)
    
    return assetList, result.Error
}


func (c* InstrumentManagementController) CreateInstrumentManagement(payload *model.InstrumentManagement) error {
	if payload == nil {
		return errors.New("Invalid asset data")
	}
	result := c.db.Create(payload)
	return result.Error
}

func (c* InstrumentManagementController) UpdateInstrumentManagement(Id uint, payload *model.InstrumentManagement) error {
	if payload == nil || Id == 0 {
		return errors.New("Invalid info")
	}
	existingAsset := new(model.InstrumentManagement)
	if err := c.db.First(existingAsset, Id).Error; err != nil {
		return err
	}
	result := c.db.Model(existingAsset).Updates(payload)
	return result.Error
}

func (c* InstrumentManagementController) DeleteInstrumentManagement(Id uint) error {
	if Id == 0 {
		return errors.New("No Id provide")
	}
	assetInfo := new(model.InstrumentManagement)
	result := c.db.Delete(&assetInfo, Id)
	return result.Error
}