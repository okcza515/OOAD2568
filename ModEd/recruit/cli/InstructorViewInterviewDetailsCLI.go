// MEP-1003 Student Recruitment
package cli

import (
	"ModEd/recruit/controller"
)

func ViewInterviewDetails(instructorViewInterviewDetailsService InstructorViewInterviewDetailsService, instructorID uint, filter string, interviewController *controller.InterviewController) {
	report := controller.InterviewReport{}
	interviews, err := instructorViewInterviewDetailsService.ViewInterviewDetails(instructorID, filter, interviewController)
	if err != nil {
		println("Error fetching interview details:", err.Error())
		return
	}
	report.DisplayReport(interviews)
}
