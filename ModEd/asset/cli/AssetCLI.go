package main

import (
	controller "ModEd/asset/controller"
	model "ModEd/asset/model"
	util "ModEd/asset/util"

	"errors"
	"flag"
	"fmt"
	"os"

	"time"

	"github.com/google/uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	fmt.Println("Hi im asset")

	var (
		database string
		path     string
	)
	flag.StringVar(&database, "database", "data/ModEd.bin", "Path of SQLite Database.")
	flag.StringVar(&path, "path", "", "Path to CSV or JSON for student registration.")
	flag.Parse()

	db, err := gorm.Open(sqlite.Open(database), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		panic("*** Error: " + path + " does not exist.\n")
	}

	instrumentLogController := controller.InstrumentLogController{Db: db}
	supplyController := controller.SupplyController{Db: db}
	migrationController := controller.MigrationController{Db: db}

	err = migrationController.MigrateToDB()
	if err != nil {
		panic("err: migration failed")
	}

	util.PrintBanner()

	exampleLog := model.InstrumentLog{
		LogID:        uuid.New(),
		Timestamp:    time.Now(),
		RefUserID:    nil,
		StaffUserID:  uuid.New(),
		Action:       model.INS_LOG_REPAIR,
		InstrumentID: uuid.New(),
		Description:  "Do something",
		RefBorrowID:  nil,
	}

	db.Create(exampleLog)

	exampleSupply := model.Supply{
		SupplyID: uuid.New(),
		SupplyLabel: "New supply",
		Description: nil,
		RoomID: uuid.New(),
		Location: nil,
		CategoryID: nil,
		Quantity: 100,
		DeletedAt: gorm.DeletedAt{},
	}

	supplyController.Create(&exampleSupply)

	supplys, err := supplyController.GetAll()

	if err != nil {
		panic("err: query supply failed")
	}

	for _, supplyLog := range *supplys {
		util.PrintStruct(supplyLog)
	}

	supply, err := supplyController.GetByID((*supplys)[len((*supplys))-1].SupplyID)

	if err != nil {
		panic("err: query supply deleted")
	}

	util.PrintStruct(*supply)

	logs, err := instrumentLogController.GetAll()
	if err != nil {
		panic("err: query failed")
	}

	for _, instrumentLog := range *logs {
		util.PrintStruct(instrumentLog)
	}
}
