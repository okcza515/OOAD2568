// MEP-1003 Student Recruitment
package cli

import (
	"ModEd/common/model"
	"ModEd/recruit/controller"
	"ModEd/recruit/util"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func InstructorCLI(
	instructorViewInterviewDetailsService InstructorViewInterviewDetailsService,
	instructorEvaluateApplicantService InstructorEvaluateApplicantService, loginCtrl *controller.LoginController) {

	instructorID, err := promptInstructorCredentials()
	if err != nil {
		fmt.Println(err)
		time.Sleep(3 * time.Second)
		return
	}

	instructorIDUint64, err := strconv.ParseUint(instructorID, 10, 32)
	instructorIDUint := uint(instructorIDUint64)

	req := controller.LoginRequest{
		ID: instructorID,
	}

	var instructor model.Instructor
	isValid, err := loginCtrl.ExecuteLogin(req, &instructor)
	if err != nil {
		fmt.Println("An error occurred while checking credentials:", err)
		time.Sleep(3 * time.Second)
		return
	}
	if !isValid {
		fmt.Println("Invalid credentials. Access denied.")
		time.Sleep(3 * time.Second)
		return
	}

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
			ViewInterviewDetails(instructorViewInterviewDetailsService, instructorIDUint)
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

func promptInstructorCredentials() (string, error) {
	var id string
	fmt.Print("Enter Instructor ID: ")
	fmt.Scanln(&id)

	if id == "" {
		return "", fmt.Errorf("Username and password are required")
	}

	return id, nil
}
