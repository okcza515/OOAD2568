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
	assetType := ""

	for inputBuffer != "back"{
		util.ClearScreen()
		util.PrintBanner()
		printAssetManagementOption()
		inputBuffer = util.GetCommandInput()

		switch inputBuffer{
		case "1":
			assetType = "Instrument"
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
				handleGetAll(facade, assetType)
			case "2":
				handleGetByID(facade, assetType)
			case "3":
				handleGetByRoomID(facade, assetType)
			case "4":
				handleCreate(facade, assetType)
			case "5":
				handleUpdate(facade, assetType)
			case "6":
				handleDelete(facade, assetType)
			}


		case "2":
			assetType = "Supply"
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
				handleGetAll(facade, assetType)
			case "2":
				handleGetByID(facade, assetType)
			case "3":
				handleGetByRoomID(facade, assetType)
			case "4":
				handleCreate(facade, assetType)
			case "5":
				handleUpdate(facade, assetType)
			case "6":
				handleDelete(facade, assetType)
			}
		}
	}
}

func handleGetAll(facade *controller.SpaceManagementControllerFacade, assetType string){
	fmt.Printf("Get all the %s management \n",assetType)
}

func handleGetByID(facade *controller.SpaceManagementControllerFacade, assetType string){
	fmt.Printf("Get the %s management by ID \n",assetType)
}

func handleGetByRoomID(facade *controller.SpaceManagementControllerFacade, assetType string){
	fmt.Printf("Get the %s management by ID \n",assetType)
}

func handleCreate(facade *controller.SpaceManagementControllerFacade, assetType string){
	fmt.Printf("Create a new %s management \n",assetType)
}

func handleUpdate(facade *controller.SpaceManagementControllerFacade, assetType string){
	fmt.Printf("Update the %s management \n",assetType)
}

func handleDelete(facade *controller.SpaceManagementControllerFacade, assetType string){
	fmt.Printf("Delete the %s management \n",assetType)
}