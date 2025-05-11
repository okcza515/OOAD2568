// MEP-1003 Student Recruitment
package cli

import (
	"ModEd/common/model"
	"ModEd/core/cli"
	"ModEd/recruit/controller"
	"ModEd/recruit/util"
	"fmt"
	"strconv"
)

type InstructorMenuState struct {
	manager                  *cli.CLIMenuStateManager
	instructorID             uint
	viewInterviewService     InstructorViewInterviewDetailsService
	evaluateApplicantService InstructorEvaluateApplicantService
	applicantReportService   ApplicantReportService
	interviewController      *controller.InterviewController
}

func NewInstructorMenuState(
	manager *cli.CLIMenuStateManager,
	instructorID uint,
	viewInterviewService InstructorViewInterviewDetailsService,
	evaluateApplicantService InstructorEvaluateApplicantService,
	applicantReportService ApplicantReportService,
	interviewController *controller.InterviewController,
) *InstructorMenuState {
	instructorMenu := &InstructorMenuState{
		manager:                  manager,
		instructorID:             instructorID,
		viewInterviewService:     viewInterviewService,
		evaluateApplicantService: evaluateApplicantService,
		applicantReportService:   applicantReportService,
		interviewController:      interviewController,
	}
	manager.AddMenu("1", NewInstructorViewInterviewDetailsMenuState(manager, instructorID, "all", viewInterviewService, instructorMenu, interviewController))
	manager.AddMenu("2", NewInstructorViewInterviewDetailsMenuState(manager, instructorID, "Pending", viewInterviewService, instructorMenu, interviewController))
	manager.AddMenu("3", NewInstructorViewInterviewDetailsMenuState(manager, instructorID, "Evaluated", viewInterviewService, instructorMenu, interviewController))
	manager.AddMenu("4", NewInstructorEvaluateApplicantMenuState(manager, instructorID, evaluateApplicantService, applicantReportService, instructorMenu))
	//manager.AddMenu("5", nil)

	return instructorMenu
}

func (m *InstructorMenuState) Render() {
	util.ClearScreen()
	fmt.Println("\n\033[1;35m╔══════════════════════════════════════╗")
	fmt.Println("║           Instructor Menu            ║")
	fmt.Println("╚══════════════════════════════════════╝\033[0m")

	fmt.Println("\n\033[1;36m[1]\033[0m  View All Interview Details")
	fmt.Println("\033[1;36m[2]\033[0m  View Pending Interviews")
	fmt.Println("\033[1;36m[3]\033[0m  View Evaluated Interviews")
	fmt.Println("\033[1;36m[4]\033[0m  Evaluate an Applicant")
	fmt.Println("\033[1;36m[5]\033[0m  Back")
	fmt.Print("\n\033[1;33mSelect an option:\033[0m ")
}

func (m *InstructorMenuState) HandleUserInput(input string) error {
	if input == "5" {
		return ErrExitMenu
	}

	err := m.manager.GoToMenu(input)
	if err != nil {
		fmt.Println("Invalid option. Try again.")
	}
	return err
}

func InstructorCLI(
	instructorDeps InstructorDependencies,
) {
	instructorID, err := promptInstructorCredentials()
	if err != nil {
		fmt.Println(err)
		return
	}

	instructorIDUint64, err := strconv.ParseUint(instructorID, 10, 32)
	if err != nil {
		fmt.Println("Invalid ID format")
		return
	}
	instructorIDUint := uint(instructorIDUint64)

	req := controller.LoginRequest{ID: instructorID}
	var instructor model.Instructor
	isValid, err := instructorDeps.LoginCtrl.ExecuteLogin(req, &instructor)
	if err != nil || !isValid {
		fmt.Println("Login failed:", err)
		return
	}

	manager := cli.NewCLIMenuManager()
	interviewCtrl := controller.NewInterviewController(instructorDeps.DB)
	menu := NewInstructorMenuState(manager, instructorIDUint, instructorDeps.ViewInterviewService, instructorDeps.EvaluateApplicantService, instructorDeps.ApplicantReportService, interviewCtrl)
	manager.SetState(menu)

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

func promptInstructorCredentials() (string, error) {
	var id string
	fmt.Print("Enter Instructor ID: ")
	fmt.Scanln(&id)

	if id == "" {
		return "", fmt.Errorf("Username and password are required")
	}

	return id, nil
}
