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

	if len(interviews) == 0 {
		fmt.Println("No interviews found for this instructor.")
		return
	}

	fmt.Println("\n==== Interview Schedule ====")
	for i, interview := range interviews {
		fmt.Printf("\nInterview #%d\n", i+1)
		fmt.Println("----------------------------------------")
		fmt.Printf("Interview ID       : %d\n", interview.InterviewID)
		fmt.Printf("Applicant Fullname : %s %s \n",
			interview.ApplicationReport.Applicant.FirstName,
			interview.ApplicationReport.Applicant.LastName)
		fmt.Printf("Application ID     : %d\n", interview.ApplicationReportID)
		fmt.Printf("Appointment Date   : %s\n", interview.ScheduledAppointment.Format("2006-01-02 15:04"))
		fmt.Printf("Interview Status   : %s\n", interview.InterviewStatus)
		fmt.Println("----------------------------------------")
	}
}
