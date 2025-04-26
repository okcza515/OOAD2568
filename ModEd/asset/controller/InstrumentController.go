package controller

// MEP-1012 Asset

import (
	"ModEd/asset/model"
	"ModEd/core"
	"fmt"
	"gorm.io/gorm"
)

type InstrumentController struct {
	db *gorm.DB
	*core.BaseController[model.Instrument]
}

type InstrumentControllerInterface interface {
	ListAll() ([]string, error)
	List(condition map[string]interface{}) ([]model.Instrument, error)
	RetrieveByID(id uint, preloads ...string) (model.Instrument, error)
	Insert(data model.Instrument) error
	UpdateByID(data model.Instrument) error
	DeleteByID(id uint) error
	InsertMany(data []model.Instrument) error
}

func (c *InstrumentController) ListAll() ([]string, error) {
	instruments := new([]model.Instrument)
	result := c.db.Find(&instruments)

	if result.Error != nil {
		return nil, result.Error
	}

	var resultList []string

	for _, instrument := range *instruments {
		resultList = append(resultList, fmt.Sprintf("[%v] %v", instrument.InstrumentCode, instrument.InstrumentLabel))
	}

	return resultList, result.Error
}
