// MEP-1013
package handler

import (
	controller "ModEd/asset/controller/spacemanagement"
	"fmt"

	// model "ModEd/asset/model/spacemanagement"
	"ModEd/asset/util"
)

func printAssetManagementOption(){
	fmt.Println("========== Asset Management ==========")
	fmt.Println("Please select your asset management type")
	fmt.Println("1. Instrument Management")
	fmt.Println("2. Supply Management")
	fmt.Println("Type 'back' to return to previous menu")
	fmt.Println("========================================")
}

func AssetManagementHandler(facade *controller.SpaceManagementControllerFacade){
	inputBuffer := ""

	for inputBuffer != "back"{
		util.ClearScreen()
		util.PrintBanner()
		printAssetManagementOption()
		inputBuffer = util.GetCommandInput()

		switch inputBuffer{
		case "1":
			fmt.Println("You have selected Instrument Management")
			fmt.Println("Please select what do you want to do?")
			fmt.Println("1. See all Instrument Management")
			fmt.Println("2. Get the Instrument Management by its ID")
			fmt.Println("3. Get the Instrument Management by the RoomID")
			fmt.Println("4. Create new Instrument Management")
			fmt.Println("5. Update the Instrument Management")
			fmt.Println("6. Delete the Instrument Management")
			choice := util.GetCommandInput()

			switch choice{
			case "1":
				handleGetAll(facade)
			case "2":
				handleGetByID(facade)
			case "3":
				handleGetByRoomID(facade)
			case "4":
				handleCreate(facade)
			case "5":
				handleUpdate(facade)
			case "6":
				handleDelete(facade)
			}


		case "2":
			fmt.Println("Supply Management")
			fmt.Println("Please select what do you want to do?")
			fmt.Println("1. See all Supply Management")
			fmt.Println("2. Get the Supply Management by its ID")
			fmt.Println("3. Get the Supply Management by the RoomID")
			fmt.Println("4. Create new Supply Management")
			fmt.Println("5. Update the Supply Management")
			fmt.Println("6. Delete the Supply Management")
			choice := util.GetCommandInput()

			switch choice{
			case "1":
				handleGetAll(facade)
			case "2":
				handleGetByID(facade)
			case "3":
				handleGetByRoomID(facade)
			case "4":
				handleCreate(facade)
			case "5":
				handleUpdate(facade)
			case "6":
				handleDelete(facade)
			}
		}
	}
}

func handleGetAll(facade *controller.SpaceManagementControllerFacade){

}

func handleGetByID(facade *controller.SpaceManagementControllerFacade){

}

func handleGetByRoomID(facade *controller.SpaceManagementControllerFacade){

}

func handleCreate(facade *controller.SpaceManagementControllerFacade){

}
func handleUpdate(facade *controller.SpaceManagementControllerFacade){

}

func handleDelete(facade *controller.SpaceManagementControllerFacade){

}