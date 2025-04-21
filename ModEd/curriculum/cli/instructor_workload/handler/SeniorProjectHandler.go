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
			projects, err := controller.GetProjectByAdvisorID(1)
			if err != nil {
				fmt.Println("Error getting projects by advisor ID:", err)
			} else {
				fmt.Println(projects)
			}
		case "2":
			projects, err := controller.GetProjectByCommitteeID(1)
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

			err := controller.CreateEvaluation(mockEvaluation, "presentation", mockCriteria)
			if err != nil {
				fmt.Println("Error creating evaluation:", err)
			} else {
				fmt.Println("Mock evaluation created successfully")
			}
			controller.CreateEvaluation(mockEvaluation, "presentation", mockCriteria)
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
