package commands

import (
	"ModEd/hr/controller"
	"flag"
	"fmt"

	"gorm.io/gorm"
)

type PullInstructorCommand struct{}

func (c *PullInstructorCommand) Execute(args []string, tx *gorm.DB) error {
	fs := flag.NewFlagSet("pull-instructor", flag.ExitOnError)
	fs.Parse(args)

	instructorController := controller.NewInstructorHRController(tx)
	if err := instructorController.MigrateInstructorRecords(); err != nil {
		return fmt.Errorf("failed to pull instructor record into instructorinfo: %w", err)
	}

	fmt.Println("Instructor record pulled successfully into InstructorInfo table.")
	return nil
}
