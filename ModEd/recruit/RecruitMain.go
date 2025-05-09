package main

import (
	common "ModEd/common/controller"
	"ModEd/recruit/cli"
	"ModEd/recruit/controller"
	"ModEd/recruit/util"
	db "ModEd/recruit/util"
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	var (
		database              string
		roundsCSVPath         string
		adminCSVPath          string
		interviewCreteriaPath string
		role                  string
	)

	curDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		return
	}
	parentDir := filepath.Dir(curDir)

	defaultDBPath := filepath.Join(parentDir, "data", "ModEd.bin")
	defaultRoundsPath := filepath.Join(parentDir, "recruit", "data", "application_rounds.csv")
	defaultAdminPath := filepath.Join(parentDir, "recruit", "data", "AdminMockup.csv")
	defaultInterviewCreteriaPath := filepath.Join(parentDir, "recruit", "data", "InterviewCriteria.csv")

	flag.StringVar(&database, "database", defaultDBPath, "")
	flag.StringVar(&roundsCSVPath, "rounds", defaultRoundsPath, "")
	flag.StringVar(&adminCSVPath, "admin", defaultAdminPath, "")
	flag.StringVar(&role, "role", "", "Specify the role (user/admin/instructor)")
	flag.StringVar(&interviewCreteriaPath, "criteria", defaultInterviewCreteriaPath, "")
	flag.Parse()

	db.InitDB(database)

	applicationReportCtrl := controller.NewApplicationReportController(db.DB)
	applicantController := controller.NewApplicantController(db.DB)
	interviewController := controller.NewInterviewController(db.DB)
	applicationRoundCtrl := controller.NewApplicationRoundController(db.DB)

	adminCtrl := controller.NewAdminController(db.DB)
	if err := adminCtrl.ReadAdminsFromCSV(defaultAdminPath); err != nil {
		fmt.Println(err)
	}

	facultyCtrl := common.NewFacultyController(db.DB)
	departmentCtrl := common.NewDepartmentController(db.DB)
	interviewCriteriaCtrl := controller.NewInterviewCriteriaCtrl(db.DB)

	instructorViewInterviewDetailsService := cli.NewInstructorViewInterviewDetailsService(db.DB, interviewController)
	instructorEvaluateApplicantService := cli.NewInstructorEvaluateApplicantService(
		db.DB,
		interviewCriteriaCtrl,
		applicationReportCtrl,
	)
	applicantRegistrationService := cli.NewApplicantRegistrationService(
		applicantController,
		applicationRoundCtrl,
		applicationReportCtrl,
		facultyCtrl,
		departmentCtrl,
	)
	applicantReportService := cli.NewApplicantReportService(db.DB, applicationReportCtrl)

	if err := applicationRoundCtrl.ReadApplicationRoundsFromCSV(roundsCSVPath); err != nil {
		fmt.Println("Failed to initialize application rounds:", err)
		return
	}
	if err := interviewCriteriaCtrl.ReadInterviewCriteriaFromCSV(interviewCreteriaPath); err != nil {
		fmt.Println("Failed to initialize interview criteria:", err)
		return
	}

	factory := &controller.DefaultLoginStrategyFactory{DB: db.DB}
	loginController := controller.LoginController{Strategy: factory.CreateStrategy(role)}

	// adminInterviewService := cli.NewAdminInterviewService(interviewController)
	adminDeps := cli.AdminDependencies{
		ApplicantController:                applicantController,
		ApplicationReportCtrl:              applicationReportCtrl,
		InterviewCtrl:                      interviewController,
		AdminCtrl:                          adminCtrl,
		LoginCtrl:                          &loginController,
		AdminInterviewService:              cli.NewAdminInterviewService(interviewController),
		AdminShowApplicationReportsService: cli.NewAdminShowApplicationReportsService(applicationReportCtrl, interviewController),
		AdminScheduleInterviewService:      cli.NewAdminScheduleInterviewService(interviewController, applicationReportCtrl),
	}

	for {
		util.ClearScreen()

		if role == "" {
			// Display the main menu
			fmt.Println("\n\033[1;34m╔══════════════════════════════════════╗")
			fmt.Println("║       Moded Recruitment System       ║")
			fmt.Println("╚══════════════════════════════════════╝\033[0m")

			// Options for role selection
			fmt.Println("\n\033[1;36m[1]\033[0m  User")
			fmt.Println("\033[1;36m[2]\033[0m  Admin")
			fmt.Println("\033[1;36m[3]\033[0m  Instructor")
			fmt.Println("\033[1;36m[4]\033[0m  Exit")
			fmt.Print("\n\033[1;33mSelect role:\033[0m ")

			// Get user input for role selection
			var roleChoice int
			fmt.Scanln(&roleChoice)

			switch roleChoice {
			case 1:
				loginController.SetStrategy(factory.CreateStrategy("user"))
				cli.UserCLI(applicantRegistrationService, applicantReportService)
			case 2:
				loginController.SetStrategy(factory.CreateStrategy("admin"))
				// cli.AdminCLI(applicantController, applicationReportCtrl, interviewController, adminCtrl, &loginController)
				cli.AdminCLI(adminDeps)
			case 3:
				loginController.SetStrategy(factory.CreateStrategy("instructor"))
				cli.InstructorCLI(instructorViewInterviewDetailsService, instructorEvaluateApplicantService, applicantReportService, &loginController, db.DB)

			case 4:
				fmt.Println("Exiting...")
				return

			default:
				fmt.Println("Invalid option. Try again.")
				continue
			}

		}
	}
}
