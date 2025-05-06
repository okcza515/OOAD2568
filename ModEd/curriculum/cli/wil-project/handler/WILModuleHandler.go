package handler

import (
	"ModEd/asset/util"
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

	wilmoduleHandler.menuManger.AddMenu("1", wilmoduleHandler.WILProjectCurriculumMenuStateHandler)
	wilmoduleHandler.menuManger.AddMenu("2", wilmoduleHandler.WILProjectApplicationMenuStateHandler)
	wilmoduleHandler.menuManger.AddMenu("3", wilmoduleHandler.WILProjectMenuStateHandler)
	wilmoduleHandler.menuManger.AddMenu("4", wilmoduleHandler.IndependentStudyMenuStateHandler)
	wilmoduleHandler.menuManger.AddMenu("exit", nil)

	return wilmoduleHandler
}

func (handler *WILModuleMenuStateHandler) Render() {
	fmt.Println("\nWIL Module Menu:")
	fmt.Println("1. WIL Project Curriculum")
	fmt.Println("2. WIL Project Application")
	fmt.Println("3. WIL Project")
	fmt.Println("4. Independent Study")
	fmt.Println("exit: Exit the module")
}

func (handler *WILModuleMenuStateHandler) HandleUserInput(input string) error {

	err := handler.menuManger.GoToMenu(input)
	if err != nil {
		fmt.Println("err: Invalid input, menu '" + input + "' doesn't exist")
		util.PressEnterToContinue()
	}

	return nil
}
