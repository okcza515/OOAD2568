package main

import (
	"ModEd/eval/controller"
	model "ModEd/eval/model"
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	// "time"
)

func main() {
	// exam := model.Examination{
	// 	ID:           1,
	// 	Exam_name:    "Final Exam - Programming",
	// 	Instructor_id: 101,
	// 	CourseId:     202,
	// 	CurriculumId: 303,
	// 	Criteria:     "Pass mark >= 50%",
	// 	Description:  "Final exam covering all topics in the course.",
	// 	Exam_date:    time.Date(2025, 5, 1, 9, 0, 0, 0, time.UTC),
	// }

	db, _ := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	db.AutoMigrate(&model.Examination{})

	coreController := controller.NewExaminationController(db)
	loggingCtrl := controller.NewLoggingDecorator(coreController)

	// loggingCtrl.CreateExam(&exam)
	fmt.Println(loggingCtrl.GetAll())
}
