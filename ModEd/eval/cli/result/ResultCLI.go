package result

import (
	result_controller "ModEd/eval/controller"
	"ModEd/eval/util"
	"strconv"

	"fmt"

	"gorm.io/gorm"
)

func RunResultCLI(db *gorm.DB) {

	resultController := result_controller.NewResultController(db)

	for {
		fmt.Println("\nResult CLI")
		fmt.Println("1. Create Result by ExamID")
		fmt.Println("2. Display Result by ExamID and StudentID")
		fmt.Println("3. Update Result by ResultID")
		fmt.Println("4. Delete Result by ResultID")
		fmt.Println("5. Exit")
		fmt.Print("Enter your choice: ")
		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			var examID uint
			fmt.Print("Enter Examination ID: ")
			fmt.Scan(&examID)
			CreateResultByExamID(db, resultController, examID)

		case 2:
			var examID, studentID uint
			fmt.Print("Enter Examination ID: ")
			fmt.Scan(&examID)
			fmt.Print("Enter Student ID: ")
			fmt.Scan(&studentID)
			DisplayResultByExamStudent(db, resultController, examID, studentID)

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

func CreateResultByExamID(db *gorm.DB, resultController *result_controller.ResultController, examID uint) {
	err := resultController.CreateResultByExamID(examID)
	if err != nil {
		fmt.Println("Failed to create results:", err)
	} else {
		fmt.Println("Results created successfully!")
	}
}

func DisplayResultByExamStudent(db *gorm.DB, resultController *result_controller.ResultController, examID uint, studentID uint) {
	results, err := resultController.GetResultByExamAndStudent(examID, studentID)
	if err != nil {
		fmt.Println("Error fetching results", err)
		return
	}

	if len(results) == 0 {
		fmt.Println("No results found for this student.")
		return
	}

	fmt.Printf("\nAll Results:\n")
	for _, result := range results {
		fmt.Printf("Result ID: %d | Examination ID: %d | Student ID: %d | Status: %s | Feedback: %s | Score: %d\n",
			result.ID, result.ExaminationID, result.StudentID, result.Status, result.Feedback, result.Score)
	}
}

func UpdateResult(db *gorm.DB, resultController *result_controller.ResultController, resultID uint) {
	feedbackText := util.PromptString("Enter feedback: ")
	scoreText := util.PromptString("Enter new score: ")
	scoreValue, err := strconv.Atoi(scoreText)
	if err != nil {
		fmt.Println("Invalid score. Please enter a valid number.")
		return
	}

	updatedData := map[string]interface{}{
		"Feedback": feedbackText,
		"Score":    scoreValue,
	}

	err = resultController.UpdateResult(resultID, updatedData)
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
