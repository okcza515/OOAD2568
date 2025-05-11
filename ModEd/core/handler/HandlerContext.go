package handler

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
)

// Wrote by MEP-1012, MEP-1010

type MenuStrategy interface {
	Execute() error
}

type FuncStrategy struct {
	Action func() error
}

func (f FuncStrategy) Execute() error {
	return f.Action()
}

type MenuItem struct {
	Label    string
	Strategy HandlerStrategy
}

type HandlerContext struct {
	title string
	menu  map[string]MenuItem
}

func NewHandlerContext() *HandlerContext {
	return &HandlerContext{
		menu: make(map[string]MenuItem),
	}
}

func (c *HandlerContext) HandleInput(userInput string) error {
	fmt.Println(c.menu[userInput].Label)

	handler, ok := c.menu[userInput]
	if !ok {
		fmt.Println("Invalid Command Input")
		return nil
	}

	if handler.Strategy == nil || handler.Strategy.Execute == nil {
		return errors.New("err : input handler not implemented")
	}

	return handler.Strategy.Execute()
}

func (c *HandlerContext) AddHandler(userInput string, headerLabel string, s HandlerStrategy) {
	_, exists := c.menu[userInput]
	if exists {
		return
	}

	c.menu[userInput] = MenuItem{
		Label:    headerLabel,
		Strategy: s,
	}
}

func (c *HandlerContext) ShowMenu() {
	fmt.Println(c.title)
	keys := make([]string, 0, len(c.menu))
	for key := range c.menu {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(i, j int) bool {
		iInt, errI := strconv.Atoi(keys[i])
		jInt, errJ := strconv.Atoi(keys[j])
		if errI == nil && errJ == nil {
			return iInt < jInt
		}
		return keys[i] < keys[j]
	})

	for _, key := range keys {
		menu := c.menu[key]
		fmt.Printf("  %s:\t%s\n", key, menu.Label)
	}
}

func (c *HandlerContext) SetMenuTitle(title string) {
	c.title = title
}

func (c *HandlerContext) AddBackHandler(strategy HandlerStrategy) {
	c.AddHandler("back", "exit to previous page", strategy)
}

// TODO for standard CLI
// Insert
// List with Pagination
// Retrieve by ID
// Filter
// Update
// Delete
