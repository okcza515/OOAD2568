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

	flag.StringVar(&database, "database", "", "Path of SQLite Database.")
	flag.StringVar(&path, "path", "", "Path to CSV or JSON for student registration.")
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

	companyDataController := controller.NewCompanyDataController(db)
	err = companyDataController.ImportCompaniesFromCSV("C:/Users/bigza/Desktop/code/OOAD2568/ModEd/data/Intership/Company.csv")
	if err != nil {
		panic("err: failed to import companies")
	}

	internStudentController := controller.InternStudentController{Connector: db}
	err = internStudentController.RegisterInternStudentsFromFile("C:/Users/bigza/Desktop/code/OOAD2568/ModEd/data/StudentList.csv")
	if err != nil {
		panic("err: failed to import students")
	}

}
