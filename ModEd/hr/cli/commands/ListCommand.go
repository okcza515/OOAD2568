package commands

import (
	"fmt"

	"gorm.io/gorm"
)

func (c *ListCommand) Execute(args []string, tx *gorm.DB) error {
	if len(args) < 1 {
		return fmt.Errorf("usage: list {student|instructor} [options]")
	}

	target := args[0]

	switch target {
	case "student":
		return listStudents(args[1:], tx)
	case "instructor":
		return nil // wait for implementation
	default:
		return fmt.Errorf("unknown target: %s", target)
	}
}
