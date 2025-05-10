// MEP-1003 Student Recruitment
package cli

import (
	"ModEd/core/cli"
	"ModEd/recruit/controller"
	"ModEd/recruit/model"
	"ModEd/recruit/util"
	"fmt"
	"time"
)

var ErrExitAdminMenu = fmt.Errorf("exit admin menu")

type AdminMenuState struct {
	manager  *cli.CLIMenuStateManager
	username string
}

func NewAdminMenuState(
	manager *cli.CLIMenuStateManager,
	username string,
	applicantController *controller.ApplicantController,
	reportService AdminShowApplicationReportsService,
	scheduleInterviewService AdminScheduleInterviewService,
	deleteInterviewService AdminInterviewService,
) *AdminMenuState {

	adminMenu := &AdminMenuState{
		manager:  manager,
		username: username,
	}

	manager.AddMenu("1", NewAdminShowApplicationReportMenuState(manager, reportService, adminMenu))
	manager.AddMenu("2", NewAdminScheduleInterviewMenuState(manager, scheduleInterviewService, adminMenu))
	manager.AddMenu("3", NewAdminDeleteInterviewMenuState(manager, deleteInterviewService, adminMenu))
	manager.AddMenu("4", nil)
	return adminMenu
}

func (a *AdminMenuState) Render() {
	util.ClearScreen()
	fmt.Println("\n\033[1;35m╔════════════════════════════════╗")
	fmt.Printf("║ Welcome, Admin: %-14s ║\n", a.username)
	fmt.Println("╚════════════════════════════════╝\033[0m")

	fmt.Println("\033[1;36m[1]\033[0m View Application Reports")
	fmt.Println("\033[1;36m[2]\033[0m Schedule Interview")
	fmt.Println("\033[1;36m[3]\033[0m Delete Interview")
	fmt.Println("\033[1;36m[4]\033[0m Back")
	fmt.Print("\n\033[1;33mSelect an option:\033[0m ")
}

func (a *AdminMenuState) HandleUserInput(input string) error {
	err := a.manager.GoToMenu(input)
	if err != nil {
		fmt.Printf("Invalid option. Try again.")
	}
	return nil
}

func AdminCLI(dep AdminDependencies) {
	username, err := AdminLogin(dep.LoginCtrl)
	if err != nil {
		fmt.Println(err)
		time.Sleep(3 * time.Second)
		return
	}

	manager := cli.NewCLIMenuManager()
	adminMenu := NewAdminMenuState(
		manager,
		username,
		dep.ApplicantController,
		dep.AdminShowApplicationReportsService,
		dep.AdminScheduleInterviewService,
		dep.AdminInterviewService,
	)
	manager.SetState(adminMenu)
	for {
		manager.Render()
		var input string
		fmt.Scanln(&input)
		manager.UserInput = input

		err := manager.HandleUserInput()
		if err != nil {
			fmt.Println("Error:", err)
		}
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
