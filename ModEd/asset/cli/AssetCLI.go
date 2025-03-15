package main

import (
	controller "ModEd/asset/controller"
	model "ModEd/asset/model"
	util "ModEd/asset/util"

	"errors"
	"flag"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
	"time"
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

	err = instrumentLogController.MigrateToDB()
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

	logs, err := instrumentLogController.GetAll()
	if err != nil {
		panic("err: query failed")
	}

	for _, instrumentLog := range *logs {
		util.PrintStruct(instrumentLog)
	}
}
