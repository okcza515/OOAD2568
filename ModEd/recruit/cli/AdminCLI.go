package cli

import (
	"ModEd/recruit/controller"
	"ModEd/recruit/model"
	"ModEd/recruit/util"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func AdminCLI(applicantController *controller.ApplicantController, applicationReportCtrl *controller.ApplicationReportController, interviewCtrl *controller.InterviewController, adminCtrl *controller.AdminController) {
	for {
		util.ClearScreen()

		// ฟังก์ชันเฉพาะสำหรับ admin
		fmt.Println("==== Admin Menu ====")
		fmt.Println("1. Manage Applicants")
		fmt.Println("2. View Application Reports")
		fmt.Println("3. Schedule Interview")
		fmt.Println("4. Delete Interview")
		fmt.Println("5. back")
		fmt.Print("Select an option: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			// ฟังก์ชันจัดการผู้สมัคร
			ManageApplicants(applicantController)
		case 2:
			// ฟังก์ชันดูรายงานการสมัคร
			//ShowApplicationReports(applicationReportCtrl)
		case 3:
			Interview(interviewCtrl)
		case 4:
			DeleteInterview(interviewCtrl)
		case 5:
			return
		default:
			fmt.Println("Invalid option. Try again.")
		}
	}
}

func Interview(interviewCtrl *controller.InterviewController) {
	var instructorID string
	var con_int_instrucID uint
	var int_ApplicantID uint
	var interviewScore float64
	var scoreInput string
	var Status string

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

	// Get Status
	fmt.Print("Enter Status: ")
	scanner.Scan()
	Status = scanner.Text()

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
			InterviewStatus:      Status,
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
			InterviewStatus:      Status,
		}
	}

	err = interviewCtrl.CreateInterview(interview)
	if err != nil {
		fmt.Println("Failed to create interview:", err)
		return
	}

	fmt.Println("Interview scheduled successfully!")
}

func DeleteInterview(interviewCtrl *controller.InterviewController) {
	var interviewID uint

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter Interview ID to delete: ")
	scanner.Scan()
	inputID := scanner.Text()
	convInterviewID, err := strconv.ParseUint(inputID, 10, 32)
	if err != nil {
		fmt.Println("Invalid Interview ID. Please enter a valid number.")
		return
	}
	interviewID = uint(convInterviewID)

	err = interviewCtrl.DeleteInterview(interviewID)
	if err != nil {
		fmt.Println("Failed to delete interview:", err)
		return
	}

	fmt.Println("Interview deleted successfully!")
}

func ManageApplicants(applicantController *controller.ApplicantController) {
	// ฟังก์ชันจัดการผู้สมัคร
	fmt.Println("Managing Applicants...")
}
