package utils

import (
	"ModEd/core"
	"fmt"
	"io"
	"os"
	"strconv"
)

type MenuItem struct {
	Title    string
	Action   func(io *core.MenuIO)
	Children []*MenuItem
}

func (m *MenuItem) AddChild(child *MenuItem) {
	m.Children = append(m.Children, child)
}

type MenuBuilder struct {
	Root *MenuItem
	IO   *core.MenuIO
}

func NewMenuBuilder(root *MenuItem, reader io.Reader, writer io.Writer) *MenuBuilder {
	if reader == nil {
		reader = os.Stdin
	}
	if writer == nil {
		writer = os.Stdout
	}

	return &MenuBuilder{
		Root: root,
		IO:   core.NewMenuIO(),
	}
}

func (mb *MenuBuilder) AddMenuPath(path []string, action func(io *core.MenuIO)) {
	current := mb.Root

	for _, title := range path {
		found := false
		for _, child := range current.Children {
			if child.Title == title {
				current = child
				found = true
				break
			}
		}

		if !found {
			newItem := &MenuItem{Title: title}
			current.AddChild(newItem)
			current = newItem
		}
	}

	current.Action = action
}

func (mb *MenuBuilder) AddMenuChild(path []string, menu *MenuItem) {
	current := mb.Root
	for _, title := range path {
		found := false
		for _, child := range current.Children {
			if child.Title == title {
				current = child
				found = true
				break
			}
		}

		if !found {
			newItem := &MenuItem{Title: title}
			current.AddChild(newItem)
			current = newItem
		}
	}

	for _, child := range current.Children {
		if child.Title == menu.Title {
			return
		}
	}

	current.AddChild(menu)
}

func (mb *MenuBuilder) Show() {
	mb.showMenu(mb.Root, "")
}

func (mb *MenuBuilder) showMenu(item *MenuItem, path string) {
	for {
		fmt.Fprintf(mb.IO.Writer, "\n%s> %s\n", path, item.Title)
		for i, child := range item.Children {
			fmt.Fprintf(mb.IO.Writer, "[%d] %s\n", i+1, child.Title)
		}
		if path != "" {
			fmt.Fprint(mb.IO.Writer, "[0] Back\n")
		} else {
			fmt.Fprint(mb.IO.Writer, "[0] Exit\n")
		}

		fmt.Fprint(mb.IO.Writer, "Select an option: ")
		choice, _ := mb.IO.ReadInput()

		index, err := strconv.Atoi(choice)
		if err != nil || index < 0 || index > len(item.Children) {
			mb.IO.Println("Invalid input, try again.")
			continue
		}

		if index == 0 {
			break
		}

		selected := item.Children[index-1]
		if len(selected.Children) > 0 {
			mb.showMenu(selected, path+"/"+selected.Title)
		} else if selected.Action != nil {
			selected.Action(mb.IO)
		} else {
			mb.IO.Println("No action assigned.")
		}
	}
}
