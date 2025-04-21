package controller

// MEP-1012 Asset

import (
	"ModEd/asset/model"
	"ModEd/core"
	"gorm.io/gorm"
	"strconv"
)

type SupplyController struct {
	db *gorm.DB
	*core.BaseController
}

type SupplyControllerInterface interface {
	getAll() ([]model.Supply, error)
	ListAll() ([]string, error)
	Insert(data core.RecordInterface) error
	RetrieveByID(id uint, preloads ...string) (*core.RecordInterface, error)
	UpdateByID(data core.RecordInterface) error
	DeleteByID(id uint) error
	InsertMany(data interface{}) error
}

func (c *SupplyController) getAll() ([]model.Supply, error) {
	supplies := new([]model.Supply)
	result := c.db.Find(&supplies)
	return *supplies, result.Error
}

func (c *SupplyController) ListAll() ([]string, error) {
	supplies := new([]model.Supply)
	result := c.db.Find(&supplies)

	if result.Error != nil {
		return nil, result.Error
	}

	var resultList []string

	for _, supplie := range *supplies {
		resultList = append(resultList, "["+supplie.UpdatedAt.String()+"] "+string(supplie.SupplyLabel)+" "+strconv.FormatUint(uint64(supplie.ID), 10))
	}

	return resultList, result.Error
}

// func (c *SupplyController) GetByID(supplyID uint) (*model.Supply, error) {
// 	supply := new(model.Supply)
// 	result := c.db.First(&supply, "ID = ?", supplyID)
// 	return supply, result.Error
// }

// func (c *SupplyController) Create(body *model.Supply) error {
// 	result := c.db.Create(body)
// 	return result.Error
// }

// func (c *SupplyController) Update(supplyID uint, body *model.Supply) error {
// 	body.ID = supplyID
// 	result := c.db.Updates(body)
// 	return result.Error
// }

// func (c *SupplyController) Delete(supplyID uint) error {
// 	result := c.db.Model(&model.Supply{}).Where("ID = ?", supplyID).Update("deleted_at", time.Now())
// 	return result.Error
// }
