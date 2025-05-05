package cli

import "errors"

// Wrote by MEP-1012

type CLIMenuStateManager struct {
	currentMenuState MenuState

	UserInput    string
	unitedStates map[string]MenuState
}

func NewCLIMenuManager() *CLIMenuStateManager {
	us := make(map[string]MenuState)
	return &CLIMenuStateManager{
		currentMenuState: nil,
		UserInput:        "",
		unitedStates:     us,
	}
}

func (manager *CLIMenuStateManager) AddMenu(menuLabel string, newState MenuState) {
	manager.unitedStates[menuLabel] = newState
}

func (manager *CLIMenuStateManager) Render() {
	manager.currentMenuState.Render()
}

func (manager *CLIMenuStateManager) HandleUserInput() error {
	return manager.currentMenuState.HandleUserInput(manager.UserInput)
}

func (manager *CLIMenuStateManager) SetState(newState MenuState) {
	manager.currentMenuState = newState
}

func (manager *CLIMenuStateManager) GetState(menu string) MenuState {
	return manager.unitedStates[menu]
}

func (manager *CLIMenuStateManager) GoToMenu(menu string) error {
	ms, ok := manager.unitedStates[menu]

	if !ok {
		return errors.New("err: menu state '" + menu + "' is not registered")
	}

	manager.currentMenuState = ms
	return nil
}
