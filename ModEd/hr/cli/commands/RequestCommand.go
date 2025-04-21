package commands

import (
	"fmt"
	"strings"

	"gorm.io/gorm"
)

func (c *RequestCommand) Execute(args []string, tx *gorm.DB) error {
    if len(args) < 2 {
        return fmt.Errorf("usage: request {student|instructor} {resign|leave} [options]")
    }

    target := strings.ToLower(args[0])
    action := strings.ToLower(args[1])

    switch target {
    case "student":
        switch action {
        case "resign":
            return RequestResignationStudent(args[2:], tx)
        case "leave":
            return nil // wait for implementation
        default:
            return fmt.Errorf("unknown action for student: %s", action)
        }
    case "instructor":
        switch action {
        case "resign":
            return RequestResignationInstructor(args[2:], tx)
		case "leave":
            return nil // wait for implementation
		case "raise":
			return nil // wait for implementation
        default:
            return fmt.Errorf("unknown action for instructor: %s", action)
        }
    default:
        return fmt.Errorf("unknown target: %s", target)
    }
}