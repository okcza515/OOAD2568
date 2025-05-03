// MEP-1003 Student Recruitment
package cli

import (
	"ModEd/recruit/controller"
	"ModEd/recruit/model"
	"ModEd/recruit/util"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func AdminCLI(applicantController *controller.ApplicantController, applicationReportCtrl *controller.ApplicationReportController, interviewCtrl *controller.InterviewController, adminCtrl *controller.AdminController, loginCtrl *controller.LoginController) {
	username, err := AdminLogin(loginCtrl)
	if err != nil {
		fmt.Println(err)
		time.Sleep(3 * time.Second)
		return
	}

	util.ClearScreen()
	fmt.Println("Login successful. Welcome,", username)
	fmt.Println("==== Admin Menu ====")
	fmt.Println("1. Manage Applicants")
	fmt.Println("2. View Application Reports")
	fmt.Println("3. Schedule Interview")
	fmt.Println("4. Delete Interview")
	fmt.Println("5. back")
	fmt.Print("Select an option: ")

	var choice int
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		ManageApplicants(applicantController)
	case 2:
		ShowApplicationReports(applicationReportCtrl)
	case 3:
		AdminScheduleInterviewCLI(interviewCtrl, applicationReportCtrl)
	case 4:
		DeleteInterview(interviewCtrl)
	case 5:
		return
	default:
		fmt.Println("Invalid option. Try again.")

	}
}

func AdminLogin(loginCtrl *controller.LoginController) (string, error) {
	var username, password string
	fmt.Print("Enter admin username: ")
	fmt.Scanln(&username)
	fmt.Print("Enter admin password: ")
	fmt.Scanln(&password)

	req := controller.LoginRequest{
		Username: username,
		Password: password,
	}

	var admin model.Admin
	isValid, err := loginCtrl.ExecuteLogin(req, &admin)
	if err != nil {
		return "", fmt.Errorf("An error occurred while checking credentials: %v", err)
	}
	if !isValid {
		return "", fmt.Errorf("Invalid credentials. Access denied.")
	}

	return username, nil
}

func DeleteInterview(interviewCtrl *controller.InterviewController) {
	var interviewID uint

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter Interview ID to delete: ")
	scanner.Scan()
	inputID := scanner.Text()
	convInterviewID, err := strconv.ParseUint(inputID, 10, 32)
	if err != nil {
		fmt.Println("Invalid Interview ID. Please enter a valid number.")
		return
	}
	interviewID = uint(convInterviewID)

	err = interviewCtrl.DeleteInterview(interviewID)
	if err != nil {
		fmt.Println("Failed to delete interview:", err)
		return
	}

	fmt.Println("Interview deleted successfully!")
	util.WaitForEnter()
}

func ManageApplicants(applicantController *controller.ApplicantController) {
	fmt.Println("Managing Applicants...")
	util.WaitForEnter()
}


func ShowApplicationReports(applicationReportCtrl *controller.ApplicationReportController) {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter Applicant ID (number): ")
	scanner.Scan()
	inputID := scanner.Text()

	applicantID, err := strconv.ParseUint(inputID, 10, 32)
	if err != nil {
		fmt.Println("Invalid applicant ID. Please enter a valid number.")
		util.WaitForEnter()
		return
	}

	report, err := applicationReportCtrl.GetApplicationReportByApplicantID(uint(applicantID))
	if err != nil {
		fmt.Println("Error retrieving report:", err)
	} else {
		fmt.Println("===== Applicant Report =====")
		fmt.Println(report)
	}
	util.WaitForEnter()
}
