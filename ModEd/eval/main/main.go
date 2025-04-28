package main

import (
	"ModEd/eval/controller"
	model "ModEd/eval/model"
	"fmt"
	"time"

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

	examFacade := controller.NewExaminationFacade(db)

	loggingCtrl := controller.NewLoggingDecorator(examFacade)

	exam := model.NewExaminationBuilder().
		Build(
			model.WithExamName("Final Exam - Programming"),
			model.WithInstructorID(101),
			model.WithCourseId(202),
			model.WithCurriculumId(303),
			model.WithCriteria("Pass mark >= 50%"),
			model.WithDescription("Final exam covering all topics in the course."),
			model.WithExamDate(time.Date(2025, 5, 1, 9, 0, 0, 0, time.UTC)),
		)

	question := model.NewQuestionBuilder().
		ID(1).
		ExamID(2).
		Detail("ข้อสอบข้อที่หนึ่ง").
		Score(10).
		Build()

	loggingCtrl.CreateExamAndQuestion(exam, question)
}
