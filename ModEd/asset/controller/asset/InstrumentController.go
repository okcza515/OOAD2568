package asset

import (
	"ModEd/asset/model/asset"
	"gorm.io/gorm"
)

type InstrumentController struct {
	Db *gorm.DB
}

func (ins *InstrumentController) CreateNewInstrument(instrumentData *[]asset.Instrument) error {
	result := ins.Db.Create(instrumentData)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
