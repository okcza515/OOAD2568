package utils

import (
	"fmt"
	"sort"
)

//Wrote by MEP-1010

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
	Strategy MenuStrategy
}

type GeneralHandlerContext struct {
	title string
	menu  map[string]MenuItem
}

func NewGeneralHandlerContext() *GeneralHandlerContext {
	return &GeneralHandlerContext{
		menu: make(map[string]MenuItem),
	}
}

func (c *GeneralHandlerContext) SetMenuTitle(title string) {
	c.title = title
}

func (c *GeneralHandlerContext) AddHandler(menuNumber string, label string, strategy MenuStrategy) {
	_, exists := c.menu[menuNumber]
	if exists {
		return
	}

	c.menu[menuNumber] = MenuItem{
		Label:    label,
		Strategy: strategy,
	}
}

func (c *GeneralHandlerContext) AddBackHandler(strategy MenuStrategy) {
	c.AddHandler("back", "exit to previous page", strategy)
}

func (c *GeneralHandlerContext) ShowMenu() {
	fmt.Println(c.title)
	keys := make([]string, 0, len(c.menu))
	for key := range c.menu {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	for _, key := range keys {
		menu := c.menu[key]
		fmt.Printf("%s: %s\n", key, menu.Label)
	}
}

func (c *GeneralHandlerContext) HandleInput(menuNumber string) error {
	menuItem, exists := c.menu[menuNumber]
	if !exists {
		fmt.Println("Invalid Command Input")
		return nil
	}

	if menuItem.Strategy == nil || menuItem.Strategy.Execute == nil {
		fmt.Println("err : input handler not implemented")
		return nil
	}

	return menuItem.Strategy.Execute()
}
