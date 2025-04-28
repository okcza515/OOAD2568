package handler

import (
	"ModEd/core/cli"
	"ModEd/curriculum/controller"
	"fmt"
)

type WILModuleMenuStateHandler struct {
	menuManger *cli.CLIMenuStateManager
	wrapper    *controller.WILModuleWrapper

	WILProjectApplicationMenuStateHandler *WILProjectApplicationMenuStateHandler
	WILProjectCurriculumMenuStateHandler  *WILProjectCurriculumMenuStateHandler
	WILProjectMenuStateHandler            *WILProjectMenuStateHandler
	IndependentStudyMenuStateHandler      *IndependentStudyMenuStateHandler
}

func NewWILModuleMenuStateHandler(manager *cli.CLIMenuStateManager, wrapper *controller.WILModuleWrapper) *WILModuleMenuStateHandler {
	wilmoduleHandler := &WILModuleMenuStateHandler{
		menuManger: manager,
		wrapper:    wrapper,
	}

	wilmoduleHandler.WILProjectApplicationMenuStateHandler = NewWILProjectApplicationMenuStateHandler(manager, wrapper, wilmoduleHandler)
	wilmoduleHandler.WILProjectCurriculumMenuStateHandler = NewWILProjectCurriculumMenuStateHandler(manager, wrapper, wilmoduleHandler)
	wilmoduleHandler.WILProjectMenuStateHandler = NewWILProjectMenuStateHandler(manager, wrapper, wilmoduleHandler)
	wilmoduleHandler.IndependentStudyMenuStateHandler = NewIndependentStudyMenuStateHandler(manager, wrapper, wilmoduleHandler)

	return wilmoduleHandler
}

func (handler *WILModuleMenuStateHandler) Render() {
	fmt.Println("\nWIL Module Menu:")
	fmt.Println("1. WIL Project Curriculum")
	fmt.Println("2. WIL Project Application")
	fmt.Println("3. WIL Project")
	fmt.Println("4. Independent Study")
	fmt.Println("0. Exit WIL Module")
}

func (handler *WILModuleMenuStateHandler) HandleUserInput(input string) error {
	switch input {
	case "1":
		handler.menuManger.SetState(handler.WILProjectCurriculumMenuStateHandler)
	case "2":
		handler.menuManger.SetState(handler.WILProjectApplicationMenuStateHandler)
	case "3":
		handler.menuManger.SetState(handler.WILProjectMenuStateHandler)
	case "4":
		handler.menuManger.SetState(handler.IndependentStudyMenuStateHandler)
	default:
		fmt.Println("invalid input")
	}

	return nil
}
