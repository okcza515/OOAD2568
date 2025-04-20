package cli

import (
	"ModEd/recruit/controller"
	"ModEd/recruit/model"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func AdminScheduleInterviewCLI(interviewCtrl *controller.InterviewController, applicationReportCtrl *controller.ApplicationReportController) {
	var instructorID string
	var con_int_instrucID uint
	var int_ApplicantID uint
	var interviewScore float64
	var scoreInput string

	scanner := bufio.NewScanner(os.Stdin)

	// Get Instructor ID
	fmt.Print("Enter Instructor ID: ")
	scanner.Scan()
	instructorID = scanner.Text()
	convInstructorID, err := strconv.ParseUint(instructorID, 10, 32)
	if err != nil {
		fmt.Println("Invalid Instructor ID. Please enter a valid number.")
		return
	}
	con_int_instrucID = uint(convInstructorID)

	// Get Applicant ID
	fmt.Print("Enter Applicant ID: ")
	scanner.Scan()
	applicantID := scanner.Text()
	convApplicantID, err := strconv.ParseUint(applicantID, 10, 32)
	if err != nil {
		fmt.Println("Invalid Applicant ID. Please enter a valid number.")
		return
	}
	int_ApplicantID = uint(convApplicantID)

	// Get Scheduled Appointment
	fmt.Print("Enter Scheduled Appointment (YYYY-MM-DD HH:MM:SS): ")
	scanner.Scan()
	scheduledTime := scanner.Text()

	scheduledTimeParsed, err := time.Parse("2006-01-02 15:04:05", scheduledTime)
	if err != nil {
		fmt.Println("Invalid date format. Use YYYY-MM-DD HH:MM:SS.")
		return
	}

	// Get Interview Score (optional)
	fmt.Print("Enter Interview Score (or press Enter to skip): ")
	scanner.Scan()
	scoreInput = scanner.Text()

	var interview *model.Interview
	if scoreInput == "" {
		interview = &model.Interview{
			InstructorID:         con_int_instrucID,
			ApplicantID:          int_ApplicantID,
			ScheduledAppointment: scheduledTimeParsed,
			InterviewScore:       nil,
			InterviewStatus:      model.Pending,
		}
	} else {
		interviewScore, err = strconv.ParseFloat(scoreInput, 64)
		if err != nil {
			fmt.Println("Invalid interview score. Please enter a valid number.")
			return
		}
		interview = &model.Interview{
			InstructorID:         con_int_instrucID,
			ApplicantID:          int_ApplicantID,
			ScheduledAppointment: scheduledTimeParsed,
			InterviewScore:       &interviewScore,
			InterviewStatus:      model.Pending,
		}
	}

	err = interviewCtrl.CreateInterview(interview)
	if err != nil {
		fmt.Println("Failed to create interview:", err)
		return
	}

	err = applicationReportCtrl.UpdateApplicationStatus(int_ApplicantID, model.InterviewStage)
	if err != nil {
		fmt.Println("Failed to update status:", err)
	} else {
		fmt.Println("Status updated successfully!")
	}

	fmt.Println("Interview scheduled successfully!")
}
