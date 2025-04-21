// MEP-1003 Student Recruitment
package cli

import (
	"ModEd/recruit/controller"
	"fmt"
)

func ViewInterviewDetails(instructorCtrl *controller.InstructorController, instructorID uint) {
	interviews, err := instructorCtrl.GetInterviewsByInstructor(instructorID)
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
