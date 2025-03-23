package main

import (
	controller "ModEd/asset/controller/Procurement"
	"ModEd/asset/util"

	// "ModEd/asset/util"
	"flag"
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	fmt.Println("Starting Procurement Module")

	var database string
	flag.StringVar(&database, "database", "data/ModEd_Procurement.bin", "Path of SQLite Database.")
	flag.Parse()

	db, err := gorm.Open(sqlite.Open(database), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	procurementController := controller.CreateProcurementController(db)

	util.PrintBanner()

	// exampleProcurement := model.Procurement{
	// 	ProcurementApprovalWorkflowID: uuid.New(),
	// 	RequestedByUserID:             uuid.New().String(),
	// 	ApprovedByUserID:              nil,
	// 	Status:                        "Pending",
	// 	CreatedAt:                     time.Now(),
	// 	UpdatedAt:                     time.Now(),
	// }

	err = procurementController.Create(&exampleProcurement)
	if err != nil {
		panic("Error: Failed to create procurement entry")
	}

	procurements, err := procurementController.ListAll()
	if err != nil {
		panic("Error: Failed to retrieve procurement list")
	}

	for _, procurement := range procurements {
		util.PrintStruct(procurement)
	}

	if len(procurements) > 0 {
		latestID := procurements[len(procurements)-1].ProcurementApprovalWorkflowID
		procurement, err := procurementController.GetByApprovalId(latestID)
		if err != nil {
			panic("Error: Failed to fetch procurement by ID")
		}
		util.PrintStruct(*procurement)
	}
}
