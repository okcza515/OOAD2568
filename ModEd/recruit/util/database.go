// MEP-1003 Student Recruitment
package util

import (
	"ModEd/recruit/model"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(path string) {
	var err error
	DB, err = gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	err = DB.AutoMigrate(&model.Admin{}, &model.Applicant{}, &model.ApplicationRound{}, &model.ApplicationReport{}, &model.Interview{})
	if err != nil {
		log.Fatalf("Failed to migrate models: %v\n", err)
	}

	log.Println("Database connected & migrated!")
}
