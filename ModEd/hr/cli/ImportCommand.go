package main

import (
	"fmt"
	"strings"

	"gorm.io/gorm"
)

func (c *ImportCommand) Execute(args []string, tx *gorm.DB) error {
	if len(args) < 1 {
		return fmt.Errorf("usage: import {student|instructor} [options]")
	}

	target := strings.ToLower(args[0])
	switch target {
	case "student":
		return importStudents(args[1:], tx)
	case "instructor":
		return importInstructor(args[1:], tx)
	default:
		return fmt.Errorf("unknown update target: %s", target)
	}
}
