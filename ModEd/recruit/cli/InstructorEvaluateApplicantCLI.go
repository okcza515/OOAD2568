// MEP-1003 Student Recruitment
package cli

import (
	"ModEd/core/cli"
	recruitUtil "ModEd/recruit/util"
	"fmt"
	"strconv"
)

type InstructorEvaluateApplicantMenuState struct {
	manager         *cli.CLIMenuStateManager
	instructorID    uint
	evaluateService InstructorEvaluateApplicantService
	reportService   ApplicantReportService
	previousState   cli.MenuState
}

func NewInstructorEvaluateApplicantMenuState(
	manager *cli.CLIMenuStateManager,
	instructorID uint,
	evaluateService InstructorEvaluateApplicantService,
	reportService ApplicantReportService,
	previousState cli.MenuState,
) *InstructorEvaluateApplicantMenuState {
	return &InstructorEvaluateApplicantMenuState{
		manager:         manager,
		instructorID:    instructorID,
		evaluateService: evaluateService,
		reportService:   reportService,
		previousState:   previousState,
	}
}

func (m *InstructorEvaluateApplicantMenuState) Render() {
	recruitUtil.ClearScreen()
	fmt.Print("Enter Application Report ID: ")
}

func (m *InstructorEvaluateApplicantMenuState) HandleUserInput(input string) error {
	convReportID, err := strconv.ParseUint(input, 10, 32)
	if err != nil {
		fmt.Println("Invalid Application Report ID. Please enter a valid number.")
		recruitUtil.WaitForEnter()
		m.manager.SetState(m.previousState)
		return nil
	}
	applicationReportID := uint(convReportID)

	hasPermission, err := m.evaluateService.HasPermissionToEvaluate(m.instructorID, applicationReportID)
	if err != nil {
		fmt.Println("Error checking permission:", err)
		recruitUtil.WaitForEnter()
		m.manager.SetState(m.previousState)
		return nil
	}
	if !hasPermission {
		fmt.Println("You do not have permission to evaluate this application.")
		recruitUtil.WaitForEnter()
		m.manager.SetState(m.previousState)
		return nil
	}

	report, err := m.reportService.GetApplicationReport(applicationReportID)
	if err != nil {
		fmt.Println("Failed to fetch application report:", err)
		recruitUtil.WaitForEnter()
		m.manager.SetState(m.previousState)
		return nil
	}

	err = m.evaluateService.EvaluateApplicant(applicationReportID, report.ApplicationRound.RoundName, report.Faculty.Name, report.Department.Name)
	if err != nil {
		fmt.Println("Error updating interview score:", err)
	} else {
		fmt.Println("Score updated successfully!")
	}

	recruitUtil.WaitForEnter()
	m.manager.SetState(m.previousState)
	return nil
}
