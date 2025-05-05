package handler

import (
	"ModEd/core"
	"fmt"
)

// Wrote by MEP-1012

type RetrieveByIDHandlerStrategy[T core.RecordInterface] struct {
	controller interface {
		RetrieveByID(id uint, preloads ...string) (T, error)
	}
	preloads []string
}

func NewRetrieveByIDHandlerStrategy[T core.RecordInterface](
	controller interface {
		RetrieveByID(id uint, preloads ...string) (T, error)
	},
	preloads ...string,
) *RetrieveByIDHandlerStrategy[T] {
	return &RetrieveByIDHandlerStrategy[T]{controller: controller, preloads: preloads}
}

func (handler RetrieveByIDHandlerStrategy[T]) Execute() error {

	var id uint
	fmt.Print("Enter ID to retrieve: ")
	_, err := fmt.Scanln(&id)

	if err != nil {
		return err
	}

	record, err := handler.controller.RetrieveByID(id, handler.preloads...)

	if err != nil {
		if err.Error() == "record not found" {
			fmt.Println("Record not found.")
			return nil
		}
		return err
	}

	fmt.Println(record.ToString())

	return nil
}
