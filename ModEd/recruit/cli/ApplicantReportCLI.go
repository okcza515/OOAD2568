// MEP-1003 Student Recruitment
package cli

import (
	"ModEd/core/cli"
	"ModEd/recruit/model"
	"ModEd/recruit/util"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ApplicantReportMenuState struct {
	manager       *cli.CLIMenuStateManager
	reportService ApplicantReportService
	parent        cli.MenuState
}

func NewApplicantReportMenuState(
	manager *cli.CLIMenuStateManager,
	reportService ApplicantReportService,
	parent cli.MenuState,
) *ApplicantReportMenuState {
	return &ApplicantReportMenuState{
		manager:       manager,
		reportService: reportService,
		parent:        parent,
	}
}

func (menu *ApplicantReportMenuState) Render() {
	util.ClearScreen()
	fmt.Print("Enter Applicantion Report ID to view the report: ")
}

func (menu *ApplicantReportMenuState) HandleUserInput(input string) error {
	applicantReportID, err := strconv.ParseUint(input, 10, 32)
	if err != nil {
		fmt.Println("Invalid Applicantion Report ID :", err)
	} else {
		report, err := menu.reportService.GetApplicationReport(uint(applicantReportID))
		if err != nil {
			fmt.Println("Error retrieving report:", err)
		} else {
			menu.reportService.DisplayReport([]*model.ApplicationReport{report})

			if report.ApplicationStatuses == "Accepted" {
				fmt.Print("\nThis applicant has been accepted. Do you want to confirm acceptance? [y/n]: ")
				reader := bufio.NewReader(os.Stdin)
				confirmInput, _ := reader.ReadString('\n')
				confirmInput = strings.TrimSpace(confirmInput)

				if confirmInput == "y" || confirmInput == "Y" {
					err := menu.reportService.ConfirmAcceptance(report.ApplicationReportID, model.Confirmed)
					if err != nil {
						fmt.Println("Could not confirm acceptance:", err)
					} else {
						fmt.Println("Acceptance confirmed successfully.")
					}
				} else if confirmInput == "n" || confirmInput == "N" {
					err = menu.reportService.ConfirmAcceptance(report.ApplicationReportID, model.Withdrawn)
					if err != nil {
						fmt.Println("Could not cancel acceptance:", err)
					} else {
						fmt.Println("Acceptance cancelled successfully.")
					}
				} else {
					fmt.Println("Invalid input. Please enter 'y' or 'n'.")
				}
			}
		}
	}

	fmt.Println("\nPress Enter to return...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
	menu.manager.SetState(menu.parent)
	return nil
}
