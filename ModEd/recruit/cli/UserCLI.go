package cli

import (
	"ModEd/recruit/controller"
	"ModEd/recruit/util"
	"bufio"
	"fmt"
	"os"

	"gorm.io/gorm"
)

func UserCLI(applicantCtrl *controller.ApplicantController, applicationRoundCtrl *controller.ApplicationRoundController, applicationReportCtrl *controller.ApplicationReportController, facultyCtrl *controller.FacultyController, departmentCtrl *controller.DepartmentController) {

	for {
		// util.ClearScreen()

		fmt.Println("\n\033[1;35m╔════════════════════════════╗")
		fmt.Println("║          User Menu         ║")
		fmt.Println("╚════════════════════════════╝\033[0m")

		fmt.Println("\n\033[1;36m[1]\033[0m  Register Applicant")
		fmt.Println("\033[1;36m[2]\033[0m  View Application Status")
		fmt.Println("\033[1;36m[3]\033[0m  View Interview Details")
		fmt.Println("\033[1;36m[4]\033[0m  Back")
		fmt.Print("\n\033[1;33mSelect an option:\033[0m ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			ApplicantRegistrationCLI(applicantCtrl, applicationRoundCtrl, applicationReportCtrl, facultyCtrl, departmentCtrl)
		case 2:
			ShowApplicantReportCLI(applicantCtrl, applicationReportCtrl)
		case 3:
			scanner := bufio.NewScanner(os.Stdin)
			fmt.Print("Enter Applicant ID: ")

			var applicantID uint
			if scanner.Scan() {
				input := scanner.Text()
				_, err := fmt.Sscanf(input, "%d", &applicantID)
				if err != nil {
					fmt.Println("Invalid Apoplicant ID.")
					return
				}
				ReportInterviewDetails(applicationReportCtrl.DB, applicantID)
			} else {
				fmt.Println("Error reading input.")
			}
		case 4:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid option.")
			return
		}
		fmt.Println("\nPress Enter to return to the menu...")
		bufio.NewReader(os.Stdin).ReadBytes('\n')
		util.ClearScreen()
	}
}

func ReportInterviewDetails(db *gorm.DB, applicantID uint) {
	status, err := controller.GetApplicationStatus(db, applicantID)
	if err != nil {
		fmt.Println("เกิดข้อผิดพลาดในการดึงสถานะการสมัคร:", err)
		return
	}

	if status != "Pass" {
		fmt.Println("คุณไม่มีสิทธิ์เข้าสัมภาษณ์ เนื่องจากสถานะของคุณคือ:", status)
		return
	}

	interview, err := controller.GetInterviewDetails(db, applicantID)
	if err != nil {
		fmt.Println("เกิดข้อผิดพลาดในการดึงรายละเอียดการสัมภาษณ์:", err)
		return
	}

	scoreText := "N/A"
	if interview.InterviewScore != nil {
		scoreText = fmt.Sprintf("%.2f", *interview.InterviewScore)
	}

	fmt.Println("--- รายละเอียดการสัมภาษณ์ ---")
	fmt.Println("วันที่:", interview.ScheduledAppointment)
	fmt.Println("คะแนนสัมภาษณ์:", scoreText)
	fmt.Println("สถานะการสัมภาษณ์:", interview.InterviewStatus)
}
