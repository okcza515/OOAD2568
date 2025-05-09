package cli

import (
	"ModEd/core/cli"
	"ModEd/recruit/util"
	"fmt"
)

type AdminShowApplicationReportMenuState struct {
	manager *cli.CLIMenuStateManager
	service AdminShowApplicationReportsService
	parent  cli.MenuState
}

func NewAdminShowApplicationReportMenuState(
	manager *cli.CLIMenuStateManager,
	service AdminShowApplicationReportsService,
	parent cli.MenuState,
) *AdminShowApplicationReportMenuState {
	return &AdminShowApplicationReportMenuState{
		manager: manager,
		service: service,
		parent:  parent,
	}
}

func (menu *AdminShowApplicationReportMenuState) Render() {
	util.ClearScreen()
	fmt.Println("\033[1;36mReport Menu\033[0m \n")
	fmt.Println("[1] View All Application Reports")
	fmt.Println("[2] View All Interview Reports")
	fmt.Println("[3] View Application Reports by Status")
	fmt.Println("[4] Back")
	fmt.Print("\nSelect an option: ")
}

// func (menu *AdminShowApplicationReportMenuState) Render() {
// 	util.ClearScreen()
// 	fmt.Print("Enter Applicantion Report ID to view the report: ")
// }

// func (menu *AdminShowApplicationReportMenuState) HandleUserInput(input string) error {
// 	applicantReportID, err := strconv.ParseUint(input, 10, 32)
// 	if err != nil {
// 		fmt.Println("Invalid Applicantion Report ID :", err)
// 	} else {
// 		report, err := menu.service.GetApplicationReport(uint(applicantReportID))
// 		if err != nil {
// 			fmt.Println("Error retrieving report:", err)
// 		} else {
// 			fmt.Println("\n===== Applicant Report =====")
// 			menu.service.DisplayReport([]*model.ApplicationReport{report})
// 		}
// 	}

//		fmt.Println("\nPress Enter to return...")
//		bufio.NewReader(os.Stdin).ReadBytes('\n')
//		menu.manager.SetState(menu.parent)
//		return nil
//	}
func (menu *AdminShowApplicationReportMenuState) HandleUserInput(input string) error {
	switch input {
	case "1":
		menu.viewAllApplicationReports()
	case "2":
		menu.viewAllInterviewReports()
	case "3":
		menu.viewReportsByStatus()
	case "4":
		menu.manager.SetState(menu.parent)
		return nil
	default:
		fmt.Println("Invalid option.")
	}
	util.WaitForEnter()
	return nil
}

func (menu *AdminShowApplicationReportMenuState) viewAllApplicationReports() {
	reports, err := menu.service.GetAllApplicationReports()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("\n===== All Application Reports =====")
	menu.service.DisplayOnlyApplicationReport(reports)
}

func (menu *AdminShowApplicationReportMenuState) viewAllInterviewReports() {
	reports, err := menu.service.GetAllApplicationReports()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("\n===== All Interview Reports =====")
	menu.service.DisplayOnlyInterviews(reports)
}

func (menu *AdminShowApplicationReportMenuState) viewReportsByStatus() {
	var status string
	fmt.Print("Enter status (e.g., Pending, Interview, Evaluated, Accepted, Rejected): ")
	fmt.Scanln(&status)
	reports, err := menu.service.GetApplicationReportsByStatus(status)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("\n===== Application Reports (Status: %s) =====\n", status)
	menu.service.DisplayOnlyApplicationReport(reports)
}
