// MEP-1003 Student Recruitment
package controller

import (
	"ModEd/recruit/model"
	"fmt"
)

type InterviewReport struct {
	Controller *InterviewController
}

func (r *InterviewReport) GetFilteredInterviews(condition map[string]interface{}) ([]*model.Interview, error) {
	return r.Controller.GetFilteredInterviews(condition)
}

func (r *InterviewReport) DisplayReport(filteredReport []*model.Interview) {
	for _, interview := range filteredReport {
		// fmt.Printf("\nInterview #%d\n", i+1)
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
