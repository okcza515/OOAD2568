package controller

import (
	"ModEd/recruit/model"
	"fmt"
)

type InterviewDataProvider interface {
	GetAllInterviews() ([]*model.Interview, error)
}

type InterviewReport struct {
	Filters           []FilterStrategy[model.Interview]
	InterviewProvider InterviewDataProvider
}

func (r *InterviewReport) GetReport() ([]model.Interview, error) {
	ptrs, err := r.InterviewProvider.GetAllInterviews()
	if err != nil {
		return nil, err
	}

	var result []model.Interview
	for _, p := range ptrs {
		result = append(result, *p)
	}
	return result, nil
}

func (r *InterviewReport) FilterReport(report []model.Interview) ([]model.Interview, error) {
	var filtered []model.Interview = report

	for _, filter := range r.Filters {
		var err error
		filtered, err = filter.Filter(filtered)
		if err != nil {
			return nil, err
		}
	}
	return filtered, nil
}

func (r *InterviewReport) DisplayReport(filteredReport []model.Interview) {
	fmt.Println("\n==== Interview Schedule ====")
	for i, interview := range filteredReport {
		fmt.Printf("\nInterview #%d\n", i+1)
		fmt.Println("----------------------------------------")
		fmt.Printf("Interview ID       : %d\n", interview.InterviewID)
		fmt.Printf("Applicant Fullname : %s %s\n",
			interview.ApplicationReport.Applicant.FirstName,
			interview.ApplicationReport.Applicant.LastName)
		fmt.Printf("Application ID     : %d\n", interview.ApplicationReportID)
		fmt.Printf("Appointment Date   : %s\n", interview.ScheduledAppointment.Format("2006-01-02 15:04"))

		if string(interview.InterviewStatus) == "Evaluated" {
			criteriaScores, err := interview.GetCriteriaScores()
			if err != nil {
				fmt.Println("Error retrieving criteria scores:", err)
			} else {
				fmt.Println("Criteria Scores    :")
				for criterion, score := range criteriaScores {
					fmt.Printf("  - %s: %.2f\n", criterion, score)
				}
			}
		}
		fmt.Printf("Interview Status   : %s\n", interview.InterviewStatus)
		fmt.Println("----------------------------------------")
	}
}
