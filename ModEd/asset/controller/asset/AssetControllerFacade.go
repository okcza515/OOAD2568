package asset

import (
	"errors"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type AssetControllerFacade struct {
	Db *gorm.DB

	BorrowInstrument BorrowInstrumentController
	Category         CategoryController
	Instrument       InstrumentController
	InstrumentLog    InstrumentLogController
	Supply           SupplyController
	SupplyLog        SupplyLogController
}

func CreateAssetControllerFacade() (*AssetControllerFacade, error) {
	database := "data/ModEd.bin"

	db, err := gorm.Open(sqlite.Open(database), &gorm.Config{})
	if err != nil {
		return nil, errors.New("err: failed to connect database")
	}

	facade := AssetControllerFacade{Db: db}

	migrationController := MigrationController{Db: db}

	err = migrationController.MigrateToDB()
	if err != nil {
		return nil, errors.New("err: failed to migrate schema")
	}

	facade.BorrowInstrument = BorrowInstrumentController{Db: db}
	facade.Category = CategoryController{Db: db}
	facade.Instrument = InstrumentController{Db: db}
	facade.InstrumentLog = InstrumentLogController{Db: db}
	facade.Supply = SupplyController{Db: db}
	facade.SupplyLog = SupplyLogController{Db: db}

	return &facade, nil
}
