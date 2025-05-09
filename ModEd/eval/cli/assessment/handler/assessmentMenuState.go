package handler

import (
	"fmt"
)

// AssessmentMenuState represents the main menu of the Assessment CLI
type AssessmentMenuState struct {
	params    *AssessmentCLIParams
	menuItems []MenuItem
}

// Add a method to add menu items
func (s *AssessmentMenuState) AddMenuItem(key, description string, action func() (MenuState, error)) {
	s.menuItems = append(s.menuItems, MenuItem{
		Key:         key,
		Description: description,
		Action:      action,
	})
}

// Enter displays the main menu
func (s *AssessmentMenuState) Enter() error {
	fmt.Println("\n===== Assessment Management =====")
	fmt.Println("1. List Assessments")
	fmt.Println("2. Create Assessment")
	fmt.Println("3. Update Assessment")
	fmt.Println("4. Delete Assessment")
	fmt.Println("5. Manage Submissions")
	fmt.Println("6. Manage Results")
	fmt.Println("back - Return to previous menu")
	return nil
}

// Exit handles exit from the main menu
func (s *AssessmentMenuState) Exit() error {
	return nil
}

// Refactor HandleInput to use menu items
func (s *AssessmentMenuState) HandleInput(input string) (MenuState, error) {
	for _, item := range s.menuItems {
		if item.Key == input {
			return item.Action()
		}
	}
	return s, fmt.Errorf("invalid choice: %s", input)
}
