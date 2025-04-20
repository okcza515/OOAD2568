package commands

import (
	"fmt"
	"strings"

	"gorm.io/gorm"
)

func (c *UpdateCommand) Execute(args []string, tx *gorm.DB) error {
	if len(args) < 1 {
		return fmt.Errorf("usage: update {student|instructor} [options]")
	}

	target := strings.ToLower(args[0])
	switch target {
	case "student":
		return updateStudent(args[1:], tx)
	case "instructor":
		return updateInstructor(args[1:], tx)
	default:
		return fmt.Errorf("unknown update target: %s", target)
	}
}
