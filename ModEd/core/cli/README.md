# Core Menu State

## CLIMenuStateManager Functions

| Function                       | Description                                            |
| ------------------------------ | ------------------------------------------------------ |
| `NewCLIMenuManager()`          | Creates a new menu manager instance.                   |
| `AddMenu(label string, state)` | Registers a new menu state under a label.              |
| `SetState(state)`              | Sets the current active state.                         |
| `GoToMenu(label string)`       | Switches to the menu state associated with the label.  |
| `Render()`                     | Calls the `Render` method of the current state.        |
| `HandleUserInput()`            | Passes the `UserInput` to the current state's handler. |
| `GetState(label string)`       | Retrieves a registered menu state by label.            |


## Adding a New Menu State

### To add a new menu:

1. Create a new state struct that implements:

```go
Render()
HandleUserInput(input string) error
```

2. Register the new menu

```go
menuManager.AddMenu("newOption", NewState)
```

```go
package main

import (
	"fmt"
	"core/cli"
)

type MenuState struct {
	menuManger *cli.CLIMenuStateManager

	State1 *State1
	State2 *State2
	State3 *State3
}

func NewMenuState(manager *cli.CLIMenuStateManager) *MenuState {
	menuState := &MenuState{
		menuManger: manager,
	}

	// Initialize the sub-states
	menuState.State1 = NewState1(manager, menuState)
	menuState.State2 = NewState2(manager, menuState)
	menuState.State3 = NewState3(manager, menuState)

	// Register menus
	menuState.menuManger.AddMenu("1", menuState.State1)
	menuState.menuManger.AddMenu("2", menuState.State2)
	menuState.menuManger.AddMenu("3", menuState.State3)
	menuState.menuManger.AddMenu("exit", nil)

	return menuState
}

func (menu *MenuState) Render() {
	fmt.Println("\nWIL Module Menu:")
	fmt.Println("1. WIL Project Curriculum")
	fmt.Println("2. WIL Project Application")
	fmt.Println("3. WIL Project")
	fmt.Println("4. Independent Study")
	fmt.Println("exit: Exit the module")
}

func (menu *MenuState) HandleUserInput(input string) error {
	err := menu.menuManger.GoToMenu(input)
	if err != nil {
		fmt.Println("err: Invalid input, menu '" + input + "' doesn't exist")
	}
	return nil
}

func main() {
	menuManager := cli.NewCLIMenuManager()
	menuState := NewMenuState(menuManager)
	menuManager.SetState(menuState)

	for {
		menuManager.Render()
		menuManager.UserInput = utils.GetUserChoice()

		if menuManager.UserInput == "exit" {
			break
		}

		err := menuManager.HandleUserInput()
		if err != nil {
			return
		}
	}
}
```