package handler

import (
	"ModEd/core/cli"
	"ModEd/curriculum/controller"
	"ModEd/curriculum/model"
	projectModel "ModEd/project/model"
	"fmt"
)

type SeniorProjectWorkloadMenuStateHandler struct {
	manager *cli.CLIMenuStateManager
	wrapper *controller.InstructorWorkloadModuleWrapper

	instructorWorkloadModuleMenuStateHandler *InstructorWorkloadModuleMenuStateHandler
}

func NewSeniorProjectModuleStateHandler(
	manager *cli.CLIMenuStateManager,
	wrapper *controller.InstructorWorkloadModuleWrapper,
	instructorWorkloadModuleMenuStateHandler *InstructorWorkloadModuleMenuStateHandler,
) *SeniorProjectWorkloadMenuStateHandler {
	return &SeniorProjectWorkloadMenuStateHandler{
		manager:                                  manager,
		wrapper:                                  wrapper,
		instructorWorkloadModuleMenuStateHandler: instructorWorkloadModuleMenuStateHandler,
	}
}

func (menu *SeniorProjectWorkloadMenuStateHandler) Render() {
	fmt.Println("\nAcademic Workload Menu:")
	fmt.Println("1. Get All Projects By Advisor ID")
	fmt.Println("2. Get All Projects By Committee ID")
	fmt.Println("3. Evaluate Project as Advisor")
	fmt.Println("4. Evaluate Project as Committee")
	fmt.Println("Type 'exit' to quit")
}

func (menu *SeniorProjectWorkloadMenuStateHandler) HandleUserInput(input string) error {
	switch input {
	case "1":
		projects, err := menu.wrapper.SeniorProjectController.GetProjectByAdvisorID(1)
		if err != nil {
			fmt.Println("Error getting projects by advisor ID:", err)
		} else {
			fmt.Println(projects)
		}
	case "2":
		projects, err := menu.wrapper.SeniorProjectController.GetProjectByCommitteeID(1)
		if err != nil {
			fmt.Println("Error getting projects by committee ID:", err)
		} else {
			fmt.Println(projects)
		}
	case "3":
		mockEvaluation := &model.ProjectEvaluation{
			GroupId:        1,
			AssignmentId:   1,
			AssignmentType: "presentation",
			Score:          0,
			Comment:        "Improved English skills",
		}

		mockCriteria := []projectModel.AssessmentCriteria{
			{CriteriaName: "Good presentation"},
			{CriteriaName: "Answered all questions"},
			{CriteriaName: "Good teamwork"},
		}
		err := menu.wrapper.SeniorProjectController.CreateEvaluation(mockEvaluation, "presentation", mockCriteria)
		if err != nil {
			fmt.Println("Error creating evaluation:", err)
		} else {
			fmt.Println("Mock evaluation created successfully")
		}
		menu.wrapper.SeniorProjectController.CreateEvaluation(mockEvaluation, "presentation", mockCriteria)
	case "4":
		fmt.Println("Evaluate Project as Committee Not implemented yet...")
	case "exit":
		fmt.Println("Exiting...")
		return nil
	default:
		fmt.Println("Invalid option")
	}
	return nil
}
