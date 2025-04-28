package handler

import (
	"ModEd/core/cli"
	"ModEd/curriculum/controller"
	"fmt"
)

type InstructorWorkloadModuleMenuStateHandler struct {
	menuManger *cli.CLIMenuStateManager
	wrapper    *controller.InstructorWorkloadModuleWrapper

	SeniorProjectWorkloadMenuStateHandler *SeniorProjectWorkloadMenuStateHandler
}

func NewInstructorWorkloadModuleMenuStateHandler(manager *cli.CLIMenuStateManager, wrapper *controller.InstructorWorkloadModuleWrapper) *InstructorWorkloadModuleMenuStateHandler {
	instructorWorkloadModuleHandler := &InstructorWorkloadModuleMenuStateHandler{
		menuManger: manager,
		wrapper:    wrapper,
	}

	instructorWorkloadModuleHandler.SeniorProjectWorkloadMenuStateHandler = NewSeniorProjectModuleStateHandler(manager, wrapper, instructorWorkloadModuleHandler)

	return instructorWorkloadModuleHandler
}

func (handler *InstructorWorkloadModuleMenuStateHandler) Render() {
	fmt.Println("\nInstructor Workload Module Menu:")
	fmt.Println("1. Load CSV Seed Data")
	fmt.Println("2. Academic")
	fmt.Println("3. Student Advisor Workload")
	fmt.Println("4. Administrative Task")
	fmt.Println("5. Senior Project")
	fmt.Println("Type 'exit' to quit")
}

func (handler *InstructorWorkloadModuleMenuStateHandler) HandleUserInput(input string) error {
	switch input {
	case "5":
		handler.menuManger.SetState(handler.SeniorProjectWorkloadMenuStateHandler)
	default:
		fmt.Println("Invalid input")
	}

	return nil
}
