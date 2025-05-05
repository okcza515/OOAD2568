package handler

// Wrote by MEP-1012

import (
	"ModEd/core/cli"
)

type ChangeMenuHandlerStrategy struct {
	menuManager *cli.CLIMenuStateManager
	nextState   cli.MenuState
}

func NewChangeMenuHandlerStrategy(
	menuManager *cli.CLIMenuStateManager,
	nextState cli.MenuState,
) *ChangeMenuHandlerStrategy {
	return &ChangeMenuHandlerStrategy{
		menuManager: menuManager,
		nextState:   nextState,
	}
}

func (handler ChangeMenuHandlerStrategy) Execute() error {
	handler.menuManager.SetState(handler.nextState)
	return nil
}
