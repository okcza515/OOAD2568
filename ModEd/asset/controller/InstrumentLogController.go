package controller

import (
	"ModEd/asset/model"

	"gorm.io/gorm"
)

type InstrumentLogController struct {
	Db *gorm.DB
}

func (c *InstrumentLogController) GetAll() (*[]model.InstrumentLog, error) {
	logs := new([]model.InstrumentLog)
	result := c.Db.Find(&logs)

	return logs, result.Error
}
