package cli

import (
	"ModEd/recruit/controller"
	"ModEd/recruit/util"
	"fmt"
)

func AdminCLI(applicantController *controller.ApplicantController, applicationReportCtrl *controller.ApplicationReportController) {

	util.ClearScreen()

	// ฟังก์ชันเฉพาะสำหรับ admin
	fmt.Println("==== Admin Menu ====")
	fmt.Println("1. Manage Applicants")
	fmt.Println("2. View Application Reports")
	fmt.Println("3. back")
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
		return
	default:
		fmt.Println("Invalid option. Try again.")
	}
}

func ManageApplicants(applicantController *controller.ApplicantController) {
	// ฟังก์ชันจัดการผู้สมัคร
	fmt.Println("Managing Applicants...")
}

// func ShowApplicationReports(applicationReportCtrl *controller.ApplicationReportController) {
// 	// ฟังก์ชันดูรายงานการสมัคร
// 	reports, err := applicationReportCtrl.GetAllReports()
// 	if err != nil {
// 		fmt.Println("Error fetching application reports:", err)
// 		return
// 	}
// 	fmt.Println("Application Reports:")
// 	for _, report := range reports {
// 		fmt.Println(report)
// 	}
// }
