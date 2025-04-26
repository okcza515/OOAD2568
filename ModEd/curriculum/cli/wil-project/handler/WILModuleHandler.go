package handler

import (
	"ModEd/core/cli"
	"ModEd/curriculum/controller"
	"fmt"
)

type WILModuleMenuStateHandler struct {
	menuManger *cli.CLIMenuStateManager
	proxy      *controller.WILModuleProxy

	WILProjectApplicationMenuStateHandler *WILProjectApplicationMenuStateHandler
	WILProjectCurriculumMenuStateHandler  *WILProjectCurriculumMenuStateHandler
	WILProjectMenuStateHandler            *WILProjectMenuStateHandler
	IndependentStudyMenuStateHandler	*IndependentStudyMenuStateHandler
}

func NewWILModuleMenuStateHandler(manager *cli.CLIMenuStateManager, proxy *controller.WILModuleProxy) *WILModuleMenuStateHandler {
	wilmoduleHandler := &WILModuleMenuStateHandler{
		menuManger: manager,
		proxy:      proxy,
	}

	wilmoduleHandler.WILProjectApplicationMenuStateHandler = NewWILProjectApplicationMenuStateHandler(manager, proxy, wilmoduleHandler)
	wilmoduleHandler.WILProjectCurriculumMenuStateHandler = NewWILProjectCurriculumMenuStateHandler(manager, proxy, wilmoduleHandler)
	wilmoduleHandler.WILProjectMenuStateHandler = NewWILProjectMenuStateHandler(manager, proxy, wilmoduleHandler)
	wilmoduleHandler.IndependentStudyMenuStateHandler = NewIndependentStudyMenuStateHandler(manager, proxy, wilmoduleHandler)

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
