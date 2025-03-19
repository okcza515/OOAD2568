// MEP-1003 Student Recruitment
package cli

import (
	"ModEd/recruit/controller"
	"ModEd/recruit/model"
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func InterviewCLI(interviewCtrl *controller.InterviewController) {
	for {
		fmt.Println("\nInterview Management")
		fmt.Println("1. Schedule Interview")
		fmt.Println("2. Delete Interview")
		fmt.Println("3. Exit")
		fmt.Print("Select an option: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			InterviewManage(interviewCtrl) // Function for scheduling an interview
		case 2:
			DeleteInterviewCLI(interviewCtrl) // Function for deleting an interview
		case 3:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

func InterviewManage(interviewCtrl *controller.InterviewController) {
	var applicantID uuid.UUID
	var instructorID string
	var con_int_instrucID uint
	var interviewScore float64
	var scoreInput string
	var Status string

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter Instructor ID: ")
	scanner.Scan()
	instructorID = scanner.Text()
	fmt.Sscanf(instructorID, "%d", &con_int_instrucID)

	fmt.Print("Enter Applicant ID: ")
	scanner.Scan()
	applicantID, _ = uuid.Parse(scanner.Text()) // Parse UUID from input

	fmt.Print("Enter Status: ")
	scanner.Scan()
	Status = scanner.Text()

	fmt.Print("Enter Scheduled Appointment (YYYY-MM-DD HH:MM:SS): ")
	scanner.Scan()
	scheduledTime := scanner.Text()

	scheduledTimeParsed, err := time.Parse("2006-01-02 15:04:05", scheduledTime)
	if err != nil {
		fmt.Println("Invalid date format. Use YYYY-MM-DD HH:MM:SS.")
		return
	}

	fmt.Print("Enter Interview Score (or press Enter to skip): ")
	scanner.Scan()
	scoreInput = scanner.Text()

	var interview *model.Interview
	if scoreInput == "" {
		interview = &model.Interview{
			ID:                   uuid.New(),
			InstructorID:         con_int_instrucID,
			ApplicantID:          applicantID,
			ScheduledAppointment: scheduledTimeParsed,
			InterviewScore:       nil,
			InterviewStatus:      Status,
		}
	} else {
		fmt.Sscanf(scoreInput, "%f", &interviewScore)
		interview = &model.Interview{
			ID:                   uuid.New(),
			InstructorID:         con_int_instrucID,
			ApplicantID:          applicantID,
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
	var interviewID uuid.UUID

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter Interview ID to delete: ")
	scanner.Scan()
	interviewID, _ = uuid.Parse(scanner.Text())

	err := interviewCtrl.DeleteInterview(interviewID)
	if err != nil {
		fmt.Println("Failed to delete interview:", err)
		return
	}

	fmt.Println("Interview deleted successfully!")
}

// ReportInterviewDetails แสดงข้อมูลสัมภาษณ์
func ReportInterviewDetails(db *gorm.DB, applicantID uuid.UUID) {
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
