package cli

import (
	"ModEd/recruit/controller"
	"bufio"
	"fmt"
	"os"
)

func EvaluateApplicant(instructorCtrl *controller.InstructorController) {
	var interviewID uint
	var score float64

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter Interview ID: ")
	scanner.Scan()
	_, err := fmt.Sscan(scanner.Text(), &interviewID)
	if err != nil {
		fmt.Println("Invalid Interview ID.")
		return
	}

	fmt.Print("Enter Interview Score: ")
	scanner.Scan()
	_, err = fmt.Sscan(scanner.Text(), &score)
	if err != nil {
		fmt.Println("Invalid score.")
		return
	}

	err = instructorCtrl.EvaluateApplicant(interviewID, score)
	if err != nil {
		fmt.Println("Error updating interview score:", err)
	} else {
		fmt.Println("Score updated successfully!")
	}
}
