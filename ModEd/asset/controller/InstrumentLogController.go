package controller

// MEP-1012 Asset

import (
	"ModEd/asset/model"
	"ModEd/core"
	"gorm.io/gorm"
	"strconv"
)

type InstrumentLogController struct {
	db *gorm.DB
	*core.BaseController[model.InstrumentLog]
}

type InstrumentLogControllerInterface interface {
	getAll() ([]model.InstrumentLog, error)
	Insert(data model.InstrumentLog) error
	InsertMany(data []model.InstrumentLog) error
	RetrieveByID(id uint, preloads ...string) (model.InstrumentLog, error)
	List(condition map[string]interface{}) ([]model.InstrumentLog, error)
}

func (c *InstrumentLogController) getAll() ([]model.InstrumentLog, error) {
	logs := new([]model.InstrumentLog)
	result := c.db.Find(&logs)

	return *logs, result.Error
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
