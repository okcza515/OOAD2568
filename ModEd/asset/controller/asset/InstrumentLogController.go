package asset

import (
	"ModEd/asset/model/asset"

	"gorm.io/gorm"
)

type InstrumentLogController struct {
	Db *gorm.DB
}

func (c *InstrumentLogController) getAll() (*[]asset.InstrumentLog, error) {
	logs := new([]asset.InstrumentLog)
	result := c.Db.Find(&logs)

	return logs, result.Error
}
