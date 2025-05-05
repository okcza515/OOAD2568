package handler

import (
	"errors"
	"fmt"
)

// Wrote by MEP-1012

type HandlerContext struct {
	handlerMap map[string]HandlerStrategy
	labelMap   map[string]string
}

func NewHandlerContext() *HandlerContext {
	return &HandlerContext{
		handlerMap: make(map[string]HandlerStrategy),
		labelMap:   make(map[string]string),
	}
}

func (c *HandlerContext) HandleInput(userInput string) error {
	fmt.Println(c.labelMap[userInput])

	handler, ok := c.handlerMap[userInput]
	if !ok {
		fmt.Println("Invalid Command Input")
		return nil
	}

	if handler == nil {
		return errors.New("err : input handler not implemented")
	}

	return handler.Execute()
}

func (c *HandlerContext) AddHandler(userInput string, headerLabel string, s HandlerStrategy) {
	c.labelMap[userInput] = headerLabel
	c.handlerMap[userInput] = s
}

// TODO for standard CLI
// Insert
// List with Pagination
// Retrieve by ID
// Filter
// Update
// Delete
