package handler

import (
	"ModEd/core/cli"

	"ModEd/curriculum/controller"
	"fmt"
)

type AcademicWorkloadMenuStateHandler struct {
	manager *cli.CLIMenuStateManager
	wrapper *controller.InstructorWorkloadModuleWrapper

	instructorWorkloadModuleMenuStateHandler *InstructorWorkloadModuleMenuStateHandler
}

func NewAcademicWorkloadMenuStateHandler(
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

func (menu *AcademicWorkloadMenuStateHandler) Render() {
	fmt.Println("1. View Class Lecture")
}

func (menu *AcademicWorkloadMenuStateHandler) HandleUserInput(input string) error {
	switch input {
	case "1":
		classList, err := menu.wrapper.ClassController.GetClasses()
		if err != nil {
			fmt.Println("Error fetching meetings:", err.Error())
			return err
		}
		for _, class := range classList {
			fmt.Printf("Class: %d, Course Name: %s\n, Schedule: %s\n", class.ClassId, class.Course.Name, class.Schedule)
		}
	case "exit":
		fmt.Println("Exiting...")
		return nil
	default:
		fmt.Println("Invalid option")
	}
	return nil
}
