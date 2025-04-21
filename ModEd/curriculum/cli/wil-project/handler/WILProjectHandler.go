package handler

import (
	"ModEd/curriculum/controller"
	"ModEd/curriculum/utils"
	"fmt"
)

func RunWILProjectHandler(controller *controller.WILProjectController) {
	for {
		printWILProjectModuleMenu()
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

func printWILProjectModuleMenu() {
	fmt.Println("\nWIL Project Menu:")
	fmt.Println("1. Create WIL Project")
	fmt.Println("2. Edit WIL Project")
	fmt.Println("3. Search WIL Project")
	fmt.Println("4. List all WIL Project")
	fmt.Println("5. Get WIL Project Detail By ID")
	fmt.Println("6. Delete WIL Project By ID")
	fmt.Println("0. Exit WIL Module")
}
