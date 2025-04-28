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
}

func NewRetrieveByIDHandlerStrategy[T core.RecordInterface](controller interface {
	RetrieveByID(id uint, preloads ...string) (T, error)
}) *RetrieveByIDHandlerStrategy[T] {
	return &RetrieveByIDHandlerStrategy[T]{controller: controller}
}

func (cs RetrieveByIDHandlerStrategy[T]) Execute() error {

	var id uint
	fmt.Print("Enter ID to retrieve: ")
	_, err := fmt.Scanln(&id)

	if err != nil {
		return err
	}

	record, err := cs.controller.RetrieveByID(id)

	if err != nil {
		return err
	}

	fmt.Println(record.ToString())

	return nil
}
