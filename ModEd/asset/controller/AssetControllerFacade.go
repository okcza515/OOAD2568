package controller

import (
	model "ModEd/asset/model"
	"ModEd/core"
	"ModEd/utils/deserializer"
	"errors"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type AssetControllerFacade struct {
	db        *gorm.DB
	migration MigrationControllerInterface

	BorrowInstrument BorrowInstrumentControllerInterface
	Category         CategoryController
	Instrument       InstrumentController
	InstrumentLog    InstrumentLogControllerInterface
	Supply           SupplyControllerInterface
	SupplyLog        SupplyLogController
}

func CreateAssetControllerFacade() (*AssetControllerFacade, error) {
	database := "data/ModEd.bin"

	db, err := gorm.Open(sqlite.Open(database), &gorm.Config{})
	if err != nil {
		return nil, errors.New("err: failed to connect database")
	}

	facade := AssetControllerFacade{db: db}

	facade.migration = &MigrationController{db: db}
	facade.BorrowInstrument = &BorrowInstrumentController{db: db, BaseController: core.NewBaseController("BorrowInstrument", db)}
	facade.Category = CategoryController{db: db}
	facade.Instrument = InstrumentController{db: db}
	facade.InstrumentLog = &InstrumentLogController{db: db, BaseController: core.NewBaseController("InstrumentLog", db)}
	facade.Supply = &SupplyController{db: db, BaseController: core.NewBaseController("Supply", db)}
	facade.SupplyLog = SupplyLogController{db: db}

	err = facade.migration.migrateToDB()
	if err != nil {
		return nil, errors.New("err: failed to migrate schema")
	}

	return &facade, nil
}

func (facade *AssetControllerFacade) loadSeedData() error {
	seedData := map[string]interface{}{
		//"BorrowInstrumentList": &[]model.BorrowInstrument{},
		"Category":       &[]model.Category{},
		"InstrumentList": &[]model.Instrument{},
		//"InstrumentLog":  &[]model.InstrumentLog{},
		//"SupplyList":     &[]model.Supply{},
		//"SupplyLog":      &[]model.SupplyLog{},
	}

	for filename, m := range seedData {
		fd, err := deserializer.NewFileDeserializer("data/asset/" + filename + ".csv")
		if err != nil {
			return err
		}

		err = fd.Deserialize(m)
		if err != nil {
			return err
		}

		result := facade.db.Create(m)
		if result.Error != nil {
			return result.Error
		}
	}

	return nil
}

func (facade *AssetControllerFacade) ResetDB() error {
	err := facade.migration.resetDB()
	if err != nil {
		return err
	}

	return nil
}

func (facade *AssetControllerFacade) ResetAndLoadDB() error {
	err := facade.migration.resetDB()
	if err != nil {
		return err
	}

	err = facade.loadSeedData()
	if err != nil {
		return err
	}

	return nil
}
