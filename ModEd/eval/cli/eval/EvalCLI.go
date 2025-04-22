// mep 1006
package main

import (
	"ModEd/eval/cli"
	"fmt"
	// "gorm.io/gorm"
)

func main() {
	fmt.Println("Hello from Main")

	cli.PrintAssignment()
	cli.PrintEvaluation()
}
