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
	*core.BaseController
}

type InstrumentControllerInterface interface {
	ListAll() ([]string, error)
	RetrieveByID(id uint, preloads ...string) (*core.RecordInterface, error)
	Insert(data core.RecordInterface) error
	UpdateByID(data core.RecordInterface) error
	DeleteByID(id uint) error
	InsertMany(data []core.RecordInterface) error
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
