package handler

// Wrote by MEP-1012

import (
	"ModEd/core"
	"fmt"
)

type ListHandlerStrategy[T core.RecordInterface] struct {
	controller interface {
		List(condition map[string]interface{}, preloads ...string) ([]T, error)
	}

	preloads []string
}

func NewListHandlerStrategy[T core.RecordInterface](
	controller interface {
		List(condition map[string]interface{}, preloads ...string) ([]T, error)
	},
	preloads ...string,
) *ListHandlerStrategy[T] {
	return &ListHandlerStrategy[T]{controller: controller, preloads: preloads}
}

func (handler ListHandlerStrategy[T]) Execute() error {
	records, err := handler.controller.List(nil, handler.preloads...)
	if err != nil {
		return err
	}

	fmt.Println(fmt.Sprintf("Total %v record(s)", len(records)))
	fmt.Println()

	for _, record := range records {
		fmt.Println(record.ToString())
	}

	return nil
}
