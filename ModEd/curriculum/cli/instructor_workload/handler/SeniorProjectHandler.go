package handler

import (
	"ModEd/curriculum/controller"
	"ModEd/curriculum/model"
	"ModEd/curriculum/utils"
	projectModel "ModEd/project/model"
	"fmt"
)

func RunSeniorProjectWorkloadHandler(controller controller.ProjectControllerService) {
	for {
		DisplaySeniorProjectWorkloadModuleMenu()
		choice := utils.GetUserChoice()
		fmt.Println("choice: ", choice)

		switch choice {
		case "1":
			fmt.Println("Get All Projects By Advisor ID Not implemented yet...")
		case "2":
			fmt.Println("Get All Projects By Committee ID Not implemented yet...")
		case "3":

			mockEvaluation := &model.ProjectEvaluation{
				GroupId:        1,
				AssignmentId:   1,
				AssignmentType: "mockAssignmentType",
				Score:          95.0,
				Comment:        "mockComment",
			}

			mockCriteria := []projectModel.AssessmentCriteria{
				{CriteriaName: "mockCriteriaName1"},
				{CriteriaName: "mockCriteriaName2"},
			}

			err := controller.CreateEvaluation(mockEvaluation, "mockStrategy", mockCriteria)
			if err != nil {
				fmt.Println("Error creating evaluation:", err)
			} else {
				fmt.Println("Mock evaluation created successfully")
			}
			controller.CreateEvaluation(mockEvaluation, "mockStrategy", mockCriteria)
		case "4":
			fmt.Println("Evaluate Project as Committee Not implemented yet...")
		case "exit":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid option")
		}
	}
}

func DisplaySeniorProjectWorkloadModuleMenu() {
	fmt.Println("\nAcademic Workload Menu:")
	fmt.Println("1. Get All Projects By Advisor ID")
	fmt.Println("2. Get All Projects By Committee ID")
	fmt.Println("3. Evaluate Project as Advisor")
	fmt.Println("4. Evaluate Project as Committee")

	fmt.Println("Type 'exit' to quit")
}
