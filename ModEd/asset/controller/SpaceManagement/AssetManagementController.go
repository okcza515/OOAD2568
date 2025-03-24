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

func (c* AssetManagementController) Create(payload *model.AssetManagement) error {
	if payload == nil {
		return errors.New("Invalid asset data")
	}
	result := c.db.Create(payload)
	return result.Error
}