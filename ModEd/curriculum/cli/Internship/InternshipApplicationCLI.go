package main

import (
	CommonModel "ModEd/common/model"
	controller "ModEd/curriculum/controller/Internship"
	model "ModEd/curriculum/model/Internship"

	"flag"
	"time"

	"github.com/google/uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	var (
		database string
		path     string
	)

	flag.StringVar(&database, "database", "C:/Users/bigza/Desktop/code/OOAD2568/ModEd/data/Intership/internship_data.bin", "Path of SQLite Database.") //TODO: Waiting for common database cli implemente
	flag.StringVar(&path, "path", "", "Path to CSV or JSON for WIL Application ")
	flag.Parse()

	connector, err := gorm.Open(sqlite.Open(database), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	InternStudent := &model.InternStudent{
		Student: CommonModel.Student{
			SID:       "65070501001",
			FirstName: "Mick",
			LastName:  "Doe",
			Email:     "mm@example.com",
			StartDate: time.Now(),
			BirthDate: time.Now(),
			Program:   CommonModel.ProgramType(1),
			Status:    CommonModel.StudentStatus(1),
		},
		InternID:     uuid.New(),
		InternStatus: model.ACTIVE,
	}

	controller := controller.CreateInternshipApplicationController(connector)

	controller.RegisterInternshipApplication(InternStudent)
}
