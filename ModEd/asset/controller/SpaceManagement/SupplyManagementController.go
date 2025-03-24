//MEP-1013
package spacemanagement

import(
	model "ModEd/asset/model/SpaceManagement"
	"gorm.io/gorm"
	"errors"
)

type SupplyManagementController struct {
	db *gorm.DB
}

func (c* SupplyManagementController) getAllSupplyManagement() (*[]model.SupplyManagement, error) {
	assetInfo := new([]model.SupplyManagement)
	result := c.db.Find(&assetInfo)
	return assetInfo, result.Error
}

func (c* SupplyManagementController) getSupplyManagementById(Id uint) (*model.SupplyManagement, error) {
	if Id == 0 {
		return nil, errors.New("No Id provide")
	}
	assetInfo := new(model.SupplyManagement)
	result := c.db.First(&assetInfo, "ID = ?", Id)
	return assetInfo, result.Error
}

func (c *SupplyManagementController) GetAssetManagementByRoomId(roomID uint) (*[]model.SupplyManagement, error) {
    if roomID == 0 {
        return nil, errors.New("No RoomID provided")
    }
    
    assetList := new([]model.SupplyManagement)
    result := c.db.Where("room_id = ?", roomID).Find(&assetList)
    
    return assetList, result.Error
}


func (c* SupplyManagementController) CreateAssetManagement(payload *model.SupplyManagement) error {
	if payload == nil {
		return errors.New("Invalid asset data")
	}
	result := c.db.Create(payload)
	return result.Error
}

func (c* SupplyManagementController) UpdateSupplyManagement(Id uint, payload *model.SupplyManagement) error {
	if payload == nil || Id == 0 {
		return errors.New("Invalid info")
	}
	existingAsset := new(model.SupplyManagement)
	if err := c.db.First(existingAsset, Id).Error; err != nil {
		return err
	}
	result := c.db.Model(existingAsset).Updates(payload)
	return result.Error
}

func (c* SupplyManagementController) DeleteSupplyManagement(Id uint) error {
	if Id == 0 {
		return errors.New("No Id provide")
	}
	assetInfo := new(model.SupplyManagement)
	result := c.db.Delete(&assetInfo, Id)
	return result.Error
}