package handler

import (
	"ModEd/core"
	"ModEd/utils/deserializer"
	"fmt"
)

type UpdateHandlerStrategy[T core.RecordInterface] struct {
	controller interface{ UpdateByID(data T) error }
}

func NewUpdateHandlerStrategy[T core.RecordInterface](controller interface{ UpdateByID(data T) error }) *UpdateHandlerStrategy[T] {
	return &UpdateHandlerStrategy[T]{controller: controller}
}

func (cs UpdateHandlerStrategy[T]) Execute() error {
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

	err = cs.controller.UpdateByID(recordModel[0])
	if err != nil {
		return err
	}

	fmt.Printf("New record: %v", recordModel[0].ToString())
	fmt.Println("Records successfully updated!")

	return nil
}
