package main

import (
	controller "ModEd/curriculum/controller/Internship"

	"errors"
	"flag"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	var (
		database string
		path     string
	)

	flag.StringVar(&database, "database", "C:/Users/bigza/Desktop/code/OOAD2568/ModEd/data/ModEd.bin", "Path of SQLite Database.")
	flag.StringVar(&path, "path", "C:/Users/bigza/Desktop/code/OOAD2568/ModEd/data/Intership/Company.csv", "Path to CSV or JSON for student registration.")
	flag.Parse()

	db, err := gorm.Open(sqlite.Open(database), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		panic("*** Error: " + path + " does not exist.\n")
	}

	migrationController := controller.MigrationController{Db: db}

	err = migrationController.MigrateToDB()
	if err != nil {
		panic("err: migration failed")
	}

}
