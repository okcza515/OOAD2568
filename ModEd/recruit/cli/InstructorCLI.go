// MEP-1003 Student Recruitment
package cli

import (
	"ModEd/recruit/util"
	"bufio"
	"fmt"
	"os"
)

func InstructorCLI(
	instructorViewInterviewDetailsService InstructorViewInterviewDetailsService,
	instructorEvaluateApplicantService InstructorEvaluateApplicantService) {
	var instructorID uint
	fmt.Print("Enter Instructor ID: ")
	fmt.Scanln(&instructorID)

	scanner := bufio.NewScanner(os.Stdin)

	for {
		util.ClearScreen()
		fmt.Println("\n==== Instructor Menu ====")
		fmt.Println("1. View Interview Details")
		fmt.Println("2. Evaluate an Applicant")
		fmt.Println("3. Exit")
		fmt.Print("Select an option: ")

		var choice int
		scanner.Scan()
		_, err := fmt.Sscan(scanner.Text(), &choice)
		if err != nil {
			fmt.Println("Invalid input, please try again.")
			continue
		}

		switch choice {
		case 1:
			ViewInterviewDetails(instructorViewInterviewDetailsService, instructorID)
			util.WaitForEnter()
		case 2:
			EvaluateApplicant(instructorEvaluateApplicantService)
		case 3:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid option. Try again.")
		}
	}
}
