package cli

import (
	"ModEd/core/cli"
	"ModEd/recruit/util"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type AdminShowApplicationReportMenuState struct {
	manager *cli.CLIMenuStateManager
	service AdminShowApplicationReportsService
	parent  cli.MenuState
}

func NewAdminShowApplicationReportMenuState(
	manager *cli.CLIMenuStateManager,
	service AdminShowApplicationReportsService,
	parent cli.MenuState,
) *AdminShowApplicationReportMenuState {
	return &AdminShowApplicationReportMenuState{
		manager: manager,
		service: service,
		parent:  parent,
	}
}

func (menu *AdminShowApplicationReportMenuState) Render() {
	util.ClearScreen()
	fmt.Print("Enter Applicant ID to view the report: ")
}

func (menu *AdminShowApplicationReportMenuState) HandleUserInput(input string) error {
	applicantID, err := strconv.ParseUint(input, 10, 32)
	if err != nil {
		fmt.Println("Invalid applicant ID:", err)
	} else {
		report, err := menu.service.GetApplicationReport(uint(applicantID))
		if err != nil {
			fmt.Println("Error retrieving report:", err)
		} else {
			fmt.Println("\n===== Applicant Report =====")
			fmt.Println(report)
		}
	}

	fmt.Println("\nPress Enter to return...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
	menu.manager.SetState(menu.parent)
	return nil
}
