package cli

import (
	"ModEd/recruit/controller"
	"ModEd/recruit/model"
	"ModEd/recruit/util"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func ShowApplicantReportCLI(
	applicantCtrl *controller.ApplicantController,
	applicationReportCtrl *controller.ApplicationReportController,
) {
	util.ClearScreen()
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter Applicant ID to view the report: ")
	scanner.Scan()
	applicantIDStr := scanner.Text()

	applicantID, err := strconv.ParseUint(applicantIDStr, 10, 32)
	if err != nil {
		fmt.Println("Invalid applicant ID:", err)
		return
	}

	applicant, err := applicantCtrl.GetApplicantByID(uint(applicantID))
	if err != nil {
		fmt.Println("Error fetching applicant details:", err)
		return
	}

	report, err := applicationReportCtrl.GetApplicationReportByApplicantID(uint(applicantID))
	if err != nil {
		fmt.Println("Error fetching application report:", err)
		return
	}

	displayApplicantReport(&applicant, report)

	fmt.Println("\n-- Press Enter to continue... --")
	scanner.Scan()
}

func displayApplicantReport(applicant *model.Applicant, report *model.ApplicationReport) {
	fmt.Println("\n==== Applicant Report ====")
	fmt.Printf("Applicant ID: %d\n", applicant.ApplicantID)
	fmt.Printf("Full Name: %s %s\n", applicant.FirstName, applicant.LastName)

	if report != nil {
		// fmt.Printf("\n==== Application Report ====\n")
		// fmt.Printf("Application Round: %s\n", report.ApplicationRoundsID)
		// fmt.Printf("Faculty: %s\n", report.FacultyID)
		// fmt.Printf("Department: %s\n", report.DepartmentID)

		fmt.Printf("\n\033[1;37;48m==== Status ==== \033[0m\n")
		printStatus(report.ApplicationStatuses)
	} else {
		fmt.Println("No application report found for this applicant.")
	}
}

func printStatus(status model.ApplicationStatus) {
	switch status {
	case model.Pending:
		fmt.Printf("\033[1;33mStatus: %s\033[0m\n", status)
	case model.InterviewStage:
		fmt.Printf("\033[1;36mStatus: %s\033[0m\n", status)
	case model.Accepted:
		fmt.Printf("\033[1;32mStatus: %s\033[0m\n", status)
	case model.Rejected:
		fmt.Printf("\033[1;31mStatus: %s\033[0m\n", status)
	default:
		fmt.Printf("Status: %s\n", status)
	}
}
