package handler

// Wrote by MEP-1012

import (
	"ModEd/core"
	"ModEd/utils/deserializer"
	"fmt"
)

type InsertHandlerStrategy[T core.RecordInterface] struct {
	controller interface{ InsertMany(data []T) error }
}

func NewInsertHandlerStrategy[T core.RecordInterface](controller interface{ InsertMany(data []T) error }) *InsertHandlerStrategy[T] {
	return &InsertHandlerStrategy[T]{controller: controller}
}

func (cs InsertHandlerStrategy[T]) Execute() error {
	path := ""
	fmt.Println("Please enter the path of the records (csv or json): ")
	_, err := fmt.Scanln(&path)
	if err != nil {
		return err
	}

	var recordModel []T

	fd, err := deserializer.NewFileDeserializer(path)
	if err != nil {
		return err
	}
	err = fd.Deserialize(&recordModel)
	if err != nil {
		return err
	}

	err = cs.controller.InsertMany(recordModel)
	if err != nil {
		return err
	}

	fmt.Println("Records successfully inserted!")

	return nil
}
