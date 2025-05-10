package utils

import (
	"fmt"
	"strconv"
	"strings"
)

type MenuComponentHandler interface {
	ExecuteMenuComponent(params []string)
	GetInputSequence() []InputPrompt
}

type InputPrompt struct {
	Label    string
	Validate func(string) error
}

type MenuComponent struct {
	Key     string
	Label   string
	Handler MenuComponentHandler
}

type MenuHandler struct {
	//defaultHandler MenuComponentHandler
	parentMenu  *MenuHandler         //a menu can go back to parent if go back (mostly used when canceling an action)
	backHandler MenuComponentHandler //a menu can go back(AKA. jump to) another method(Handler) if go back (special case when canceling an action)
	items       []MenuComponent
	menuTitle   string
}

func NewMenuHandler(title string) *MenuHandler {
	return &MenuHandler{
		items:     make([]MenuComponent, 0),
		menuTitle: title,
	}
}

func (handler *MenuHandler) AppendComponent(label string, itemHandler MenuComponentHandler) {
	key := strconv.Itoa(len(handler.items) + 1)
	handler.items = append(handler.items, MenuComponent{
		Key:     key,
		Label:   label,
		Handler: itemHandler,
	})
}

func (handler *MenuHandler) showMenu() {
	fmt.Printf("\n--- %s ---\n", handler.menuTitle)

	for _, item := range handler.items {
		fmt.Printf("%s. %s\n", item.Key, item.Label)
	}

	if handler.parentMenu != nil || handler.backHandler != nil {
		fmt.Println("-1. Back")
	}

	fmt.Print("Select an option: ")
	var input string
	fmt.Scanln(&input)
	input = strings.TrimSpace(input)

	if input == "-1" {
		if handler.parentMenu != nil {
			handler.parentMenu.showMenu()
			return
		}
		if handler.backHandler != nil {
			handler.backHandler.ExecuteMenuComponent(nil)
			return
		}
		return
	}

	for _, item := range handler.items {
		if item.Key == input {
			if subMenu, ok := item.Handler.(*MenuHandler); ok {
				subMenu.showMenu()
			} else {
				handler.collectAndExecuteInputs(item.Handler)
			}
			return
		}
	}

	fmt.Println("Invalid selection")
	handler.showMenu()
}

func (handler *MenuHandler) collectAndExecuteInputs(itemHandler MenuComponentHandler) {
	var inputs []string
	prompts := itemHandler.GetInputSequence()

	for _, prompt := range prompts {
		for {
			fmt.Printf("%s: ", prompt.Label)
			var input string
			fmt.Scanln(&input)
			input = strings.TrimSpace(input)

			if input == "-1" {
				handler.showMenu() // Return to current menu
				return
			}

			if err := prompt.Validate(input); err != nil {
				fmt.Printf("  Invalid input: %v\n", err)
				continue
			}

			inputs = append(inputs, input)
			break
		}
	}

	itemHandler.ExecuteMenuComponent(inputs)
	handler.showMenu() // return to menu after execution
}

func (handler *MenuHandler) ExecuteMenuComponent(_ []string) {
	handler.showMenu()
}

func (handler *MenuHandler) GetInputSequence() []InputPrompt {
	return nil
}

func (handler *MenuHandler) SetBackHandler(itemHandler MenuComponentHandler) {
	handler.backHandler = itemHandler
}

func (handler *MenuHandler) SetParentMenu(parent *MenuHandler) {
	handler.parentMenu = parent
}
