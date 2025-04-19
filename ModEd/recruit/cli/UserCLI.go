package cli

import (
	"ModEd/recruit/controller"
	"ModEd/recruit/util"
	"fmt"
)

func UserCLI(applicantController *controller.ApplicantController, applicationRoundCtrl *controller.ApplicationRoundController, applicationReportCtrl *controller.ApplicationReportController, facultyCtrl *controller.FacultyController, departmentCtrl *controller.DepartmentController) {

	util.ClearScreen()

	// ฟังก์ชันเฉพาะสำหรับผู้ใช้ทั่วไป
	fmt.Println("==== User Menu ====")
	fmt.Println("1. Register Applicant")
	fmt.Println("2. View Application Status")
	fmt.Println("3. back")
	fmt.Print("Select an option: ")

	var choice int
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		// เรียกใช้ฟังก์ชันสมัครผู้สมัคร
		RegisterApplicantCLI(applicantController, applicationRoundCtrl, applicationReportCtrl, facultyCtrl, departmentCtrl)
	case 2:
		// เรียกดูสถานะของผู้สมัคร
		statuses, err := applicationReportCtrl.GetApplicantStatus()
		if err != nil {
			fmt.Println("Error fetching applicant statuses:", err)
			return
		}
		fmt.Println("Applicant Statuses:")
		for _, status := range statuses {
			fmt.Println(status)
		}
	default:
		fmt.Println("Invalid option.")
		return
	}
}
