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
	manager                            *cli.CLIMenuStateManager
	username                           string
	applicantController                *controller.ApplicantController
	adminShowApplicationReportsService AdminShowApplicationReportsService
	adminScheduleInterviewService      AdminScheduleInterviewService
	adminDeleteInterviewService        AdminInterviewService
	reportMenu                         cli.MenuState
	scheduleMenu                       cli.MenuState
	deleteMenu                         cli.MenuState
}

func NewAdminMenuState(
	manager *cli.CLIMenuStateManager,
	username string,
	applicantController *controller.ApplicantController,
	reportService AdminShowApplicationReportsService,
	scheduleInterviewService AdminScheduleInterviewService,
	deleteInterviewService AdminInterviewService,
) *AdminMenuState {
	menu := &AdminMenuState{
		manager:                            manager,
		username:                           username,
		applicantController:                applicantController,
		adminShowApplicationReportsService: reportService,
		adminScheduleInterviewService:      scheduleInterviewService,
		adminDeleteInterviewService:        deleteInterviewService,
	}
	menu.reportMenu = NewAdminShowApplicationReportMenuState(manager, reportService, menu)
	menu.scheduleMenu = NewAdminScheduleInterviewMenuState(manager, scheduleInterviewService, menu)
	menu.deleteMenu = NewAdminDeleteInterviewMenuState(manager, deleteInterviewService, menu)

	return menu
}

func (a *AdminMenuState) Render() {
	util.ClearScreen()
	fmt.Println("\n\033[1;35m╔════════════════════════════════╗")
	fmt.Printf("║ Welcome, Admin: %-16s ║\n", a.username)
	fmt.Println("╚════════════════════════════════╝\033[0m")

	fmt.Println("\n\033[1;36m[1]\033[0m Manage Applicants")
	fmt.Println("\033[1;36m[2]\033[0m View Application Reports")
	fmt.Println("\033[1;36m[3]\033[0m Schedule Interview")
	fmt.Println("\033[1;36m[4]\033[0m Delete Interview")
	fmt.Println("\033[1;36m[5]\033[0m Back")
	fmt.Print("\n\033[1;33mSelect an option:\033[0m ")
}

func (a *AdminMenuState) HandleUserInput(input string) error {
	switch input {
	case "1":
		ManageApplicants(a.applicantController)
		//util.WaitForEnter()
	case "2":
		a.manager.SetState(a.reportMenu)
	case "3":
		a.manager.SetState(a.scheduleMenu)
	case "4":
		a.manager.SetState(a.deleteMenu)
	case "5":
		return ErrExitAdminMenu
	default:
		fmt.Println("Invalid option. Try again.")
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
	menu := NewAdminMenuState(
		manager,
		username,
		dep.ApplicantController,
		dep.AdminShowApplicationReportsService,
		dep.AdminScheduleInterviewService,
		dep.AdminInterviewService,
	)
	manager.SetState(menu)

	for {
		manager.Render()
		var input string
		fmt.Scanln(&input)
		manager.UserInput = input

		err := manager.HandleUserInput()
		if err == ErrExitAdminMenu {
			break
		} else if err != nil {
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
