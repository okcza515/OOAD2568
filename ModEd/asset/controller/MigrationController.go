package controller

import (
	"ModEd/asset/model"
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
		&model.ReturnInstrument{},
		&model.Supply{},
		&model.SupplyLog{},
	)
	if err != nil {
		return errors.New("err: migration failed")
	}

	return nil
}
