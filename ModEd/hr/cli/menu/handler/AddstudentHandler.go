package handler

import (
	"ModEd/core"
	"fmt"
)

type AddStudentStrategy[T core.RecordInterface] struct {
	controller interface{ InsertOne(data interface{}) error }
}

func NewAddStudentStrategy[T core.RecordInterface](controller interface{ InsertOne(data interface{}) error }) *AddStudentStrategy[T] {
	return &AddStudentStrategy[T]{controller: controller}
}

func (handler AddStudentStrategy[T]) Execute() error {
	id := ""
	fmt.Printf("Please enter the student ID: ")
	fmt.Scanln(&id)

	name := ""
	fmt.Printf("Please enter the student name: ")
	fmt.Scanln(&name)

	return nil
}
