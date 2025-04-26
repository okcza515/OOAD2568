package controller

// MEP-1012 Asset

import (
	"ModEd/asset/model"
	"ModEd/core"
	"strconv"

	"gorm.io/gorm"
)

type SupplyController struct {
	db *gorm.DB
	*core.BaseController[model.Supply]
}

type SupplyControllerInterface interface {
	ListAll() ([]string, error)
	List(condition map[string]interface{}) ([]model.Supply, error)
	RetrieveByID(id uint, preloads ...string) (model.Supply, error)
	Insert(data model.Supply) error
	UpdateByID(data model.Supply) error
	DeleteByID(id uint) error
	InsertMany(data []model.Supply) error
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
