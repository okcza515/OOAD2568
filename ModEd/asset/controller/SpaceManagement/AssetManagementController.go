//MEP-1013
package spacemanagement

import(
	model "ModEd/asset/model/SpaceManagement"
	"gorm.io/gorm"
	"errors"
)

type AssetManagementController struct {
	db *gorm.DB
}

func (c* AssetManagementController) getAll() (*[]model.AssetManagement, error) {
	assetInfo := new([]model.AssetManagement)
	result := c.db.Find(&assetInfo)
	return assetInfo, result.Error
}

func (c* AssetManagementController) getById(Id uint) (*model.AssetManagement, error) {
	if Id == 0 {
		return nil, errors.New("No Id provide")
	}
	assetInfo := new(model.AssetManagement)
	result := c.db.First(&assetInfo, "ID = ?", Id)
	return assetInfo, result.Error
}

func (c *AssetManagementController) GetByRoomId(roomID uint) (*[]model.AssetManagement, error) {
    if roomID == 0 {
        return nil, errors.New("No RoomID provided")
    }
    
    assetList := new([]model.AssetManagement)
    result := c.db.Where("room_id = ?", roomID).Find(&assetList)
    
    return assetList, result.Error
}


func (c* AssetManagementController) CreateAssetManagement(payload *model.AssetManagement) error {
	if payload == nil {
		return errors.New("Invalid asset data")
	}
	result := c.db.Create(payload)
	return result.Error
}

func (c* AssetManagementController) UpdateAssetManagement(Id uint, payload *model.AssetManagement) error {
	if payload == nil || Id == 0 {
		return errors.New("Invalid info")
	}
	existingAsset := new(model.AssetManagement)
	if err := c.db.First(existingAsset, Id).Error; err != nil {
		return err
	}
	result := c.db.Model(existingAsset).Updates(payload)
	return result.Error
}

func (c* AssetManagementController) DeleteAssetManagement(Id uint) error {
	if Id == 0 {
		return errors.New("No Id provide")
	}
	assetInfo := new(model.AssetManagement)
	result := c.db.Delete(&assetInfo, Id)
	return result.Error
}