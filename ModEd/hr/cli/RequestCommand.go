package main

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
			return requestResignationStudent(args[2:], tx)
		case "leave":
			return requestLeaveStudent(args[2:], tx)
		default:
			return fmt.Errorf("unknown action for student: %s", action)
		}
	case "instructor":
		switch action {
		case "resign":
			return requestResignationInstructor(args[2:], tx)
		case "leave":
			return requestLeaveInstructor(args[2:], tx)
		case "raise":
			return requestRaiseInstructor(args[2:], tx)
		default:
			return fmt.Errorf("unknown action for instructor: %s", action)
		}
	default:
		return fmt.Errorf("unknown target: %s", target)
	}
}
