package controller

import (
	"ModEd/recruit/model"
	"fmt"
)

type ApplicationDataProvider interface {
	GetAllApplicationReports() ([]*model.ApplicationReport, error)
}

type ApplicationReport struct {
	Filters             []FilterStrategy[model.ApplicationReport]
	ApplicationProvider ApplicationDataProvider
}

func (r *ApplicationReport) GetReport() ([]model.ApplicationReport, error) {
	ptrs, err := r.ApplicationProvider.GetAllApplicationReports()
	if err != nil {
		return nil, err
	}

	var result []model.ApplicationReport
	for _, p := range ptrs {
		result = append(result, *p)
	}
	return result, nil
}

func (r *ApplicationReport) FilterReport(report []model.ApplicationReport) ([]model.ApplicationReport, error) {
	var filtered []model.ApplicationReport = report

	for _, filter := range r.Filters {
		var err error
		filtered, err = filter.Filter(filtered)
		if err != nil {
			return nil, err
		}
	}
	return filtered, nil
}

func (r *ApplicationReport) DisplayReport(filteredReport []model.ApplicationReport) {
	fmt.Println("\n==== Interview Schedule ====")
	for _, report := range filteredReport {
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
	default:
		fmt.Printf("Status: %s\n", status)
	}
}
