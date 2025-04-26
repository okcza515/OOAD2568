package controller

// MEP-1012 Asset

import (
	"ModEd/asset/model"
	"ModEd/core"

	"gorm.io/gorm"
)

type SupplyLogController struct {
	db *gorm.DB
	*core.BaseController[model.SupplyLog]
}

type SupplyLogControllerInterface interface {
	GetAll() (*[]model.SupplyLog, error)
	Insert(data core.RecordInterface) error
	RetrieveByID(id uint, preloads ...string) (core.RecordInterface, error)
}

func (c *SupplyLogController) GetAll() (*[]model.SupplyLog, error) {
	suppliesLogs := new([]model.SupplyLog)
	result := c.db.Find(&suppliesLogs)
	return suppliesLogs, result.Error
}
