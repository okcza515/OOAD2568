// MEP-1003 Student Recruitment
package main

import (
	"ModEd/recruit/cli"
	"ModEd/recruit/controller"
	db "ModEd/recruit/util"
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"runtime"

	"github.com/google/uuid"
)

func clearScreen() {
	switch runtime.GOOS {
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls") // Windows clear screen command
		cmd.Stdout = os.Stdout
		cmd.Run()
	default:
		cmd := exec.Command("clear") // Linux/macOS clear screen command
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func main() {
	var (
		database          string
		roundsCSVPath     string
		facultyCSVPath    string
		departmentCSVPath string
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
	flag.Parse()

	db.InitDB(database)

	applicationReportCtrl := controller.CreateApplicationReportController(db.DB)
	applicantController := controller.CreateApplicantController(db.DB)
	applicationRoundCtrl := controller.CreateApplicationRoundController(db.DB)

	//temp
	facultyCtrl := controller.NewFacultyController(db.DB)
	if err := facultyCtrl.ReadFacultyFromCSV(facultyCSVPath); err != nil {
		fmt.Println(err)
	}

	departmentCtrl := controller.NewDepartmentController(db.DB)

	if err := departmentCtrl.ReadDepartmentFromCSV(departmentCSVPath); err != nil {
		fmt.Println(err)
	}

	instructorController := controller.CreateInstructorController(db.DB)

	interviewController := controller.CreateInterviewController(db.DB)

	if err := applicationRoundCtrl.ReadApplicationRoundsFromCSV(roundsCSVPath); err != nil {
		fmt.Println("Failed to initialize application rounds:", err)
		return
	}

	for {
		clearScreen()

		fmt.Println("\n==== Student Recruitment System ====")
		fmt.Println("1. Register Applicant")
		fmt.Println("2. Instructor Menu")
		fmt.Println("3. Interview Management")
		fmt.Println("4. View Interview Details")
		fmt.Println("5. Show Application Rounds")
		fmt.Println("6. Show Applicant status")
		fmt.Println("7. Exit")
		fmt.Print("Select an option: ")

		var choice int
		fmt.Scan(&choice)

		scanner := bufio.NewScanner(os.Stdin)

		switch choice {
		case 1:
			cli.RegisterApplicantCLI(applicantController, applicationRoundCtrl, applicationReportCtrl, facultyCtrl, departmentCtrl)
		case 2:
			cli.InstructorCLI(instructorController)
		case 3:
			cli.InterviewCLI(interviewController)
		case 4:
			var applicantID uuid.UUID // Change to uint type
			fmt.Print("Enter Applicant ID: ")
			scanner.Scan()
			applicantID, _ = uuid.Parse(scanner.Text())
			cli.ReportInterviewDetails(db.DB, applicantID) // Pass db.DB and applicantID (as uint)
		case 5:
			rounds, err := applicationRoundCtrl.GetAllRounds()
			if err != nil {
				fmt.Println("Error fetching application rounds:", err)
				continue
			}
			fmt.Println("\n===== Application Rounds =====")
			for _, round := range rounds {
				fmt.Println("-", round.RoundName)
			}
		case 6:
			statuses, err := applicationReportCtrl.GetApplicantStatus()
			if err != nil {
				fmt.Println("Error fetching applicant statuses: %v", err)
			}

			// Print the fetched statuses
			fmt.Println("Applicant Statuses:")
			for _, status := range statuses {
				fmt.Println(status)
			}

		case 7:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid option. Try again.")
		}
		fmt.Println("\nPress Enter to continue...")
		fmt.Scanln() // Wait for user input before clearing the screen
		fmt.Scanln() // Needed to capture extra newline
	}
}
