// MEP-1003 Student Recruitment
package cli

import (
	"fmt"
)

func ViewInterviewDetails(instructorViewInterviewDetailsService InstructorViewInterviewDetailsService, instructorID uint) {
	interviews, err := instructorViewInterviewDetailsService.ViewInterviewDetails(instructorID)
	if err != nil {
		fmt.Println("Error retrieving interviews:", err)
		return
	}

	fmt.Println("\n==== Interview Schedule ====")
	for _, interview := range interviews {
		fmt.Printf("ID: %d | Applicant ID: %d | Date: %s | Score: ",
			interview.ID, interview.ApplicantID, interview.ScheduledAppointment)
		if interview.InterviewScore != nil {
			fmt.Println(*interview.InterviewScore)
		} else {
			fmt.Println("Not Assigned")
		}
	}
}
