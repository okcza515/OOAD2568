// MEP-1003 Student Recruitment
package controller

import (
	"ModEd/recruit/model"
	"fmt"
)

type ApplicationReport struct {
	Controller *ApplicationReportController
}

func (r *ApplicationReport) GetFilteredReport(condition map[string]interface{}) ([]*model.ApplicationReport, error) {
	return r.Controller.GetFilteredApplication(condition)
}

func (r *ApplicationReport) DisplayReport(filteredReport []model.ApplicationReport) {
	for i, report := range filteredReport {
		fmt.Printf("\nApplication #%d\n", i+1)
		fmt.Println("----------------------------------------")
		fmt.Printf("Applicant ID: %d\n", report.Applicant.ApplicantID)
		fmt.Printf("Full Name: %s %s\n", report.Applicant.FirstName, report.Applicant.LastName)
		fmt.Printf("Email: %s\n", report.Applicant.Email)
		fmt.Printf("Birth Date: %s\n", report.Applicant.BirthDate.Format("2006-01-02"))
		fmt.Printf("Phone: %s\n", report.Applicant.Phonenumber)
		fmt.Printf("GPA: %.2f\n", report.Applicant.GPAX)

		fmt.Println("\n==== Application Info ====")
		fmt.Printf("Program: %s\n", report.Program.String())
		fmt.Printf("Round: %s\n", report.ApplicationRound.RoundName)
		roundData, err := report.Applicant.GetRoundInfo()
		if err != nil {
			fmt.Println("Error retrieving round data:", err)
		} else {
			fmt.Println("Round Information:")
			for field, data := range roundData {
				fmt.Printf("  - %s: %s\n", field, data)
			}
		}
		fmt.Printf("Faculty: %s\n", report.Faculty.Name)
		fmt.Printf("Department: %s\n", report.Department.Name)

		fmt.Printf("\n\033[1;37;48m==== Status ==== \033[0m\n")
		printStatus(report.ApplicationStatuses)
		fmt.Println("----------------------------------------")
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
	case model.Confirmed:
		fmt.Printf("\033[1;35mStatus: %s\033[0m\n", status)
	case model.Withdrawn:
		fmt.Printf("\033[0;37mStatus: %s\033[0m\n", status)
	default:
		fmt.Printf("Status: %s\n", status)
	}
}
