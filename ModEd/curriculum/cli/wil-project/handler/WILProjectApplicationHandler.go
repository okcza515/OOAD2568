package handler

import (
	"ModEd/curriculum/controller"
	"ModEd/curriculum/utils"
	"fmt"
)

func RunWILProjectApplicationHandler(controller *controller.WILProjectApplicationController) {
	for {
		printWILProjectApplicationModuleMenu()
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

func printWILProjectApplicationModuleMenu() {
	fmt.Println("\nWIL Project Application Menu:")
	fmt.Println("1. Create WIL Project Application")
	fmt.Println("2. Edit WIL Project Application")
	fmt.Println("3. Search WIL Project Application")
	fmt.Println("4. List all WIL Project Application")
	fmt.Println("5. Get WIL Project Application By ID")
	fmt.Println("6. Delete WIL Project Application")
	fmt.Println("0. Exit WIL Module")
}
