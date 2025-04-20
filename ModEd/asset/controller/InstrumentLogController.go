package controller

import (
	"ModEd/asset/model"
	"gorm.io/gorm"
	"strconv"
)

type InstrumentLogController struct {
	db *gorm.DB
}

func (c *InstrumentLogController) getAll() (*[]model.InstrumentLog, error) {
	logs := new([]model.InstrumentLog)
	result := c.db.Find(&logs)

	return logs, result.Error
}

func (c *InstrumentLogController) ListAll() ([]string, error) {
	logs := new([]model.InstrumentLog)
	result := c.db.Find(&logs)

	if result.Error != nil {
		return nil, result.Error
	}

	var resultList []string

	for _, log := range *logs {
		resultList = append(resultList, "["+log.UpdatedAt.String()+"] "+string(log.Action)+" "+strconv.FormatUint(uint64(log.InstrumentID), 10))
	}

	return resultList, result.Error
}
