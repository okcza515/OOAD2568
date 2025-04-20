package controller

import (
	model2 "ModEd/asset/model"
	"errors"

	"gorm.io/gorm"
)

type MigrationController struct {
	db *gorm.DB
}

func (c *MigrationController) migrateToDB() error {
	err := c.db.AutoMigrate(
		&model2.InstrumentLog{},
		&model2.Instrument{},
		&model2.BorrowInstrument{},
		&model2.Category{},
		&model2.Instrument{},
		&model2.InstrumentLog{},
		&model2.Supply{},
		&model2.SupplyLog{},
	)
	if err != nil {
		return errors.New("err: migration failed")
	}

	return nil
}

func (c *MigrationController) dropDB() error {
	err := c.db.Migrator().DropTable(
		&model2.InstrumentLog{},
		&model2.Instrument{},
		&model2.BorrowInstrument{},
		&model2.Category{},
		&model2.Instrument{},
		&model2.InstrumentLog{},
		&model2.Supply{},
		&model2.SupplyLog{},
	)
	if err != nil {
		return err
	}

	return nil
}

func (c *MigrationController) resetDB() error {
	err := c.dropDB()
	if err != nil {
		return err
	}

	err = c.migrateToDB()
	if err != nil {
		return err
	}

	return nil
}
