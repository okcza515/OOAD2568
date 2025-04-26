package cli

// Wrote by MEP-1012

type CLIMenuStateManager struct {
	currentMenuState MenuState

	UserInput string
}

func NewCLIMenuManager() *CLIMenuStateManager {
	return &CLIMenuStateManager{currentMenuState: nil, UserInput: ""}
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
