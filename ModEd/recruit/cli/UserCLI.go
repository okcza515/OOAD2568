// MEP-1003 Student Recruitment
package cli

import (
	"ModEd/core/cli"
	"ModEd/recruit/util"
	"fmt"
)

type UserMenuState struct {
	manager *cli.CLIMenuStateManager
}

func NewUserMenuState(
	manager *cli.CLIMenuStateManager,
	applicantService ApplicantRegistrationService,
	reportService ApplicantReportService,
) *UserMenuState {
	userMenu := &UserMenuState{
		manager: manager,
	}
	manager.AddMenu("1", NewApplicantRegistrationMenuState(manager, applicantService, userMenu))
	manager.AddMenu("2", NewApplicantReportMenuState(manager, reportService, userMenu))
	//manager.AddMenu("3", nil)

	return userMenu

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

func (u *UserMenuState) HandleUserInput(input string) error {
	if input == "3" {
		return ErrExitMenu
	}

	err := u.manager.GoToMenu(input)
	if err != nil {
		fmt.Println("Invalid option. Try again.")
	}
	return nil

}

func UserCLI(applicantDeps ApplicantDependencies) {
	manager := cli.NewCLIMenuManager()
	userMenu := NewUserMenuState(manager, applicantDeps.ApplicantRegistrationService, applicantDeps.ApplicantReportService)

	manager.SetState(userMenu)

	for {
		manager.Render()

		var input string
		fmt.Scanln(&input)
		manager.UserInput = input

		err := manager.HandleUserInput()
		if err == ErrExitMenu {
			break
		} else if err != nil {
			fmt.Println("Error:", err)
		}
	}
}
