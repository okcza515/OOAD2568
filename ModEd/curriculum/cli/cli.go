package main

import (
	"ModEd/curriculum/cli/curriculum"
	"fmt"
)

// TODO: not sure is this a good approach to do at all might need to discuss further
func main() {
	for {
		displayMainMenu()
		choice := getUserChoice()

		if handleUserChoice(choice) {
			return
		}
	}
}

func displayMainMenu() {
	fmt.Println("\nModEd CLI Menu")
	fmt.Println("1. Curriculum")
	fmt.Println("2. Class")
	fmt.Println("3. Course")
	fmt.Println("0. Exit")
}

func getUserChoice() string {
	var choice string
	fmt.Print("Enter your choice: ")
	fmt.Scanln(&choice)
	return choice
}

func handleUserChoice(choice string) bool {
	switch choice {
	case "1":
		curriculum.RunCurriculumCLI()
		return true
	case "2":
		curriculum.RunClassCLI()
		return false
	case "3":
		curriculum.RunCourseCLI()
		return false
	case "0":
		fmt.Println("Exiting...")
		return true
	default:
		fmt.Println("Invalid option")
		return false
	}
}
