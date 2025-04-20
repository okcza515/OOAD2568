package cli

import (
	"ModEd/recruit/controller"
	"ModEd/recruit/util"
	"fmt"
)

func UserCLI(applicantCtrl *controller.ApplicantController, applicationRoundCtrl *controller.ApplicationRoundController, applicationReportCtrl *controller.ApplicationReportController, facultyCtrl *controller.FacultyController, departmentCtrl *controller.DepartmentController) {

	util.ClearScreen()

	fmt.Println("==== User Menu ====")
	fmt.Println("1. Register Applicant")
	fmt.Println("2. View Application Status")
	fmt.Println("3. back")
	fmt.Print("Select an option: ")

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
