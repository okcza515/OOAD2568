// MEP-1003 Student Recruitment
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
		database          string
		roundsCSVPath     string
		facultyCSVPath    string
		departmentCSVPath string
		adminCSVPath      string
		role              string
	)

	curDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		return
	}
	parentDir := filepath.Dir(curDir)

	defaultDBPath := filepath.Join(parentDir, "data", "ModEd.bin")
	defaultRoundsPath := filepath.Join(parentDir, "recruit", "data", "application_rounds.csv")
	defaultfacultyPath := filepath.Join(parentDir, "recruit", "data", "fac.csv")
	defaultdepartmentPath := filepath.Join(parentDir, "recruit", "data", "dp.csv")
	defaultAdminPath := filepath.Join(parentDir, "recruit", "data", "AdminMockup.csv")

	flag.StringVar(&database, "database", defaultDBPath, "")
	flag.StringVar(&roundsCSVPath, "rounds", defaultRoundsPath, "")
	flag.StringVar(&facultyCSVPath, "faculty", defaultfacultyPath, "")
	flag.StringVar(&departmentCSVPath, "department", defaultdepartmentPath, "")
	flag.StringVar(&adminCSVPath, "admin", defaultAdminPath, "")
	flag.StringVar(&role, "role", "", "Specify the role (user/admin/instructor)")
	flag.Parse()

	db.InitDB(database)

	applicationReportCtrl := controller.CreateApplicationReportController(db.DB)
	applicantController := controller.NewApplicantController(db.DB)
	interviewController := controller.CreateInterviewController(db.DB)
	applicationRoundCtrl := controller.CreateApplicationRoundController(db.DB)

	adminCtrl := controller.CreateAdminController(db.DB)
	if err := adminCtrl.ReadAdminsFromCSV(defaultAdminPath); err != nil {
		fmt.Println(err)
	}

	facultyCtrl := common.CreateFacultyController(db.DB)

	departmentCtrl := common.CreateDepartmentController(db.DB)

	instructorController := controller.CreateInstructorController(db.DB)

	if err := applicationRoundCtrl.ReadApplicationRoundsFromCSV(roundsCSVPath); err != nil {
		fmt.Println("Failed to initialize application rounds:", err)
		return
	}

	for {
		util.ClearScreen()

		if role == "" {
			fmt.Println("\n\033[1;34m╔══════════════════════════════════════╗")
			fmt.Println("║       Moded Recruitment System       ║")
			fmt.Println("╚══════════════════════════════════════╝\033[0m")

			fmt.Println("\n\033[1;36m[1]\033[0m  User")
			fmt.Println("\033[1;36m[2]\033[0m  Admin")
			fmt.Println("\033[1;36m[3]\033[0m  Instructor")
			fmt.Println("\033[1;36m[4]\033[0m  Exit")
			fmt.Print("\n\033[1;33mSelect role:\033[0m ")

			var roleChoice int
			fmt.Scanln(&roleChoice)

			switch roleChoice {
			case 1:
				cli.UserCLI(applicantController, applicationRoundCtrl, applicationReportCtrl, facultyCtrl, departmentCtrl)
			case 2:
				cli.AdminCLI(applicantController, applicationReportCtrl, interviewController, adminCtrl)
			case 3:
				cli.InstructorCLI(instructorController)
			case 4:
				fmt.Println("Existing...")
				return
			default:
				fmt.Println("Invalid option. Try again.")
				continue
			}
		}

	}
}
