package asset

import (
	"ModEd/asset/model/asset"
	"gorm.io/gorm"
)

type InstrumentController struct {
	db *gorm.DB
}

func (ins *InstrumentController) CreateNewInstrument(instrumentData *[]asset.Instrument) error {
	result := ins.db.Create(instrumentData)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
