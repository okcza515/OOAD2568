// MEP-1003 Student Recruitment
package cli

import (
	"ModEd/common/model"
	"ModEd/core/cli"
	"ModEd/recruit/controller"
	"ModEd/recruit/util"
	"fmt"
	"strconv"

	"gorm.io/gorm"
)

var ErrExitInstructorMenu = fmt.Errorf("exit instructor menu")

type InstructorMenuState struct {
	manager                  *cli.CLIMenuStateManager
	instructorID             uint
	viewInterviewService     InstructorViewInterviewDetailsService
	evaluateApplicantService InstructorEvaluateApplicantService
	applicantReportService   ApplicantReportService
	viewInterviewMenu        cli.MenuState
	evaluateApplicantMenu    cli.MenuState
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
	return &InstructorMenuState{
		manager:                  manager,
		instructorID:             instructorID,
		viewInterviewService:     viewInterviewService,
		evaluateApplicantService: evaluateApplicantService,
		applicantReportService:   applicantReportService,
		interviewController:      interviewController,
	}
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
	switch input {
	case "1":
		menu := NewInstructorViewInterviewDetailsMenuState(
			m.manager, m.instructorID, "all", m.viewInterviewService, m, m.interviewController)
		m.manager.SetState(menu)
	case "2":
		menu := NewInstructorViewInterviewDetailsMenuState(
			m.manager, m.instructorID, "Pending", m.viewInterviewService, m, m.interviewController)
		m.manager.SetState(menu)
	case "3":
		menu := NewInstructorViewInterviewDetailsMenuState(
			m.manager, m.instructorID, "Evaluated", m.viewInterviewService, m, m.interviewController)
		m.manager.SetState(menu)
	case "4":
		menu := NewInstructorEvaluateApplicantMenuState(
			m.manager, m.instructorID, m.evaluateApplicantService, m.applicantReportService, m)
		m.manager.SetState(menu)
	case "5":
		return ErrExitInstructorMenu
	default:
		fmt.Println("Invalid option. Try again.")
	}
	return nil
}

func InstructorCLI(
	viewInterviewService InstructorViewInterviewDetailsService,
	evaluateApplicantService InstructorEvaluateApplicantService,
	applicantReportService ApplicantReportService,
	loginCtrl *controller.LoginController,
	db *gorm.DB,
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
	isValid, err := loginCtrl.ExecuteLogin(req, &instructor)
	if err != nil || !isValid {
		fmt.Println("Login failed:", err)
		return
	}

	manager := cli.NewCLIMenuManager()
	interviewCtrl := controller.NewInterviewController(db)
	menu := NewInstructorMenuState(manager, instructorIDUint, viewInterviewService, evaluateApplicantService, applicantReportService, interviewCtrl)
	manager.SetState(menu)

	for {
		manager.Render()
		var input string
		fmt.Scanln(&input)
		manager.UserInput = input

		err := manager.HandleUserInput()
		if err == ErrExitInstructorMenu {
			return
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
