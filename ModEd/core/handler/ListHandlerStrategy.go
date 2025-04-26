package handler

// Wrote by MEP-1012

import (
	"ModEd/core"
	"fmt"
)

type ListHandlerStrategy[T core.RecordInterface] struct {
	controller interface {
		List(condition map[string]interface{}) ([]T, error)
	}
}

func NewListHandlerStrategy[T core.RecordInterface](controller interface {
	List(condition map[string]interface{}) ([]T, error)
}) *ListHandlerStrategy[T] {
	return &ListHandlerStrategy[T]{controller: controller}
}

func (cs ListHandlerStrategy[T]) Execute() error {
	records, err := cs.controller.List(nil)
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
