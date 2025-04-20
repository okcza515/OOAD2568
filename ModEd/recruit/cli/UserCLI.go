package cli

import (
	"ModEd/recruit/controller"
	"ModEd/recruit/util"
	"fmt"
)

func UserCLI(applicantCtrl *controller.ApplicantController, applicationRoundCtrl *controller.ApplicationRoundController, applicationReportCtrl *controller.ApplicationReportController, facultyCtrl *controller.FacultyController, departmentCtrl *controller.DepartmentController) {

	util.ClearScreen()

	fmt.Println("\n\033[1;35m╔════════════════════════════╗")
	fmt.Println("║          User Menu         ║")
	fmt.Println("╚════════════════════════════╝\033[0m")

	fmt.Println("\n\033[1;36m[1]\033[0m  Register Applicant")
	fmt.Println("\033[1;36m[2]\033[0m  View Application Status")
	fmt.Println("\033[1;36m[3]\033[0m  Back")
	fmt.Print("\n\033[1;33mSelect an option:\033[0m ")

	var choice int
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		RegisterApplicantCLI(applicantCtrl, applicationRoundCtrl, applicationReportCtrl, facultyCtrl, departmentCtrl)
	case 2:
		ShowApplicantReportCLI(applicantCtrl, applicationReportCtrl)
	default:
		fmt.Println("Invalid option.")
		return
	}
}
