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

	migrationController := MigrationController{db: db}

	err = migrationController.MigrateToDB()
	if err != nil {
		return nil, errors.New("err: failed to migrate schema")
	}

	facade.BorrowInstrument = BorrowInstrumentController{db: db}
	facade.Category = CategoryController{db: db}
	facade.Instrument = InstrumentController{db: db}
	facade.InstrumentLog = InstrumentLogController{db: db}
	facade.Supply = SupplyController{db: db}
	facade.SupplyLog = SupplyLogController{db: db}

	return &facade, nil
}
