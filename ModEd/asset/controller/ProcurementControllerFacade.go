// MEP-1014
package controller

import (
	"errors"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type ProcurementControllerFacade struct {
	db        *gorm.DB
	migration ProcurementMigrationController

	RequestedItem               InstrumentRequestController
	SupplierSelectionController SupplierSelectionController
	BudgetApproval              BudgetApprovalController
	Procurement                 ProcurementController
	TOR                         TORController
	Acceptance                  AcceptanceApprovalController
	// BudgetAllocation            BudgetAllocationController
	//ProcurementApproval ProcurementApprovalController
}

func CreateProcurementControllerFacade() (*ProcurementControllerFacade, error) {
	database := "data/ModEd.bin"

	db, err := gorm.Open(sqlite.Open(database), &gorm.Config{})
	if err != nil {
		return nil, errors.New("err: failed to connect database")
	}

	facade := ProcurementControllerFacade{db: db}

	facade.migration = ProcurementMigrationController{db: db}
	facade.RequestedItem = InstrumentRequestController{db: db}
	facade.SupplierSelectionController = SupplierSelectionController{db: db}
	facade.BudgetApproval = BudgetApprovalController{db: db}
	facade.Procurement = ProcurementController{db: db}
	facade.TOR = TORController{db: db}
	facade.Acceptance = AcceptanceApprovalController{db: db}
	// facade.BudgetAllocation = BudgetAllocationController{db: db}

	// fmt.Println("I'm In Facade yippie!")
	err = facade.migration.migrateToDB()
	if err != nil {
		return nil, errors.New("err: failed to migrate schema")
	}

	return &facade, nil
}

func (f *ProcurementControllerFacade) GetDB() *gorm.DB {
	return f.db
}
