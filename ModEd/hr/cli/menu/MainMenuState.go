package menu

import (
	"ModEd/core/cli"
	"ModEd/core/handler"
	"fmt"
)

type HRMainMenuState struct {
	manager        *cli.CLIMenuStateManager
	handlerContext *handler.HandlerContext
}

func NewHRMainMenuState(manager *cli.CLIMenuStateManager) *HRMainMenuState {
	handlerContext := handler.NewHandlerContext()
	state := &HRMainMenuState{
		manager:        manager,
		handlerContext: handlerContext,
	}

    manager.AddMenu(string(MENU_ADD), NewAddMenuState(manager))
	return state
}

func (state *HRMainMenuState) Render() {
	fmt.Println("=== HR Main Menu ===")
	fmt.Println("1. Add Personnel")
	fmt.Println("2. List Personnel")
	// implement the remaining menu options
	fmt.Println("0. Exit")
	fmt.Print("Enter your choice: ")
}

// func (state *HRMainMenuState) HandleUserInput(input string) error {
// 	// Map user input to the corresponding HRMenuEnum
// 	var menuOption HRMenuEnum
// 	switch input {
// 	case "1":
// 		menuOption = MENU_ADD
// 	case "2":
// 		menuOption = MENU_LIST
// 	case "3":
// 		menuOption = MENU_MIGRATE
// 	case "0":
// 		menuOption = "exit"
// 	default:
// 		fmt.Println("Invalid choice. Please try again.")
// 		return nil
// 	}

// 	// Execute the corresponding action from the map
// 	if action, exists := state.menuActions[menuOption]; exists {
// 		return action()
// 	}

// 	fmt.Println("Invalid choice. Please try again.")
// 	return nil
// }
