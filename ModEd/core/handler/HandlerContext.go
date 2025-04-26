package handler

import "errors"

type HandlerContext struct {
	strategy HandlerStrategy
}

func (c *HandlerContext) Execute() error {
	if c.strategy == nil {
		return errors.New("err: strategy is not set")
	}

	return c.strategy.Execute()
}

func (c *HandlerContext) SetStrategy(s HandlerStrategy) {
	c.strategy = s
}

// TODO for standard CLI
// Insert
// List with Pagination
// Retrieve by ID
// Filter
// Update
// Delete
