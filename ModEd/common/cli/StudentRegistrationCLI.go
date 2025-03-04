package main

import (
	"ModEd/common/controller"
	"ModEd/common/util"
	"errors"
	"flag"
	"fmt"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	var (
		database string
		path     string
	)
	flag.StringVar(&database, "database", "data/ModEd.bin", "Path of SQLite Database.")
	flag.StringVar(&path, "path", "", "Path to CSV or JSON for student registration.")
	flag.Parse()

	connector, err := gorm.Open(sqlite.Open(database), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		fmt.Printf("*** Error: %s does not exist.\n", path)
		return
	}

	controller := controller.CreateStudentRegistration(connector)

	mapper, err := util.CreateMapper(path)
	if err != nil {
		panic(err)
	}

	students := mapper.Map()
	controller.Register(students)

	retrieved, err := controller.GetAll()
	if err != nil {
		panic(err)
	}
	for _, student := range retrieved {
		fmt.Printf("%s %s %s\n", student.SID, student.FirstName, student.LastName)
	}
}
