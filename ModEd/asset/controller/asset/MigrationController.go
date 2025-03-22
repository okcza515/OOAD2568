package asset

import (
	model "ModEd/asset/model/asset"
	"errors"

	"gorm.io/gorm"
)

type MigrationController struct {
	Db *gorm.DB
}

func (c *MigrationController) MigrateToDB() error {
	err := c.Db.AutoMigrate(
		&model.InstrumentLog{},
		&model.Instrument{},
		&model.BorrowInstrument{},
		&model.Category{},
		&model.Instrument{},
		&model.InstrumentLog{},
		&model.Supply{},
		&model.SupplyLog{},
	)
	if err != nil {
		return errors.New("err: migration failed")
	}

	return nil
}
