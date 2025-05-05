// MEP-1003 Student Recruitment
package cli

import (
	"ModEd/recruit/controller"
	"ModEd/recruit/model"
	"ModEd/recruit/util"
	"fmt"
	"time"
)

func AdminCLI(dep AdminDependencies) {
	username, err := AdminLogin(dep.LoginCtrl)
	if err != nil {
		fmt.Println(err)
		time.Sleep(3 * time.Second)
		return
	}

	util.ClearScreen()
	fmt.Println("Login successful. Welcome,", username)
	for {
		fmt.Println("==== Admin Menu ====")
		fmt.Println("1. Manage Applicants")
		fmt.Println("2. View Application Reports")
		fmt.Println("3. Schedule Interview")
		fmt.Println("4. Delete Interview")
		fmt.Println("5. Back")
		fmt.Print("Select an option: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			ManageApplicants(dep.ApplicantController)
		case 2:
			// ShowApplicationReports(dep.ApplicationReportCtrl)
			AdminShowApplicationReportsCLI(dep.AdminShowApplicationReportsService)
			util.WaitForEnter()
		case 3:
			AdminScheduleInterviewCLI(dep.AdminScheduleInterviewService)
		case 4:
			AdminDeleteInterviewCLI(dep.AdminInterviewService)
			util.WaitForEnter()
		case 5:
			return 
		default:
			fmt.Println("Invalid option. Try again.")
			time.Sleep(1 * time.Second)
		}
		util.ClearScreen()
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

func ManageApplicants(applicantController *controller.ApplicantController) {
	fmt.Println("Managing Applicants...")
	util.WaitForEnter()
}
