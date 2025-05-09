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
		}
	}

	fmt.Println("\nPress Enter to return...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
	menu.manager.SetState(menu.parent)
	return nil
}
