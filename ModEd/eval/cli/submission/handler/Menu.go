// MEP-1007
package handler

import (
	"fmt"
	"os"
)

type Executable interface {
	Execute()
}

type menuItem struct {
	label string
	exec  Executable
}

type MenuHandler struct {
	prompt         string
	items          []menuItem
	showBack       bool
	backHandler    Executable
	defaultHandler Executable
}

func NewMenuHandler(prompt string, showBack bool) *MenuHandler {
	return &MenuHandler{
		prompt:   prompt,
		showBack: showBack,
	}
}

func (m *MenuHandler) Add(label string, exec Executable) {
	m.items = append(m.items, menuItem{label: label, exec: exec})
}

func (m *MenuHandler) SetBackHandler(handler Executable) {
	m.backHandler = handler
}

func (m *MenuHandler) SetDefaultHandler(handler Executable) {
	m.defaultHandler = handler
}

func (m *MenuHandler) Execute() {
	for {
		fmt.Println("\n" + m.prompt)

		for i, item := range m.items {
			fmt.Printf("[%d] %s\n", i+1, item.label)
		}

		if m.showBack {
			fmt.Println("[b] Back")
		} else {
			fmt.Println("[x] Exit")
			os.Exit(0)
		}

		fmt.Print("Choose: ")
		var input string
		fmt.Scanln(&input)

		switch input {
		case "x":
			if !m.showBack {
				os.Exit(0)
			}
		case "b":
			if m.showBack {
				if m.backHandler != nil {
					m.backHandler.Execute()
				}
				return
			}
		default:
			var index int
			_, err := fmt.Sscanf(input, "%d", &index)
			if err == nil && index >= 1 && index <= len(m.items) {
				m.items[index-1].exec.Execute()
			} else if m.defaultHandler != nil {
				m.defaultHandler.Execute()
			} else {
				fmt.Println("Invalid option.")
			}
		}
	}
}

type DummyCommand struct{}

func (d DummyCommand) Execute() {
	fmt.Println("Empty handler, please implement")
}