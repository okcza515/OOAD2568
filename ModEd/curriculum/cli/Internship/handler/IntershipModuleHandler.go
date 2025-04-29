// MEP-1009 Student Internship
package handler

import (
	"ModEd/core/cli"
	"ModEd/curriculum/controller"
	model "ModEd/curriculum/model"
	"fmt"
	"time"
)

type InternShipModuleMenuStateHandler struct {
	menuManager *cli.CLIMenuStateManager
	wrapper     *controller.InternshipModuleWrapper
}

func NewInternShipModuleMenuStateHandler(manager *cli.CLIMenuStateManager, wrapper *controller.InternshipModuleWrapper) *InternShipModuleMenuStateHandler {
	InternshipModuleHandler := &InternShipModuleMenuStateHandler{
		menuManager: manager,
		wrapper:     wrapper,
	}

	return InternshipModuleHandler
}

func (handler *InternShipModuleMenuStateHandler) Render() {
	fmt.Println("\n==== Internship Application System ====")
	fmt.Println("1. Create Internship Application")
	fmt.Println("2. Evaluation Student Performance")
	fmt.Println("3. Evaluation Student Report")
	fmt.Println("4. Update Approval Status")
	fmt.Println("Type 'exit' to quit")
	fmt.Print("Enter your choice: ")
}

func (handler *InternShipModuleMenuStateHandler) HandleUserInput(input string) error {
	switch input {
	case "1":
		// Create a emporary internship application
		application := &model.InternshipApplication{
			TurninDate:            time.Now(),
			ApprovalAdvisorStatus: model.WAIT,
			ApprovalCompanyStatus: model.WAIT,
			AdvisorCode:           0,
			InternshipReportId:    0,
			SupervisorReviewId:    0,
			CompanyId:             0,
			StudentCode:           "65070501070",
		}
		handler.wrapper.InternshipApplication.RegisterInternshipApplications([]*model.InternshipApplication{application})
	case "2":
		handler.wrapper.Report.UpdateReportScore("65070501070", 0)
	case "3":
		handler.wrapper.Review.UpdateReviewScore("65070501070", 0, 0)
	case "4":
		handler.wrapper.Approved.UpdateApprovalStatuses("65070501070", model.APPROVED, model.APPROVED)
	default:
		fmt.Println("Invalid input")
	}

	return nil
}
