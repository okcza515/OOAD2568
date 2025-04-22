//MEP-1006

package main

import (
	"ModEd/eval/cli/assignment"
	"ModEd/eval/cli/evaluation"
	"fmt"
	// "gorm.io/gorm"
)

func main() {
	fmt.Println("Hello from Manager")
	assignment.PrintAssignment()
	evaluation.PrintEvaluation()
}
