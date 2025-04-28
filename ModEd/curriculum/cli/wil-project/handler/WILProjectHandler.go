// MEP-1010 Work Integrated Learning (WIL)
package handler

import (
	"ModEd/asset/util"
	"ModEd/core/cli"
	"ModEd/curriculum/controller"
	"fmt"
)

type WILProjectMenuStateHandler struct {
	manager *cli.CLIMenuStateManager
	wrapper *controller.WILModuleWrapper

	wilModuleMenuStateHandler *WILModuleMenuStateHandler
}

func NewWILProjectMenuStateHandler(
	manager *cli.CLIMenuStateManager, wrapper *controller.WILModuleWrapper, wilModuleMenuStateHandler *WILModuleMenuStateHandler,
) *WILProjectMenuStateHandler {
	return &WILProjectMenuStateHandler{
		manager:                   manager,
		wrapper:                   wrapper,
		wilModuleMenuStateHandler: wilModuleMenuStateHandler,
	}
}

func (menu *WILProjectMenuStateHandler) Render() {
	fmt.Println("\nWIL Project Menu:")
	fmt.Println("1. Create WIL Project")
	fmt.Println("2. Edit WIL Project")
	fmt.Println("3. Search WIL Project")
	fmt.Println("4. List all WIL Project")
	fmt.Println("5. Get WIL Project Detail By ID")
	fmt.Println("6. Delete WIL Project By ID")
	fmt.Println("0. Exit WIL Module")
}

func (menu *WILProjectMenuStateHandler) HandleUserInput(input string) error {
	switch input {
	case "1":
		fmt.Println("1 Not implemented yet...")
	case "2":
		fmt.Println("2 Not implemented yet...")
	case "3":
		fmt.Println("3 Not implemented yet...")
	case "4":
		fmt.Println("4 Not implemented yet...")
	case "5":
		fmt.Println("5 Not implemented yet...")
	case "6":
		fmt.Println("6 Not implemented yet...")
	case "0":
		menu.manager.SetState(menu.wilModuleMenuStateHandler)
		return nil
	default:
		fmt.Println("Invalid Command")
	}

	util.PressEnterToContinue()

	return nil
}
