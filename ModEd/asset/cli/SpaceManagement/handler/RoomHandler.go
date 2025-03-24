package handler

import (
	controller "ModEd/asset/controller/spacemanagement"
	"ModEd/asset/util"
	"fmt"
)

func printOption() {
	fmt.Println("===== Room Management =====")
	fmt.Println("1. Add new Room")
	fmt.Println("2. List all Rooms")
	fmt.Println("3. Get detail of a Room")
	fmt.Println("4. Update a Room")
	fmt.Println("5. Delete a Room")
	fmt.Println("Type 'back' to return to previous menu")
	fmt.Println("===========================")
}

func RoomHandler(facade *controller.SpaceManagementControllerFacade) {
	inputBuffer := ""

	for inputBuffer != "back" {
		util.ClearScreen()
		util.PrintBanner()
		printOption()
		inputBuffer = util.GetCommandInput()

		switch inputBuffer {
		case "1":
			fmt.Println("Add new Room")
		case "2":
			fmt.Println("List all Room")
		case "3":
			fmt.Println("Get detail of a Room")
		case "4":
			fmt.Println("Update a Room")
		case "5":
			fmt.Println("Delete a Room")
		}

		util.ClearScreen()
	}

	util.ClearScreen()
}
