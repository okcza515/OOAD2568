package commands

import (
	"ModEd/hr/controller"
	"flag"
	"fmt"

	"gorm.io/gorm"
)

type ListInstructorCommand struct{}

func (cmd *ListInstructorCommand) Execute(args []string, tx *gorm.DB) error {
	fs := flag.NewFlagSet("list-instructor", flag.ExitOnError)
	fs.Parse(args)

	instructorController := controller.NewInstructorHRController(tx)
	instructorInfos, err := instructorController.GetAllInstructors()
	if err != nil {
		return fmt.Errorf("error listing instructors: %v", err)
	}

	fmt.Println("Human Resource Instructor Info:")
	for _, i := range instructorInfos {
		fmt.Printf("IID: %s | Name: %s %s | Gender: %s | Email: %s | Department: %s | Phone: %s\n", i.InstructorCode, i.FirstName, i.LastName, i.Gender, i.Email, *i.Department, i.PhoneNumber)
	}
	return nil
}
