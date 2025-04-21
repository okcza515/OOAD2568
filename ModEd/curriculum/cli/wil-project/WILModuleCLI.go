package wilproject

import (
	"ModEd/curriculum/utils"
	"fmt"
)

func RunWILModuleCLI() {
	fmt.Println("test")
	for {
		printWILModuleMenu()
		choice := utils.GetUserChoice()

		switch choice {
		case "1":
			fmt.Println("1 Not implemented yet...")
		case "2":
			fmt.Println("2 Not implemented yet...")
		case "0":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid option")
		}
	}
}

func printWILModuleMenu() {
	fmt.Println("\nWIL Module Menu:")
	fmt.Println("1. WIL Project")
	fmt.Println("2. WIL Project Application")
	fmt.Println("0. Exit WIL Module")
}
