package main

import (
	"ModEd/eval/cli"
	"ModEd/eval/controller"
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("quiz.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}

	err = db.AutoMigrate(&controller.QuizInput{})
	if err != nil {
		fmt.Println("Migration error:", err)
	}

	quizController := controller.NewQuizController(db)
	app := cli.NewQuizCLI(quizController)
	app.Run()
}
