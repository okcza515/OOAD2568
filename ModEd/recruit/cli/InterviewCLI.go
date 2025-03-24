package cli

import (
	"ModEd/recruit/controller"
	"ModEd/recruit/model"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"

	"gorm.io/gorm"
)

func InterviewCLI(interviewCtrl *controller.InterviewController) {
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

func DeleteInterviewCLI(interviewCtrl *controller.InterviewController) {
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

// ReportInterviewDetails แสดงข้อมูลสัมภาษณ์
func ReportInterviewDetails(db *gorm.DB, applicantID uint) {
	// ดึงสถานะจาก interview_status
	status, err := controller.GetApplicationStatus(db, applicantID)
	if err != nil {
		fmt.Println("เกิดข้อผิดพลาดในการดึงสถานะการสมัคร:", err)
		return
	}

	if status != "Pass" {
		fmt.Println("คุณไม่มีสิทธิ์เข้าสัมภาษณ์ เนื่องจากสถานะของคุณคือ:", status)
		return
	}

	// ดึงข้อมูลสัมภาษณ์
	interview, err := controller.GetInterviewDetails(db, applicantID)
	if err != nil {
		fmt.Println("เกิดข้อผิดพลาดในการดึงรายละเอียดการสัมภาษณ์:", err)
		return
	}

	//  เช็คว่า `InterviewScore` มีค่าหรือไม่
	scoreText := "N/A"
	if interview.InterviewScore != nil {
		scoreText = fmt.Sprintf("%.2f", *interview.InterviewScore) // แสดงทศนิยม 2 ตำแหน่ง
	}

	// แสดงรายละเอียดการสัมภาษณ์
	fmt.Println("--- รายละเอียดการสัมภาษณ์ ---")
	fmt.Println("วันที่:", interview.ScheduledAppointment)
	fmt.Println("คะแนนสัมภาษณ์:", scoreText)
	fmt.Println("สถานะการสัมภาษณ์:", interview.InterviewStatus)
}
