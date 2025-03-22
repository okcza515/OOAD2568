package asset

import (
	model "ModEd/asset/model/asset"
	"gorm.io/gorm"
	"strconv"
)

type InstrumentLogController struct {
	Db *gorm.DB
}

func (c *InstrumentLogController) getAll() (*[]model.InstrumentLog, error) {
	logs := new([]model.InstrumentLog)
	result := c.Db.Find(&logs)

	return logs, result.Error
}

func (c *InstrumentLogController) ListAll() ([]string, error) {
	logs := new([]model.InstrumentLog)
	result := c.Db.Find(&logs)

	if result.Error != nil {
		return nil, result.Error
	}

	var resultList []string

	for _, log := range *logs {
		resultList = append(resultList, "["+log.UpdatedAt.String()+"] "+string(log.Action)+" "+strconv.FormatUint(uint64(log.InstrumentID), 10))
	}

	return resultList, result.Error
}
