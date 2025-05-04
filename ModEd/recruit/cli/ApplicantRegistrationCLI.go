// MEP-1003 Student Recruitment
package cli

import (
	"ModEd/core/cli"
	"ModEd/recruit/util"
	"bufio"
	"fmt"
	"os"
)

type ApplicantRegistrationMenuState struct {
	manager          *cli.CLIMenuStateManager
	applicantService ApplicantRegistrationService
	parent           cli.MenuState
}

func NewApplicantRegistrationMenuState(
	manager *cli.CLIMenuStateManager,
	applicantService ApplicantRegistrationService,
	parent cli.MenuState,
) *ApplicantRegistrationMenuState {
	return &ApplicantRegistrationMenuState{
		manager:          manager,
		applicantService: applicantService,
		parent:           parent,
	}
}

func (menu *ApplicantRegistrationMenuState) Render() {
	util.ClearScreen()
	fmt.Println("==== Applicant Registration ====")
	fmt.Println("1. Register manually")
	fmt.Println("2. Register using CSV/JSON file")
	fmt.Println("3. Back")
	fmt.Print("Select option: ")
}

func (menu *ApplicantRegistrationMenuState) HandleUserInput(input string) error {
	scanner := bufio.NewScanner(os.Stdin)

	switch input {
	case "1":
		menu.applicantService.RegisterManually(scanner)
	case "2":
		menu.applicantService.RegisterFromFile(scanner)
	case "3":
		menu.manager.SetState(menu.parent)
	default:
		fmt.Println("Invalid option.")
	}

	fmt.Println("\nPress Enter to return...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
	return nil
}
