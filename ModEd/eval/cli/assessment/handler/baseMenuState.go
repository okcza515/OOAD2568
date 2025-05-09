package handler

import (
	"fmt"
)

type MenuItem struct {
	Key         string
	Description string
	Action      func() (MenuState, error)
}

type BaseMenuState struct {
	Name   string
	Items  []MenuItem
	Parent MenuState
}

// NewBaseMenuState creates a new base menu state
func NewBaseMenuState(name string, parent MenuState) *BaseMenuState {
	return &BaseMenuState{
		Name:   name,
		Parent: parent,
		Items:  []MenuItem{},
	}
}

// AddMenuItem adds a menu item to this state
func (s *BaseMenuState) AddMenuItem(key, description string, action func() (MenuState, error)) {
	s.Items = append(s.Items, MenuItem{
		Key:         key,
		Description: description,
		Action:      action,
	})
}

// AddBackItem adds a back item to this menu
func (s *BaseMenuState) AddBackItem() {
	s.AddMenuItem("0", "Back", func() (MenuState, error) {
		return s.Parent, BackCommand
	})
}

// AddExitItem adds an exit item to this menu
func (s *BaseMenuState) AddExitItem() {
	s.AddMenuItem("0", "Exit", func() (MenuState, error) {
		return nil, ExitCommand
	})
}

// Display shows the menu
func (s *BaseMenuState) Display() {
	fmt.Printf("\n%s Menu:\n", s.Name)
	for _, item := range s.Items {
		fmt.Printf("%s. %s\n", item.Key, item.Description)
	}
}

func (s *BaseMenuState) HandleInput(input string) (MenuState, error) {
	for _, item := range s.Items {
		if item.Key == input {
			return item.Action()
		}
	}
	fmt.Println("Invalid option")
	return nil, nil
}

func (s *BaseMenuState) GetName() string {
	return s.Name
}
