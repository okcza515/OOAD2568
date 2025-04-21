package examination

import (
	question_controller "ModEd/eval/controller/examination"
	"ModEd/eval/model"
	"ModEd/eval/util"
	"fmt"
	"log"
	"strconv"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func RunQuestionCLI() {
	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Connection failed:", err)
	}

	questionController := question_controller.NewQuestionController(db)

	for {
		fmt.Println("\nQuestion CLI")
		fmt.Println("1. Create Question")
		fmt.Println("2. Update Question")
		fmt.Println("3. Exit")
		fmt.Print("Enter your choice: ")
		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			CreateQuestion(questionController)
		case 2:
			var id uint
			fmt.Print("Enter Question ID to update: ")
			fmt.Scan(&id)
			UpdateQuestion(questionController, id)
		case 3:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice")
		}
	}

}

func promptQuestionType() (model.QuestionType, error) {
	fmt.Println("Enter Question type:")
	fmt.Println("1. Multiple Choice")
	fmt.Println("2. Short Answer")
	fmt.Println("3. True False")
	fmt.Println("4. Subjective")
	fmt.Print("Enter your choice: ")

	var choice int
	fmt.Scan(&choice)

	switch choice {
	case 1:
		return model.MULTIPLECHOICE, nil
	case 2:
		return model.SHORTANSWER, nil
	case 3:
		return model.TRUEFALSE, nil
	case 4:
		return model.SUBJECTIVE, nil
	default:
		return "", fmt.Errorf("invalid question type choice")
	}
}

func CreateQuestion(controller *question_controller.QuestionController) {
	examIdText := util.PromptString("Enter Exam ID: ")
	examIdUint, err := strconv.ParseUint(examIdText, 10, 64)
	if err != nil {
		fmt.Println("Invalid exam id input:", err)
		return
	}

	questionType, err := promptQuestionType()
	if err != nil {
		fmt.Println(err)
		return
	}

	detail := util.PromptString("Enter Question detail: ")
	answer := util.PromptString("Enter Correct answer: ")
	score, err := util.PromptFloat("Enter Score: ")
	if err != nil {
		fmt.Println("Invalid score input:", err)
		return
	}

	question := &model.Question{
		Exam_id:         uint(examIdUint),
		Question_type:   questionType,
		Question_detail: detail,
		Correct_answer:  answer,
		Score:           score,
	}

	if err := controller.CreateQuestion(question); err != nil {
		fmt.Println("Failed to create question:", err)
	} else {
		fmt.Println("Question created successfully!")
	}
}

func UpdateQuestion(controller *question_controller.QuestionController, id uint) {
	questionType, err := promptQuestionType()
	if err != nil {
		fmt.Println(err)
		return
	}

	detail := util.PromptString("Enter new Question detail: ")
	answer := util.PromptString("Enter new Correct answer: ")
	score, err := util.PromptFloat("Enter new Score: ")
	if err != nil {
		fmt.Println("Invalid score input:", err)
		return
	}

	updated := &model.Question{
		Question_type:   questionType,
		Question_detail: detail,
		Correct_answer:  answer,
		Score:           score,
	}

	if err := controller.UpdateQuestion(id, updated); err != nil {
		fmt.Println("Failed to update question:", err)
	} else {
		fmt.Println("Question updated successfully!")
	}
}
