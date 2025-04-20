package examination

import (
	result_controller "ModEd/eval/controller/examination"
	"bufio"
	"os"
	"strings"

	"fmt"

	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func RunResultCLI() {
	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Connection failed:", err)
	}

	resultController := result_controller.NewResultController(db)

	for {
		fmt.Println("\nResult CLI")
		fmt.Println("1. Create Result")
		fmt.Println("2. Display All Results")
		// Display by student
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
			DisplayAllResults(db, resultController)

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

func DisplayAllResults(db *gorm.DB, resultController *result_controller.ResultController) {
	results, err := resultController.GetAllResults()
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
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter new score: ")
	scoreText, _ := reader.ReadString('\n')
	scoreText = strings.TrimSpace(scoreText)

	fmt.Print("Enter feedback: ")
	feedbackText, _ := reader.ReadString('\n')
	feedbackText = strings.TrimSpace(feedbackText)

	updatedData := map[string]interface{}{
		"Feedback": feedbackText,
		"Score": scoreText,
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
