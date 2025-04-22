// MEP-1013
package handler

import (
	controller "ModEd/asset/controller"
	model "ModEd/asset/model/spacemanagement"
	"ModEd/asset/util"
	"fmt"
	"strconv"
)

func printAssetManagementOption() {
	fmt.Println("========== Asset Management ==========")
	fmt.Println("Please select your asset management type")
	fmt.Println("1. Instrument Management")
	fmt.Println("2. Supply Management")
	fmt.Println("Type 'back' to return to previous menu")
	fmt.Println("========================================")
}

func AssetManagementHandler(facade *controller.SpaceManagementControllerFacade) {
	inputBuffer := ""
	assetType := ""

	for inputBuffer != "back" {
		util.ClearScreen()
		util.PrintBanner()
		printAssetManagementOption()
		inputBuffer = util.GetCommandInput()

		switch inputBuffer {
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

			switch choice {
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

			switch choice {
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

func handleGetAll(facade *controller.SpaceManagementControllerFacade, assetType string) {
	fmt.Printf("Getting all %s management records...\n", assetType)

	var assetTypeEnum controller.AssetType
	switch assetType {
	case "Instrument":
		assetTypeEnum = controller.Instrument
	case "Supply":
		assetTypeEnum = controller.Supply
	default:
		fmt.Println("Error: Invalid asset type")
		return
	}

	// Get data through the facade
	assets, err := facade.AssetManagement.GetAllAsset(assetTypeEnum)
	if err != nil {
		fmt.Printf("Error retrieving %s data: %v\n", assetType, err)
		return
	}

	// Display results
	fmt.Printf("\n=== %s Management List ===\n", assetType)
	switch items := assets.(type) {
	case []model.InstrumentManagement:
		for _, item := range items {
			fmt.Printf("ID: %d | Name: %s | Room: %d |\n",
				item.ID, item.InstrumentLabel, item.RoomID)
		}
	case []model.SupplyManagement:
		for _, item := range items {
			fmt.Printf("ID: %d | Name: %s | Room: %d | Quantity: %d\n",
				item.ID, item.SupplyLabel, item.RoomID, item.Quantity)
		}
	default:
		fmt.Println("Error: Unknown data type returned")
	}

	fmt.Println("\nPress Enter to continue...")
	util.GetCommandInput()
}

func handleGetByID(facade *controller.SpaceManagementControllerFacade, assetType string) {
	fmt.Printf("Please insert your %s management ID\n", assetType)
	ID := util.GetCommandInput()

	// Convert ID string to uint
	idNum, err := strconv.ParseUint(ID, 10, 32)
	if err != nil {
		fmt.Printf("Error: Invalid ID format - %v\n", err)
		return
	}

	fmt.Printf("Getting the %s management at the ID %s... \n", assetType, ID)

	var assetTypeEnum controller.AssetType
	switch assetType {
	case "Instrument":
		assetTypeEnum = controller.Instrument
	case "Supply":
		assetTypeEnum = controller.Supply
	default:
		fmt.Println("Error: Invalid asset type")
		return
	}

	asset, err := facade.AssetManagement.GetAssetById(assetTypeEnum, uint(idNum))
	if err != nil {
		fmt.Printf("Error retrieving %s data: %v\n", assetType, err)
		return
	}

	fmt.Printf("\n=== %s Details ===\n", assetType)
	switch item := asset.(type) {
	case *model.InstrumentManagement:
		fmt.Printf("ID: %d | Name: %s | Room: %d\n",
			item.ID, item.InstrumentLabel, item.RoomID)
	case *model.SupplyManagement:
		fmt.Printf("ID: %d | Name: %s | Room: %d | Quantity: %d\n",
			item.ID, item.SupplyLabel, item.RoomID, item.Quantity)
	default:
		fmt.Println("Error: Unknown data type returned")
	}

	fmt.Println("\nPress Enter to continue...")
	util.GetCommandInput()
}

func handleGetByRoomID(facade *controller.SpaceManagementControllerFacade, assetType string) {
	fmt.Printf("Please insert Room ID to get %s management: \n", assetType)
	roomID := util.GetCommandInput()

	// Convert roomID string to uint
	roomIDNum, err := strconv.ParseUint(roomID, 10, 32)
	if err != nil {
		fmt.Printf("Error: Invalid Room ID format - %v\n", err)
		return
	}

	var assetTypeEnum controller.AssetType
	switch assetType {
	case "Instrument":
		assetTypeEnum = controller.Instrument
	case "Supply":
		assetTypeEnum = controller.Supply
	default:
		fmt.Println("Error: Invalid asset type")
		return
	}

	assets, err := facade.AssetManagement.GetAssetByRoomId(assetTypeEnum, uint(roomIDNum))
	if err != nil {
		fmt.Printf("Error retrieving %s data for room %s: %v\n", assetType, roomID, err)
		return
	}

	fmt.Printf("\n=== %s in Room %s ===\n", assetType, roomID)
	switch items := assets.(type) {
	case []model.InstrumentManagement:
		if len(items) == 0 {
			fmt.Printf("No %s found in room %s\n", assetType, roomID)
			return
		}
		for _, item := range items {
			fmt.Printf("ID: %d | Name: %s | Room: %d\n",
				item.ID, item.InstrumentLabel, item.RoomID)
		}
	case []model.SupplyManagement:
		if len(items) == 0 {
			fmt.Printf("No %s found in room %s\n", assetType, roomID)
			return
		}
		for _, item := range items {
			fmt.Printf("ID: %d | Name: %s | Room: %d | Quantity: %d\n",
				item.ID, item.SupplyLabel, item.RoomID, item.Quantity)
		}
	default:
		fmt.Println("Error: Unknown data type returned")
	}

	fmt.Println("\nPress Enter to continue...")
	util.GetCommandInput()
}

func handleCreate(facade *controller.SpaceManagementControllerFacade, assetType string) {
	fmt.Printf("Create a new %s management\n", assetType)

	// Get room ID
	fmt.Print("Enter Room ID: ")
	roomID := util.GetCommandInput()
	roomIDNum, err := strconv.ParseUint(roomID, 10, 32)
	if err != nil {
		fmt.Printf("Error: Invalid Room ID format - %v\n", err)
		return
	}

	var assetTypeEnum controller.AssetType
	var payload interface{}

	switch assetType {
	case "Instrument":
		assetTypeEnum = controller.Instrument
		fmt.Print("Enter Instrument Name: ")
		name := util.GetCommandInput()
		fmt.Print("Enter Description: ")

		payload = &model.InstrumentManagement{
			RoomID:          uint(roomIDNum),
			InstrumentLabel: name,
		}

	case "Supply":
		assetTypeEnum = controller.Supply
		fmt.Print("Enter Supply Name: ")
		name := util.GetCommandInput()
		fmt.Print("Enter Quantity: ")
		qtyStr := util.GetCommandInput()
		qty, err := strconv.ParseInt(qtyStr, 10, 32)
		if err != nil {
			fmt.Printf("Error: Invalid quantity format - %v\n", err)
			return
		}
		fmt.Print("Enter Description: ")

		payload = &model.SupplyManagement{
			RoomID:      uint(roomIDNum),
			SupplyLabel: name,
			Quantity:    int(qty),
		}

	default:
		fmt.Println("Error: Invalid asset type")
		return
	}

	err = facade.AssetManagement.CreateAsset(assetTypeEnum, payload)
	if err != nil {
		fmt.Printf("Error creating %s: %v\n", assetType, err)
		return
	}

	fmt.Printf("\nSuccessfully created new %s!\n", assetType)
	fmt.Println("Press Enter to continue...")
	util.GetCommandInput()
}

func handleUpdate(facade *controller.SpaceManagementControllerFacade, assetType string) {
	fmt.Printf("Update the %s management\n", assetType)

	// Get ID of asset to update
	fmt.Print("Enter ID to update: ")
	ID := util.GetCommandInput()
	idNum, err := strconv.ParseUint(ID, 10, 32)
	if err != nil {
		fmt.Printf("Error: Invalid ID format - %v\n", err)
		return
	}

	// Get room ID
	fmt.Print("Enter new Room ID: ")
	roomID := util.GetCommandInput()
	roomIDNum, err := strconv.ParseUint(roomID, 10, 32)
	if err != nil {
		fmt.Printf("Error: Invalid Room ID format - %v\n", err)
		return
	}

	var assetTypeEnum controller.AssetType
	var payload interface{}

	switch assetType {
	case "Instrument":
		assetTypeEnum = controller.Instrument
		fmt.Print("Enter new Instrument Name: ")
		name := util.GetCommandInput()
		fmt.Print("Enter new Description: ")

		payload = &model.InstrumentManagement{
			// ID:             uint(idNum),
			RoomID:          uint(roomIDNum),
			InstrumentLabel: name,
		}

	case "Supply":
		assetTypeEnum = controller.Supply
		fmt.Print("Enter new Supply Name: ")
		name := util.GetCommandInput()
		fmt.Print("Enter new Quantity: ")
		qtyStr := util.GetCommandInput()
		qty, err := strconv.ParseInt(qtyStr, 10, 32)
		if err != nil {
			fmt.Printf("Error: Invalid quantity format - %v\n", err)
			return
		}
		fmt.Print("Enter new Description: ")

		payload = &model.SupplyManagement{
			// ID:          uint(idNum),
			RoomID:      uint(roomIDNum),
			SupplyLabel: name,
			Quantity:    int(qty),
		}

	default:
		fmt.Println("Error: Invalid asset type")
		return
	}

	err = facade.AssetManagement.UpdateAsset(assetTypeEnum, uint(idNum), payload)
	if err != nil {
		fmt.Printf("Error updating %s: %v\n", assetType, err)
		return
	}

	fmt.Printf("\nSuccessfully updated %s with ID %s!\n", assetType, ID)
	fmt.Println("Press Enter to continue...")
	util.GetCommandInput()
}

func handleDelete(facade *controller.SpaceManagementControllerFacade, assetType string) {
	fmt.Printf("Delete the %s management\n", assetType)

	// Get ID of asset to delete
	fmt.Print("Enter ID to delete: ")
	ID := util.GetCommandInput()
	idNum, err := strconv.ParseUint(ID, 10, 32)
	if err != nil {
		fmt.Printf("Error: Invalid ID format - %v\n", err)
		return
	}

	var assetTypeEnum controller.AssetType
	switch assetType {
	case "Instrument":
		assetTypeEnum = controller.Instrument
	case "Supply":
		assetTypeEnum = controller.Supply
	default:
		fmt.Println("Error: Invalid asset type")
		return
	}

	// Confirm deletion
	fmt.Printf("Are you sure you want to delete %s with ID %s? (y/n): ", assetType, ID)
	confirm := util.GetCommandInput()
	if confirm != "y" {
		fmt.Println("Deletion cancelled")
		return
	}

	// Delete the asset
	err = facade.AssetManagement.DeleteAsset(assetTypeEnum, uint(idNum))
	if err != nil {
		fmt.Printf("Error deleting %s: %v\n", assetType, err)
		return
	}

	fmt.Printf("\nSuccessfully deleted %s with ID %s!\n", assetType, ID)
	fmt.Println("Press Enter to continue...")
	util.GetCommandInput()
}
