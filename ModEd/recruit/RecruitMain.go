// MEP-1003 Student Recruitment
package main

import (
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

	flag.StringVar(&database, "database", defaultDBPath, "")
	flag.StringVar(&roundsCSVPath, "rounds", defaultRoundsPath, "")
	flag.StringVar(&facultyCSVPath, "faculty", defaultfacultyPath, "")
	flag.StringVar(&departmentCSVPath, "department", defaultdepartmentPath, "")
	flag.StringVar(&role, "role", "", "Specify the role (user/admin/instructor)")
	flag.Parse()

	db.InitDB(database)

	applicationReportCtrl := controller.CreateApplicationReportController(db.DB)
	applicantController := controller.NewApplicantController(db.DB)
	interviewController := controller.CreateInterviewController(db.DB)
	applicationRoundCtrl := controller.CreateApplicationRoundController(db.DB)

	facultyCtrl := controller.NewFacultyController(db.DB)
	if err := facultyCtrl.ReadFacultyFromCSV(facultyCSVPath); err != nil {
		fmt.Println(err)
	}

	departmentCtrl := controller.NewDepartmentController(db.DB)
	if err := departmentCtrl.ReadDepartmentFromCSV(departmentCSVPath); err != nil {
		fmt.Println(err)
	}

	instructorController := controller.CreateInstructorController(db.DB)

	if err := applicationRoundCtrl.ReadApplicationRoundsFromCSV(roundsCSVPath); err != nil {
		fmt.Println("Failed to initialize application rounds:", err)
		return
	}

	for {
		util.ClearScreen()

		if role == "" {
			fmt.Println("\n==== Student Recruitment System ====")
			fmt.Println("1. User")
			fmt.Println("2. Admin")
			fmt.Println("3. Instructor")
			fmt.Println("4. Exit")
			fmt.Print("Select role: ")

			var roleChoice int
			fmt.Scanln(&roleChoice)

			switch roleChoice {
			case 1:
				cli.UserCLI(applicantController, applicationRoundCtrl, applicationReportCtrl, facultyCtrl, departmentCtrl)
			case 2:
				cli.AdminCLI(applicantController, applicationReportCtrl, interviewController)
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
