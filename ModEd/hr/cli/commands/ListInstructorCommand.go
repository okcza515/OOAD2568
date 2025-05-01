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

	instructorInfos, err := controller.GetAllInstructors(tx)
	if err != nil {
		return fmt.Errorf("error listing instructors: %v", err)
	}

	// TODO: Add more fields to display
	fmt.Println("Human Resource Instructor Info:")
	for _, i := range instructorInfos {
		fmt.Printf("IID: %s | Name: %s %s |", i.InstructorCode, i.FirstName, i.LastName)
	}
	return nil
}
