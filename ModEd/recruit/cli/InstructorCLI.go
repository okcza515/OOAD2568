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
	instructorEvaluateApplicantService InstructorEvaluateApplicantService, applicantReportService ApplicantReportService, loginCtrl *controller.LoginController) {

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
		fmt.Println("\n\033[1;35m╔══════════════════════════════════════╗")
		fmt.Println("║           Instructor Menu            ║")
		fmt.Println("╚══════════════════════════════════════╝\033[0m")

		fmt.Println("\n\033[1;36m[1]\033[0m  View All Interview Details")
		fmt.Println("\033[1;36m[2]\033[0m  View Pending Interviews")
		fmt.Println("\033[1;36m[3]\033[0m  View Evaluated Interviews")
		fmt.Println("\033[1;36m[4]\033[0m  Evaluate an Applicant")
		fmt.Println("\033[1;36m[5]\033[0m  Exit")
		fmt.Print("\n\033[1;33mSelect an option:\033[0m ")
		var choice int
		scanner.Scan()
		_, err := fmt.Sscan(scanner.Text(), &choice)
		if err != nil {
			fmt.Println("Invalid input, please try again.")
			continue
		}

		switch choice {
		case 1:
			ViewInterviewDetails(instructorViewInterviewDetailsService, instructorIDUint, "all")
			util.WaitForEnter()
		case 2:
			ViewInterviewDetails(instructorViewInterviewDetailsService, instructorIDUint, "Pending")
			util.WaitForEnter()
		case 3:
			ViewInterviewDetails(instructorViewInterviewDetailsService, instructorIDUint, "Evaluated")
			util.WaitForEnter()
		case 4:
			EvaluateApplicant(instructorEvaluateApplicantService, applicantReportService, instructorIDUint)
		case 5:
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
