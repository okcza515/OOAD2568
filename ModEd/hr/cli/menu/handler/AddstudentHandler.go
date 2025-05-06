package handler

import (
	"ModEd/core"
	"ModEd/hr/controller"
	"fmt"

	"gorm.io/gorm"
)

type AddStudentStrategy[T core.RecordInterface] struct {
	controller interface{ InsertOne(data interface{}) error }
}

func NewAddStudentStrategy[T core.RecordInterface](controller interface{ InsertOne(data interface{}) error }) *AddStudentStrategy[T] {
	return &AddStudentStrategy[T]{controller: controller}
}

func (handler AddStudentStrategy[T]) Execute(db *gorm.DB) error {
	studentCode := core.GetUserInput("Enter student ID: ")
	firstName := core.GetUserInput("Enter student first name: ")
	lastName := core.GetUserInput("Enter student last name: ")
	email := core.GetUserInput("Enter student email: ")
	startDate := core.GetUserInput("Enter student start date: ")
	birthDate := core.GetUserInput("Enter student birth date: ")
	program := core.GetUserInput("Enter student program: ")
	department := core.GetUserInput("Enter student department: ")
	status := core.GetUserInput("Enter student status: ")
	gender := core.GetUserInput("Enter student gender: ")
	citizenID := core.GetUserInput("Enter student citizen ID: ")
	phoneNumber := core.GetUserInput("Enter student phone number: ")
	advisorCode := core.GetUserInput("Enter student advisor code: ")

	studentController := controller.NewStudentHRController(db)
	err := studentController.AddStudent(
		studentCode,
		firstName,
		lastName,
		email,
		startDate,
		birthDate,
		program,
		department,
		status,
		gender,
		citizenID,
		phoneNumber,
		advisorCode,
	)

	if err != nil {
		return fmt.Errorf("failed to add student: %w", err)
	}

	fmt.Println("Student added and HR info updated successfully!")

	return nil
}
