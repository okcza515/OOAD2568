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
}

func NewWILModuleMenuStateHandler(manager *cli.CLIMenuStateManager, proxy *controller.WILModuleProxy) *WILModuleMenuStateHandler {
	wilmoduleHandler := &WILModuleMenuStateHandler{
		menuManger: manager,
		proxy:      proxy,
	}

	wilmoduleHandler.WILProjectApplicationMenuStateHandler = NewWILProjectApplicationMenuStateHandler(manager, proxy, wilmoduleHandler)
	wilmoduleHandler.WILProjectCurriculumMenuStateHandler = NewWILProjectCurriculumMenuStateHandler(manager, proxy, wilmoduleHandler)

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
	//case "3":
	//	//SupplyHandler(menu.controllerFacade)
	//	handler.menuManger.SetState(handler.instrumentMenu)
	//case "4":
	//	handler.menuManger.SetState(handler.instrumentMenu)
	case "5":
		fmt.Println("Not implemented yet...")
	case "6":
		fmt.Println("Not implemented yet...")
	default:
		fmt.Println("invalid input")
	}

	return nil
}
