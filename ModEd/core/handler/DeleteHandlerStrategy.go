package handler

// Wrote by MEP-1012

import (
	"ModEd/core"
	"fmt"
	"strconv"
)

type DeleteHandlerStrategy[T core.RecordInterface] struct {
	controller interface {
		DeleteByID(id uint) error
		List(condition map[string]interface{}, preloads ...string) ([]T, error)
	}
}

func NewDeleteHandlerStrategy[T core.RecordInterface](controller interface {
	DeleteByID(id uint) error
	List(condition map[string]interface{}, preloads ...string) ([]T, error)
}) *DeleteHandlerStrategy[T] {
	return &DeleteHandlerStrategy[T]{controller: controller}
}

func (handler DeleteHandlerStrategy[T]) Execute() error {

	records, err := handler.controller.List(nil)
	if err != nil {
		return err
	}

	fmt.Println(fmt.Sprintf("Total %v record(s)", len(records)))
	fmt.Println()

	for _, record := range records {
		fmt.Println(record.ToString())
	}

	fmt.Print("Enter ID to delete: ")
	var inputBuffer string
	fmt.Scanln(&inputBuffer)

	idNum, err := strconv.Atoi(inputBuffer)
	if err != nil {
		fmt.Println("Invalid ID format. Please enter a valid number.")
		return nil
	}
	err = handler.controller.DeleteByID(uint(idNum))
	if err != nil {
		fmt.Println("Error deleting record:", err)
		return nil
	}
	fmt.Println("Record successfully deleted!")
	return nil
}
