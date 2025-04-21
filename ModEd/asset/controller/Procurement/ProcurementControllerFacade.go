// MEP-1014
package procurement

import (
	"errors"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type ProcurementControllerFacade struct {
	db        *gorm.DB
	migration MigrationController

	RequestedItem               InstrumentRequestController
	SupplierSelectionController SupplierSelectionController
}

func CreateProcurementControllerFacade() (*ProcurementControllerFacade, error) {
	database := "data/ModEd.bin"

	db, err := gorm.Open(sqlite.Open(database), &gorm.Config{})
	if err != nil {
		return nil, errors.New("err: failed to connect database")
	}

	facade := ProcurementControllerFacade{db: db}

	facade.migration = MigrationController{db: db}
	facade.RequestedItem = InstrumentRequestController{db: db}
	facade.SupplierSelectionController = SupplierSelectionController{db: db}

	// fmt.Println("I'm In Facade yippie!")
	err = facade.migration.migrateToDB()
	if err != nil {
		return nil, errors.New("err: failed to migrate schema")
	}

	return &facade, nil
}
