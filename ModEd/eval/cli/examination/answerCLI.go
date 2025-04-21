package examination

import (
	answer_controller "ModEd/eval/controller/examination"
	"ModEd/eval/util"
	// "bufio"
	// "os"
	// "strings"

	"fmt"

	// "log"

	// "gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func RunAnswerCLI(db *gorm.DB) {

	answerController := answer_controller.NewAnswerController(db)

	for {
		fmt.Println("\nAnswer CLI")
		fmt.Println("1. Display All Answers by QuestionID")
		fmt.Println("2. Display Answer specific Student in specific Question")
		fmt.Println("3. Create Answer by Student")
		fmt.Println("4. Update Answer by Student")
		fmt.Println("5. Delete Answer")
		fmt.Println("6. Exit")
		fmt.Print("Enter your choice: ")
		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			var questionID uint
			fmt.Print("Enter Question ID: ")
			fmt.Scan(&questionID)
			DisplayAnswersByQuestionID(db, answerController, questionID)

		case 2:
			var studentID, questionID uint
			fmt.Print("Enter Student ID: ")
			fmt.Scan(&studentID)
			fmt.Print("Enter Question ID: ")
			fmt.Scan(&questionID)
			DisplayAnswerByStudentQuestion(db, answerController, studentID, questionID)

		case 3:
			var studentID, questionID uint
			fmt.Print("Enter Student ID: ")
			fmt.Scan(&studentID)
			fmt.Print("Enter Question ID: ")
			fmt.Scan(&questionID)
			CreateAnswer(db, answerController, studentID, questionID)

		case 4:
			var answerID uint
			fmt.Print("Enter Answer ID to update: ")
			fmt.Scan(&answerID)
			UpdateAnswer(db, answerController, answerID)

		case 5:
			var answerID uint
			fmt.Print("Enter Answer ID to delete: ")
			fmt.Scan(&answerID)
			DeleteAnswer(db, answerController, answerID)

		case 6:
			fmt.Println("Exiting...")
			return

		default:
			fmt.Println("Invalid choice")
		}
	}
}

func DisplayAnswersByQuestionID(db *gorm.DB, answerController *answer_controller.AnswerController, questionID uint) {
	answers, err := answerController.GetAnswersByQuestion(questionID)
	if err != nil {
		fmt.Println("Error fetching answers for Question ID:", err)
		return
	}

	if len(answers) == 0 {
		fmt.Println("No answers found for this question.")
		return
	}

	fmt.Printf("\nAnswers for Question ID: %d\n", questionID)
	for _, answer := range answers {
		fmt.Printf("Answer ID: %d | Student ID: %d | Student Name: %s | Answer: %s\n",
			answer.ID, answer.StudentID, answer.Student.FirstName+" "+answer.Student.LastName, answer.Answer)
	}
}

func DisplayAnswerByStudentQuestion(db *gorm.DB, answerController *answer_controller.AnswerController, studentID uint, questionID uint) {
	answer, err := answerController.GetAnswerByQuestionAndStudent(questionID, studentID)
	if err != nil {
		fmt.Println("Error fetching answer for Student ID:", studentID, "and Question ID:", questionID, err)
		return
	}

	fmt.Printf("\nAnswer for Student ID: %d and Question ID: %d\n", studentID, questionID)
	fmt.Printf("Answer ID: %d | Student Name: %s | Answer: %s\n",
		answer.ID, answer.Student.FirstName+" "+answer.Student.LastName, answer.Answer)
}

func CreateAnswer(db *gorm.DB, answerController *answer_controller.AnswerController, studentID, questionID uint) {
	answerText := util.PromptString("Enter your Answer: ")

	err := answerController.SubmitAnswer(questionID, studentID, answerText)
	if err != nil {
		fmt.Println("Failed to create answer:", err)
	} else {
		fmt.Println("Answer created successfully!")
	}
}

func UpdateAnswer(db *gorm.DB, answerController *answer_controller.AnswerController, answerID uint) {
	answerText := util.PromptString("Enter new answer: ")

	updatedData := map[string]interface{}{
		"Answer": answerText,
	}

	err := answerController.UpdateAnswerByID(answerID, updatedData)
	if err != nil {
		fmt.Println("Failed to update answer:", err)
	} else {
		fmt.Println("Answer updated successfully!")
	}
}

func DeleteAnswer(db *gorm.DB, answerController *answer_controller.AnswerController, answerID uint) {
	err := answerController.DeleteAnswerByID(answerID)
	if err != nil {
		fmt.Println("Failed to delete answer:", err)
	} else {
		fmt.Println("Answer deleted successfully!")
	}
}
