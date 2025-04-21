// MEP-1003 Student Recruitment
package cli

import (
	"ModEd/recruit/controller"
	"ModEd/recruit/model"
	"ModEd/recruit/util"
	"bufio"
	"fmt"
	"os"
	"strconv"

	"gorm.io/gorm"
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

	report, err := applicationReportCtrl.GetFullApplicationReportByApplicantID(uint(applicantID))
	if err != nil {
		fmt.Println("Error fetching application report:", err)
		return
	}

	displayApplicantReport(report)

	if report != nil && report.ApplicationStatuses == model.InterviewStage {
		ReportInterviewDetails(applicationReportCtrl.DB, uint(applicantID))
	}
}

func displayApplicantReport(report *model.ApplicationReport) {
	fmt.Println("\n==== Applicant Report ====")
	fmt.Printf("Applicant ID: %d\n", report.Applicant.ApplicantID)
	fmt.Printf("Full Name: %s %s\n", report.Applicant.FirstName, report.Applicant.LastName)
	fmt.Printf("Email: %s\n", report.Applicant.Email)
	fmt.Printf("Phone: %s\n", report.Applicant.Phonenumber)
	fmt.Printf("GPA: %.2f\n", report.Applicant.GPAX)

	fmt.Println("\n==== Application Info ====")
	fmt.Printf("Round: %s\n", report.ApplicationRound.RoundName)
	fmt.Printf("Faculty: %s\n", report.Faculty.Name)
	fmt.Printf("Department: %s\n", report.Department.Name)

	fmt.Printf("\n\033[1;37;48m==== Status ==== \033[0m\n")
	printStatus(report.ApplicationStatuses)
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

func ReportInterviewDetails(db *gorm.DB, applicantID uint) {
	interview, err := controller.GetInterviewDetails(db, applicantID)
	if err != nil {
		fmt.Println("An error occurred while fetching the interview details:", err)
		return
	}

	scoreText := "N/A"
	if interview.InterviewScore != nil {
		scoreText = fmt.Sprintf("%.2f", *interview.InterviewScore)
	}

	fmt.Println("\n==== Interview Details ====")
	fmt.Println("Scheduled Date:", interview.ScheduledAppointment)
	fmt.Println("Interview Score:", scoreText)
	fmt.Println("Interview Status:", interview.InterviewStatus)
}
