// MEP-1003 Student Recruitment
package cli

import (
	"ModEd/recruit/controller"
	"ModEd/recruit/util"
	"bufio"
	"fmt"
	"os"
)

func UserCLI(applicantCtrl *controller.ApplicantController, applicationRoundCtrl *controller.ApplicationRoundController, applicationReportCtrl *controller.ApplicationReportController, facultyCtrl *controller.FacultyController, departmentCtrl *controller.DepartmentController) {

	for {
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
			ApplicantRegistrationCLI(applicantCtrl, applicationRoundCtrl, applicationReportCtrl, facultyCtrl, departmentCtrl)
		case 2:
			ShowApplicantReportCLI(applicantCtrl, applicationReportCtrl)
		case 3:
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
