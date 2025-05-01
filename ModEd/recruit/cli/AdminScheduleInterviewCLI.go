// MEP-1003 Student Recruitment
package cli

import (
	"ModEd/recruit/controller"
	"ModEd/recruit/model"
	recruitUtil "ModEd/recruit/util"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func AdminScheduleInterviewCLI(interviewCtrl *controller.InterviewController, applicationReportCtrl *controller.ApplicationReportController) {
	var instructorID string
	var con_int_instrucID uint
	var applicationReportID uint

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter Instructor ID: ")
	scanner.Scan()
	instructorID = scanner.Text()
	convInstructorID, err := strconv.ParseUint(instructorID, 10, 32)
	if err != nil {
		fmt.Println("Invalid Instructor ID. Please enter a valid number.")
		return
	}
	con_int_instrucID = uint(convInstructorID)

	fmt.Print("Enter Application Report ID: ")
	scanner.Scan()
	reportID := scanner.Text()
	convReportID, err := strconv.ParseUint(reportID, 10, 32)
	if err != nil {
		fmt.Println("Invalid Application Report ID. Please enter a valid number.")
		return
	}
	applicationReportID = uint(convReportID)

	applicationReport, err := applicationReportCtrl.GetApplicationReportByID(applicationReportID)
	if err != nil {
		fmt.Println("Failed to retrieve ApplicationReport:", err)
		return
	}

	if applicationReport.ApplicationStatuses != "Pending" {
		fmt.Println("\n\033[1;33m⚠ You cannot assign interview details at this stage.\033[0m")
		fmt.Printf("%d current application status is: \033[1;31m%s \031[0m \n", convReportID, applicationReport.ApplicationStatuses)
		scanner.Scan()
		return
	}

	fmt.Print("Enter Scheduled Appointment (YYYY-MM-DD HH:MM:SS): ")
	scanner.Scan()
	scheduledTime := scanner.Text()

	scheduledTimeParsed, err := time.Parse("2006-01-02 15:04:05", scheduledTime)
	if err != nil {
		fmt.Println("Invalid date format. Use YYYY-MM-DD HH:MM:SS.")
		recruitUtil.WaitForEnter()
		return
	}

	var interview *model.Interview
	interview = &model.Interview{
		InstructorID:         con_int_instrucID,
		ApplicationReportID:  applicationReportID,
		ScheduledAppointment: scheduledTimeParsed,
		CriteriaScores:       "",
		TotalScore:           0,
		EvaluatedAt:          time.Time{},
		InterviewStatus:      model.Pending,
	}

	err = interviewCtrl.CreateInterview(interview)
	if err != nil {
		fmt.Println("Failed to create interview:", err)
		recruitUtil.WaitForEnter()
		return
	}
	fmt.Printf("Inserted Interview ID: %d\n", interview.InterviewID)

	err = applicationReportCtrl.UpdateApplicationStatus(applicationReport.ApplicationReportID, model.InterviewStage)
	if err != nil {
		fmt.Println("Failed to update status:", err)
	} else {
		fmt.Println("Status updated successfully!")
	}

	fmt.Println("Interview scheduled successfully!")

	recruitUtil.WaitForEnter()
}
