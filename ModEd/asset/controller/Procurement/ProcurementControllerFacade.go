package Procurement

import (
	"errors"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type ProcurementControllerFacade struct {
	db *gorm.DB

	migration     MigrationController
	requestedItem ItemRequestController
}

func CreateProcurementControllerFacade() (*ProcurementControllerFacade, error) {
	database := "data/ModEd.bin"

	db, err := gorm.Open(sqlite.Open(database), &gorm.Config{})
	if err != nil {
		return nil, errors.New("err: failed to connect database")
	}

	facade := ProcurementControllerFacade{db: db}

	facade.requestedItem = ItemRequestController{}

	err = facade.migration.MigrateToDB()
	if err != nil {
		return nil, errors.New("err: failed to migrate schema")
	}

	return &facade, nil
}
