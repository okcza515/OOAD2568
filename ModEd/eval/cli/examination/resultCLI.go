package examination

import (
	result_controller "ModEd/eval/controller"
	"ModEd/eval/util"
	// "bufio"
	// "os"
	// "strings"

	"fmt"

	// "log"

	// "gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func RunResultCLI(db *gorm.DB) {

	resultController := result_controller.NewResultController(db)

	for {
		fmt.Println("\nResult CLI")
		fmt.Println("1. Create Result")
		fmt.Println("2. Display All Results by StudentID")
		fmt.Println("3. Update Result by ResultID")
		fmt.Println("4. Delete Result by ResultID")
		fmt.Println("5. Exit")
		fmt.Print("Enter your choice: ")
		var choice int
		fmt.Scan(&choice)

		switch choice {
		// case 1:
		// 	CreateResults(db, resultController)

		case 2:
			var studentID uint
			fmt.Print("Enter Student ID: ")
			fmt.Scan(&studentID)
			DisplayResultsByStudentID(db, resultController, studentID)

		case 3:
			var resultID uint
			fmt.Print("Enter Result ID: ")
			fmt.Scan(&resultID)
			UpdateResult(db, resultController, resultID)

		case 4:
			var resultID uint
			fmt.Print("Enter Result ID: ")
			fmt.Scan(&resultID)
			DeleteResult(db, resultController, resultID)

		case 5:
			fmt.Println("Exiting...")
			return

		default:
			fmt.Println("Invalid choice")
		}
	}
}

// func CreateResults(db *gorm.DB, resultController *result_controller.ResultController) {
// 	err := resultController.CreateResults()
// 	if err != nil {
// 		fmt.Println("Failed to create results:", err)
// 	} else {
// 		fmt.Println("Results created successfully!")
// 	}
// }

func DisplayResultsByStudentID(db *gorm.DB, resultController *result_controller.ResultController, studentID uint) {
	results, err := resultController.GetResultByStudent(studentID)
	if err != nil {
		fmt.Println("Error fetching results", err)
		return
	}

	if len(results) == 0 {
		fmt.Println("No answers found for this question.")
		return
	}

	fmt.Printf("\nAll Results")
	for _, result := range results {
		fmt.Printf("Result ID: %d | Examination ID: %d | Student ID: %d | Status: %s | Feedback: %s | Score: %d\n",
			result.ID, result.ExaminationID, result.StudentID, result.Status, result.Feedback, result.Score)
	}
}

func UpdateResult(db *gorm.DB, resultController *result_controller.ResultController, resultID uint) {
	scoreText := util.PromptString("Enter new score: ")
	feedbackText := util.PromptString("Enter feedback: ")

	updatedData := map[string]interface{}{
		"Feedback": feedbackText,
		"Score":    scoreText,
	}

	err := resultController.UpdateResult(resultID, updatedData)
	if err != nil {
		fmt.Println("Failed to update result:", err)
	} else {
		fmt.Println("Result updated successfully!")
	}
}

func DeleteResult(db *gorm.DB, answerController *result_controller.ResultController, resultID uint) {
	err := answerController.DeleteResult(resultID)
	if err != nil {
		fmt.Println("Failed to delete result:", err)
	} else {
		fmt.Println("Result deleted successfully!")
	}
}
