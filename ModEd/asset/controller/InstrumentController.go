package controller

import (
	"ModEd/asset/model"
	"gorm.io/gorm"
)

type InstrumentController struct {
	Db *gorm.DB
}

func (ins *InstrumentController) CreateNewInstrument(instrumentData *[]model.Instrument) error {
	result := ins.Db.Create(instrumentData)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
