package main

import (
	controller "ModEd/curriculum/controller/wil-project"
	model "ModEd/curriculum/model/wil-project"
	"ModEd/utils/deserializer"
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

	flag.StringVar(&database, "database", "data/wil-project/WILProject.bin", "Path of SQLite Database.") //TODO: Waiting for common database cli implemente
	flag.StringVar(&path, "path", "", "Path to CSV or JSON for WIL Application ")
	flag.Parse()

	deserializer, err := deserializer.NewFileDeserializer(path)
	if err != nil {
		panic(err)
	}

	connector, err := gorm.Open(sqlite.Open(database), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		fmt.Printf("*** Error: %s does not exist.\n", path)
		return
	}

	controller := controller.CreateWILProjectApplicationController(connector)

	var applications []*model.WILProjectApplication
	if err := deserializer.Deserialize(&applications); err != nil {
		panic(err)
	}

	controller.RegisterWILProjectApplications(applications)

	retrieved, err := controller.GetAllWILProjectApplications()
	if err != nil {
		panic(err)
	}

	for _, application := range retrieved {
		fmt.Printf("%s %s %s %s\n", application.WILProjectApplicationId, application.ProjectName, application.ProjectDetail, application.Company)
	}
}
