package asset

import (
	model "ModEd/asset/model/asset"

	"gorm.io/gorm"
)

type InstrumentLogController struct {
	Db *gorm.DB
}

func (c *InstrumentLogController) getAll() (*[]model.InstrumentLog, error) {
	logs := new([]model.InstrumentLog)
	result := c.Db.Find(&logs)

	return logs, result.Error
}
