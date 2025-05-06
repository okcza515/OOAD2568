package menu

import (
    "fmt"
    "ModEd/core/cli"
)

type HRMainMenuState struct {
    manager    *cli.CLIMenuStateManager
    menuActions map[HRMenuEnum]func() error
}

func NewHRMainMenuState(manager *cli.CLIMenuStateManager) *HRMainMenuState {
    state := &HRMainMenuState{
        manager: manager,
        menuActions: map[HRMenuEnum]func() error{
            MENU_ADD: func() error {
                return manager.GoToMenu("AddInstructor")
            },
            MENU_LIST: func() error {
                return manager.GoToMenu("ViewInstructors")
            },
            MENU_MIGRATE: func() error {
                fmt.Println("Migrating data...")
                return nil
            },
            "exit": func() error {
                fmt.Println("Exiting HR module...")
                return nil
            },
        },
    }
    return state
}

func (state *HRMainMenuState) Render() {
    fmt.Println("=== HR Main Menu ===")
    fmt.Println("1. Add Personnel")
    fmt.Println("2. List Personnel")
    fmt.Println("3. Migrate Data")
    fmt.Println("0. Exit")
    fmt.Print("Enter your choice: ")
}

func (state *HRMainMenuState) HandleUserInput(input string) error {
    // Map user input to the corresponding HRMenuEnum
    var menuOption HRMenuEnum
    switch input {
    case "1":
        menuOption = MENU_ADD
    case "2":
        menuOption = MENU_LIST
    case "3":
        menuOption = MENU_MIGRATE
    case "0":
        menuOption = "exit"
    default:
        fmt.Println("Invalid choice. Please try again.")
        return nil
    }

    // Execute the corresponding action from the map
    if action, exists := state.menuActions[menuOption]; exists {
        return action()
    }

    fmt.Println("Invalid choice. Please try again.")
    return nil
}