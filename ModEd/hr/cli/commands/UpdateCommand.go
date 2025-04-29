package commands

import (
	"fmt"
	"strings"

	"gorm.io/gorm"
)

func (c *UpdateCommand) Execute(args []string, tx *gorm.DB) error {
	if len(args) < 1 {
		return fmt.Errorf("usage: update {student|instructor} {info|status} [options]")
	}

	target := strings.ToLower(args[0])
	action := strings.ToLower(args[1])
	switch target {
	case "student":
		switch action {
		case "info":
			return updateStudentInfo(args[2:], tx)
		case "status":
			return updateStudentStatus(args[2:], tx)
		default:
			return fmt.Errorf("unknown action for student: %s", action)
		}
	case "instructor":
		switch action {
		case "info":
			return updateInstructorInfo(args[2:], tx)
		default:
			return fmt.Errorf("unknown action for instructor: %s", action)
		}
	default:
		return fmt.Errorf("unknown update target: %s", target)
	}
}
