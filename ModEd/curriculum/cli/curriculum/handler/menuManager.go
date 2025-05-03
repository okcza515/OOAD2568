package handler

import (
	"ModEd/curriculum/utils"
	"errors"
)

var ExitCommand = errors.New("exit")

type MenuManager struct {
	Actions   map[string]func() error
	UserInput string
}

type MenuManagerInterface interface {
	Execute(choice string) error
	Run()
}

func NewMenuManager(actions map[string]func() error) *MenuManager {
	return &MenuManager{
		Actions: actions,
	}
}

func (m *MenuManager) Execute(choice string) error {
	return m.Actions[choice]()
}

func (m *MenuManager) AddAction(choice string, action func() error) {
	m.Actions[choice] = action
}

func (m *MenuManager) HandlerUserInput(printMenu func()) string {
	printMenu()
	m.UserInput = utils.GetUserChoice()
	return m.UserInput
}
