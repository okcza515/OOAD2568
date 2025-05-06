package menu

// MEP-1012 Asset

import (
	"ModEd/asset/controller"
	"ModEd/asset/model"
	"ModEd/asset/util"
	"ModEd/core/cli"
	"ModEd/core/handler"
	"fmt"
)

type CategoryMenuState struct {
	manager        *cli.CLIMenuStateManager
	handlerContext *handler.HandlerContext
}

func NewCategoryMenuState(
	manager *cli.CLIMenuStateManager,
) *CategoryMenuState {
	controllerInstance := controller.GetAssetInstance().Category

	handlerContext := handler.NewHandlerContext()

	insertHandler := handler.NewInsertHandlerStrategy[model.Category](controllerInstance)
	listHandler := handler.NewListHandlerStrategy[model.Category](controllerInstance)
	updateHandler := handler.NewUpdateHandlerStrategy[model.Category](controllerInstance)
	deleteHandler := handler.NewDeleteHandlerStrategy[model.Category](controllerInstance)
	backHandler := handler.NewChangeMenuHandlerStrategy(manager, manager.GetState(string(MENU_ASSET)))

	handlerContext.AddHandler("1", "Add New Category", insertHandler)
	handlerContext.AddHandler("2", "List all Category", listHandler)
	handlerContext.AddHandler("3", "Get full detail of an Category", nil)
	handlerContext.AddHandler("4", "Update an Category", updateHandler)
	handlerContext.AddHandler("5", "Delete an Category", deleteHandler)
	handlerContext.AddHandler("back", "Back to main menu", backHandler)

	return &CategoryMenuState{
		manager:        manager,
		handlerContext: handlerContext,
	}
}

func (menu *CategoryMenuState) Render() {
	fmt.Println()
	fmt.Println(":/asset/Category")
	fmt.Println()
	fmt.Println("Category Management")
	fmt.Println("Your options are...")
	menu.handlerContext.ShowMenu()
	fmt.Println("  exit:\tExit the program (or Ctrl+C is fine ¯\\\\_(ツ)_/¯)")
	fmt.Println()
}

func (menu *CategoryMenuState) HandleUserInput(input string) error {
	fmt.Println(menu.handlerContext)
	err := menu.handlerContext.HandleInput(input)
	if err != nil {
		fmt.Println(err)
	}

	if input != "back" {
		util.PressEnterToContinue()
	}

	return nil
}
