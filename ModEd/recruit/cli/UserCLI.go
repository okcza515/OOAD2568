// MEP-1003 Student Recruitment
package cli

import (
	"ModEd/core/cli"
	"ModEd/recruit/util"
	"errors"
	"fmt"
)

var ErrExitUserMenu = errors.New("exit user menu")

type UserMenuState struct {
	manager          *cli.CLIMenuStateManager
	applicantService ApplicantRegistrationService
	reportService    ApplicantReportService
	registrationMenu cli.MenuState
	reportMenu       cli.MenuState
}

func NewUserMenuState(
	manager *cli.CLIMenuStateManager,
	applicantService ApplicantRegistrationService,
	reportService ApplicantReportService,
) *UserMenuState {
	menu := &UserMenuState{
		manager:          manager,
		applicantService: applicantService,
		reportService:    reportService,
	}
	menu.registrationMenu = NewApplicantRegistrationMenuState(manager, applicantService, menu)
	menu.reportMenu = NewApplicantReportMenuState(manager, reportService, menu)

	return menu

}

func (menu *UserMenuState) Render() {
	util.ClearScreen()

	fmt.Println("\n\033[1;35m╔════════════════════════════╗")
	fmt.Println("║          User Menu         ║")
	fmt.Println("╚════════════════════════════╝\033[0m")

	fmt.Println("\n\033[1;36m[1]\033[0m  Register Applicant")
	fmt.Println("\033[1;36m[2]\033[0m  View Application Status")
	fmt.Println("\033[1;36m[3]\033[0m  Back")
	fmt.Print("\n\033[1;33mSelect an option:\033[0m ")
}

func (menu *UserMenuState) HandleUserInput(input string) error {
	switch input {
	case "1":
		menu.manager.SetState(menu.registrationMenu)
	case "2":
		menu.manager.SetState(menu.reportMenu)
	case "3":
		fmt.Println("Exiting...")
		return ErrExitUserMenu
	default:
		fmt.Println("Invalid option.")
	}

	// fmt.Println("\nPress Enter to continue...")
	// bufio.NewReader(os.Stdin).ReadBytes('\n')
	return nil
}

func UserCLI(applicantRegistrationService ApplicantRegistrationService, applicantReportService ApplicantReportService) {
	manager := cli.NewCLIMenuManager()
	userMenu := NewUserMenuState(manager, applicantRegistrationService, applicantReportService)

	manager.SetState(userMenu)

	for {
		manager.Render()

		var input string
		fmt.Scanln(&input)
		manager.UserInput = input

		err := manager.HandleUserInput()
		if err == ErrExitUserMenu {
			return
		} else if err != nil {
			fmt.Println("Error:", err)
		}
	}
}
