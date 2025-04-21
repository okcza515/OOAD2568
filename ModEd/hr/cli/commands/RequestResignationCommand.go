package commands

// import (
// 	"fmt"
// 	"strings"

// 	"gorm.io/gorm"
// )

// func (c *RequestResignationCommand) Execute(args []string, tx *gorm.DB) error {
// 	if len(args) < 1 {
// 		return fmt.Errorf("usage: requestResignation {student|instructor} [options]")
// 	}

// 	target := strings.ToLower(args[0])
// 	switch target {
// 	case "student":
// 		return RequestResignationStudent(args[1:], tx)
// 	case "instructor":
// 		return RequestResignationInstructor(args[1:], tx)
// 	default:
// 		return fmt.Errorf("unknown requestResignation target: %s", target)
// 	}
// }
