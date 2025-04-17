package asset

import (
	model "ModEd/asset/model/asset"
	"errors"

	"gorm.io/gorm"
)

type MigrationController struct {
	db *gorm.DB
}

func (c *MigrationController) migrateToDB() error {
	err := c.db.AutoMigrate(
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

func (c *MigrationController) dropDB() error {
	err := c.db.Migrator().DropTable(
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
