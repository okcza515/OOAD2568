package cli

import (
	"ModEd/core/cli"
	"ModEd/recruit/util"
	"bufio"
	"fmt"
	"os"
	"strings"
)

type AdminConfirmToStudentMenuState struct {
	manager      *cli.CLIMenuStateManager
	service      ConfirmedApplicantToStudentService
	parent       cli.MenuState
	reportViewer *AdminShowApplicationReportMenuState
}

func NewAdminConfirmToStudentMenuState(
	manager *cli.CLIMenuStateManager,
	service ConfirmedApplicantToStudentService,
	parent cli.MenuState,
	reportViewer *AdminShowApplicationReportMenuState,
) *AdminConfirmToStudentMenuState {
	return &AdminConfirmToStudentMenuState{
		manager:      manager,
		service:      service,
		parent:       parent,
		reportViewer: reportViewer,
	}
}

func (menu *AdminConfirmToStudentMenuState) Render() {
	util.ClearScreen()
	menu.reportViewer.viewReportsByStatusConfirmed()

}

func (menu *AdminConfirmToStudentMenuState) HandleUserInput(input string) error {
	fmt.Println("confirmed applicants to student [y/n]: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	buffer := scanner.Text()
	input = strings.TrimSpace(buffer)

	if input != "y" && input != "Y" {
		fmt.Println("Operation cancelled.")
	} else {
		err := menu.service.TransferConfirmedApplicants()
		if err != nil {
			fmt.Println("Error during transfer:", err)
		} else {
			fmt.Println("Transfer completed successfully.")
		}
	}
	util.WaitForEnter()
	menu.manager.SetState(menu.parent)
	return nil
}

func (menu *AdminShowApplicationReportMenuState) viewReportsByStatusConfirmed() {
	status := "Confirmed"
	reports, err := menu.service.GetApplicationReportsByStatus(status)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if len(reports) == 0 {
		fmt.Println("No reports found for status:", status)
		util.WaitForEnter()
		menu.manager.SetState(menu.parent)
		return
	}
	fmt.Printf("\n===== Application Reports (Status: %s) =====\n", status)
	menu.service.DisplayOnlyApplicationReport(reports)
}
