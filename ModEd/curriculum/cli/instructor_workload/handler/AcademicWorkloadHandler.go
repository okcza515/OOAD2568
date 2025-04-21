package handler

import (
	"ModEd/curriculum/utils"
	"fmt"
)

func RunAcademicWorkloadHandler() {
	for {
		DisplayAcademicWorkloadModuleMenu()
		choice := utils.GetUserChoice()
		fmt.Println("choice: ", choice)

		switch choice {
		case "1":
			fmt.Println("Add Class Lecture Not implemented yet...")
		case "2":
			fmt.Println("Edit Class Lecture Not implemented yet...")
		case "3":
			fmt.Println("Delete Class Lecture Not implemented yet...")
		case "4":
			fmt.Println("List all Class Lectures Not implemented yet...")
		case "5":
			fmt.Println("Get Class Lecture By ID Not implemented yet...")
		case "6":
			fmt.Println("Add Class Material Not implemented yet...")
		case "7":
			fmt.Println("Edit Class Material Not implemented yet...")
		case "8":
			fmt.Println("Delete Class Material Not implemented yet...")
		case "9":
			fmt.Println("List all Class Materials Not implemented yet...")
		case "10":
			fmt.Println("Get Class Material By ID Not implemented yet...")
		case "exit":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid option")
		}
	}
}

func DisplayAcademicWorkloadModuleMenu() {
	fmt.Println("\nAcademic Workload Menu:")
	fmt.Println("1. Add Class Lecture")
	fmt.Println("2. Edit Class Lecture")
	fmt.Println("3. Delete Class Lecture")
	fmt.Println("4. List all Class Lectures")
	fmt.Println("5. Get Class Lecture By ID")

	fmt.Println("6. Add Class Material")
	fmt.Println("7. Edit Class Material")
	fmt.Println("8. Delete Class Material")
	fmt.Println("9. List all Class Materials")
	fmt.Println("10. Get Class Material By ID")

	fmt.Println("Type 'exit' to quit")
}
