package commands

import (
	"ModEd/hr/controller"
	"flag"
	"fmt"

	"gorm.io/gorm"
)

type PullStudentCommand struct{}

func (c *PullStudentCommand) Execute(args []string, tx *gorm.DB) error {
	fs := flag.NewFlagSet("pull-student", flag.ExitOnError)
	fs.Parse(args)

	studentController := controller.NewStudentHRController(tx)
	if err := studentController.MigrateStudentRecords(); err != nil {
		return fmt.Errorf("failed to pull student record into studentinfo: %w", err)
	}

	fmt.Println("Student record pulled successfully into StudentInfo table.")
	return nil
}
