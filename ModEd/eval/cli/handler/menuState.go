// MEP-1007
package handler
import (
	"fmt"
)

type MainMenuState struct {
	username   string
}

type MenuState interface {
    Enter() error
    Exit() error
    HandleInput(input string) (MenuState, error)
}


func (m *MainMenuState) Enter() error {
	fmt.Println("\n=== Main Menu ===")
	fmt.Println("2. Exit")
	return nil
}

func (m *MainMenuState) Exit() error {
	fmt.Println("Exiting Main Menu...")
	return nil
}

func (m *MainMenuState) HandleInput(input string) (MenuState, error) {
	fmt.Println("You entered:", input)

	switch input {
	case "2":
		fmt.Println("Exiting...")
		return nil, nil 
	default:
		return m, nil 
	}
}

