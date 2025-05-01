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
	// Command-line flags for file paths and role
	var (
		database      string
		roundsCSVPath string
		adminCSVPath  string
		role          string
	)

	// Get the current working directory
	curDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		return
	}
	parentDir := filepath.Dir(curDir)

	// Default file paths for various resources
	defaultDBPath := filepath.Join(parentDir, "data", "ModEd.bin")
	defaultRoundsPath := filepath.Join(parentDir, "recruit", "data", "application_rounds.csv")
	defaultAdminPath := filepath.Join(parentDir, "recruit", "data", "AdminMockup.csv")

	// Parse command-line flags
	flag.StringVar(&database, "database", defaultDBPath, "")
	flag.StringVar(&roundsCSVPath, "rounds", defaultRoundsPath, "")
	flag.StringVar(&adminCSVPath, "admin", defaultAdminPath, "")
	flag.StringVar(&role, "role", "", "Specify the role (user/admin/instructor)")
	flag.Parse()

	// Initialize the database
	db.InitDB(database)

	// Create controllers
	applicationReportCtrl := controller.CreateApplicationReportController(db.DB)
	applicantController := controller.NewApplicantController(db.DB)
	interviewController := controller.CreateInterviewController(db.DB)
	applicationRoundCtrl := controller.CreateApplicationRoundController(db.DB)

	// Create admin controller and read admins from CSV
	adminCtrl := controller.CreateAdminController(db.DB)
	if err := adminCtrl.ReadAdminsFromCSV(defaultAdminPath); err != nil {
		fmt.Println(err)
	}

	// Other controllers
	facultyCtrl := common.CreateFacultyController(db.DB)
	departmentCtrl := common.CreateDepartmentController(db.DB)
	instructorViewInterviewDetailsService := cli.NewInstructorViewInterviewDetailsService(db.DB)
	instructorEvaluateApplicantService := cli.NewInstructorEvaluateApplicantService(db.DB)
	applicantRegistrationService := cli.NewApplicantRegistrationService(
		applicantController,
		applicationRoundCtrl,
		applicationReportCtrl,
		facultyCtrl,
		departmentCtrl,
	)
	applicantReportService := cli.NewApplicantReportService(db.DB)
	interviewService := cli.NewInterviewService(db.DB)

	// Initialize application rounds from CSV
	if err := applicationRoundCtrl.ReadApplicationRoundsFromCSV(roundsCSVPath); err != nil {
		fmt.Println("Failed to initialize application rounds:", err)
		return
	}

	loginController := controller.LoginController{
		Strategy: controller.NewLoginStrategy(role, db.DB),
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
				loginController.SetStrategy(controller.NewLoginStrategy("user", db.DB))
				cli.UserCLI(applicantRegistrationService, applicantReportService, interviewService, loginController)
			case 2:
				loginController.SetStrategy(controller.NewLoginStrategy("admin", db.DB))
				cli.AdminCLI(applicantController, applicationReportCtrl, interviewController, adminCtrl, &loginController)
			case 3:
				loginController.SetStrategy(controller.NewLoginStrategy("instructor", db.DB))
				cli.InstructorCLI(instructorViewInterviewDetailsService, instructorEvaluateApplicantService, applicantReportService, &loginController)
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
