package asset

import (
	"ModEd/asset/model/asset"
	"errors"

	"gorm.io/gorm"
)

type MigrationController struct {
	Db *gorm.DB
}

func (c *MigrationController) MigrateToDB() error {
	err := c.Db.AutoMigrate(
		&asset.InstrumentLog{},
		&asset.Instrument{},
		&asset.BorrowInstrument{},
		&asset.Category{},
		&asset.Instrument{},
		&asset.InstrumentLog{},
		&asset.Supply{},
		&asset.SupplyLog{},
	)
	if err != nil {
		return errors.New("err: migration failed")
	}

	return nil
}
