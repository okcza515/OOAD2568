package utils

import (
	"fmt"
	"sort"
)

type MenuItem struct {
	Label   string
	Handler func() error
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

func (c *GeneralHandlerContext) HandleInput(menuNumber string) error {
	menuItem, ok := c.menu[menuNumber]
	if !ok {
		fmt.Println("Invalid Command Input")
		return nil
	}

	if menuItem.Handler == nil {
		fmt.Println("err : input handler not implemented")
		return nil
	}

	return menuItem.Handler()
}

func (c *GeneralHandlerContext) SetMenuTitle(title string) {
	c.title = title
}

func (c *GeneralHandlerContext) AddHandler(menuNumber string, label string, handler func() error) {
	_, ok := c.menu[menuNumber]
	if ok {
		return
	}

	c.menu[menuNumber] = MenuItem{
		Label:   label,
		Handler: handler,
	}
	return
}

func (c *GeneralHandlerContext) AddBackHandler(handler func() error) {
	_, ok := c.menu["back"]
	if ok {
		return
	}

	c.menu["back"] = MenuItem{
		Label:   "exit to previous page",
		Handler: handler,
	}
	return
}

func (c *GeneralHandlerContext) ShowMenu() error {
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
	return nil
}
