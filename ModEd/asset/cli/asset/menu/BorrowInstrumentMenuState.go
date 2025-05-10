package menu

import (
	"ModEd/asset/controller"
	"ModEd/asset/model"
	"ModEd/asset/util"
	"ModEd/core/cli"
	"ModEd/core/handler"
	"fmt"
)

// MEP-1012 Asset

type BorrowInstrumentMenuState struct {
	manager        *cli.CLIMenuStateManager
	handlerContext *handler.HandlerContext
}

func NewBorrowInstrumentMenuState(
	manager *cli.CLIMenuStateManager,
) *BorrowInstrumentMenuState {
	controllerInstance := controller.GetAssetInstance().BorrowInstrument

	handlerContext := handler.NewHandlerContext()

	insertHandler := handler.NewInsertHandlerStrategy[model.BorrowInstrument](controllerInstance)
	listHandler := handler.NewListHandlerStrategy[model.BorrowInstrument](controllerInstance)
	//returnHandler := handler.NewReturnHandlerStrategy[model.BorrowInstrument](controllerInstance)
	deleteHandler := handler.NewDeleteHandlerStrategy[model.BorrowInstrument](controllerInstance)
	backHandler := handler.NewChangeMenuHandlerStrategy(manager, manager.GetState(string(MENU_ASSET)))

	handlerContext.AddHandler("1", "Add New BorrowInstrument", insertHandler)
	handlerContext.AddHandler("2", "List all BorrowInstrument", listHandler)
	handlerContext.AddHandler("3", "Get full detail of an BorrowInstrument", nil)
	//handlerContext.AddHandler("4", "return an BorrowInstrument", updateHandler)
	//

	handlerContext.AddHandler("5", "Delete an BorrowInstrument", deleteHandler)
	handlerContext.AddHandler("back", "Back to main menu", backHandler)

	return &BorrowInstrumentMenuState{
		manager:        manager,
		handlerContext: handlerContext,
	}
}

func (menu *BorrowInstrumentMenuState) Render() {
	fmt.Println()
	fmt.Println(":/asset/BorrowInstrumentMenuState")
	fmt.Println()
	fmt.Println("BorrowInstrument Management")
	fmt.Println("Your options are...")
	menu.handlerContext.ShowMenu()
	fmt.Println("  exit:\tExit the program (or Ctrl+C is fine ¯\\\\_(ツ)_/¯)")
	fmt.Println()
}

func (menu *BorrowInstrumentMenuState) HandleUserInput(input string) error {
	err := menu.handlerContext.HandleInput(input)
	if err != nil {
		fmt.Println(err)
	}

	if input != "back" {
		util.PressEnterToContinue()
	}

	return nil
}
