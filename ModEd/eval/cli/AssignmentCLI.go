// mep 1006
package cli

import (
	"ModEd/eval/controller"
	"fmt"
	// "gorm.io/gorm"
)

type AssignmentCLI struct {
	controller controller.AssignmentController
}

func NewAssignmentCLI(controller controller.AssignmentController) *AssignmentCLI {
	return &AssignmentCLI{controller: controller}
}

func (cli *AssignmentCLI) Run() {
	for {
		var option int
		fmt.Println("\nQuiz CLI Menu")
		fmt.Println("1. List Assignment")
		fmt.Println("2. Create Assignment")
		fmt.Println("3. Update Assignment")
		fmt.Println("4. Delete Assignment")
		fmt.Println("5. Exit")
		fmt.Print("Choose an option: ")
		fmt.Scan(&option)

		switch option {
		case 1:
			cli.listAssignment()
		case 2:
			return
		default:
			fmt.Println("Invalid option")
		}
	}
}

func (cli *AssignmentCLI) listAssignment() {
	assignments, err := cli.controller.GetAll()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	for _, a := range assignments {
		fmt.Printf("ID: %d | Title: %s | Released: %t\n", a.AssignmentId, a.Title, a.Released)
	}
}
