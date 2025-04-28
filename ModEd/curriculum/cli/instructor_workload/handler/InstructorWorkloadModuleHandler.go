package handler

import (
	"ModEd/core/cli"
	"ModEd/curriculum/controller"
)

type InstructorWorkloadModuleMenuStateHandler struct {
	menuManger *cli.CLIMenuStateManager
	wrapper    *controller.InstructorWorkloadModuleWrapper
}

func NewInstructorWorkloadModuleMenuStateHandler(manager *cli.CLIMenuStateManager, wrapper *controller.InstructorWorkloadModuleWrapper) *InstructorWorkloadModuleMenuStateHandler {
	instructorWorkloadModuleHandler := &InstructorWorkloadModuleMenuStateHandler{
		menuManger: manager,
		wrapper:    wrapper,
	}
	return instructorWorkloadModuleHandler
}
