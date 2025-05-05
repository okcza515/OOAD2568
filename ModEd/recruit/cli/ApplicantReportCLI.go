// MEP-1003 Student Recruitment
package cli

import (
	"ModEd/core/cli"
	"ModEd/recruit/controller"
	"ModEd/recruit/model"
	"ModEd/recruit/util"
	recruitUtil "ModEd/recruit/util"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type ApplicantReportMenuState struct {
	manager          *cli.CLIMenuStateManager
	reportService    ApplicantReportService
	interviewService InterviewService
	parent           cli.MenuState
}

func NewApplicantReportMenuState(
	manager *cli.CLIMenuStateManager,
	reportService ApplicantReportService,
	interviewService InterviewService,
	parent cli.MenuState,
) *ApplicantReportMenuState {
	return &ApplicantReportMenuState{
		manager:          manager,
		reportService:    reportService,
		interviewService: interviewService,
		parent:           parent,
	}
}

func (menu *ApplicantReportMenuState) Render() {
	util.ClearScreen()
	fmt.Print("Enter Applicantion Report ID to view the report: ")
}

func (menu *ApplicantReportMenuState) HandleUserInput(input string) error {
	applicantionID, err := strconv.ParseUint(input, 10, 32)
	if err != nil {
		fmt.Println("Invalid applicant ID:", err)
	} else {
		reportList, err := menu.reportService.GetFullApplicationReportByApplicationID(uint(applicantionID))
		if err != nil {
			fmt.Println("Failed to fetch application report:", err)
			recruitUtil.WaitForEnter()
			return nil
		}

		if len(reportList) == 0 {
			fmt.Println("No application report found.")
			recruitUtil.WaitForEnter()
			return nil
		}

		report := reportList[0]
		if err != nil {
			fmt.Println("Error fetching application report:", err)
		} else {
			reportDisplay := controller.ApplicationReport{}
			reportDisplay.DisplayReport([]model.ApplicationReport{report})

			if report.ApplicationStatuses == model.InterviewStage {
				ReportInterviewDetails(menu.interviewService, report.Applicant.ApplicantID)
			}
		}
	}

	fmt.Println("\nPress Enter to return...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
	menu.manager.SetState(menu.parent)
	return nil
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

func ReportInterviewDetails(interviewService InterviewService, applicantID uint) {
	interview, err := interviewService.GetInterviewDetails(applicantID)
	if err != nil {
		fmt.Println("An error occurred while fetching the interview details:", err)
		return
	}

	scoreText := "N/A"
	if interview.TotalScore != 0 {
		scoreText = fmt.Sprintf("%.2f", *&interview.TotalScore)
	}

	fmt.Println("\n==== Interview Details ====")
	fmt.Println("Scheduled Date:", interview.ScheduledAppointment)
	fmt.Println("Interview Score:", scoreText)
	fmt.Println("Interview Status:", interview.InterviewStatus)
}
