package main

import (
	// "ModEd/eval/controller"
	model "ModEd/eval/model"
	"fmt"
	// "time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		return
	}

	db.AutoMigrate(&model.Examination{}, &model.Question{}, &model.Answer{})
}
