// MEP-1013
package menu

import (
	controller "ModEd/asset/controller"
	model "ModEd/asset/model"
	"ModEd/asset/util"
	"ModEd/core/cli"
	"ModEd/core/handler"
	"fmt"
	"gorm.io/gorm"
	spaceManagementHandler "ModEd/asset/handler"
)

type SupplyManagementMenuState struct{
	manager			*cli.CLIMenuStateManager
	handlerContext	*handler.HandlerContext
}

func (menu *SupplyManagementMenuState) Render(){
	fmt.Println("========== Instrument Management ==========")
	fmt.Println("Please select your action")
	fmt.Println("1. See all Supply Managements")
	fmt.Println("2. Get the Supply Management by its ID")
	fmt.Println("3. Get the Supply Managements by the RoomID")
	fmt.Println("4. Create new Supply Management")
	fmt.Println("5. Update the Supply Management")
	fmt.Println("6. Delete the Supply Management")
	fmt.Println("Type 'back' to return to previous menu")
	fmt.Println("========================================")
}

func (menu *SupplyManagementMenuState) HandleUserInput(input string) error{
	err := menu.handlerContext.HandleInput(input)
	if err != nil {
		fmt.Println("Error handling user input", err)
	}
	if input == "back" {

		util.PressEnterToContinue()

	}
	return err
}

func NewSupplyMenuState(db *gorm.DB, manager *cli.CLIMenuStateManager, spaceManagementMenu *SpaceManagementState) *SupplyManagementMenuState {
	if db == nil {
		fmt.Println("Error: Database connection is nil")
		return &SupplyManagementMenuState{
			manager:        manager,
			handlerContext: handler.NewHandlerContext(),
		}
	}

	controllerManager := controller.GetSpaceManagementInstance(db)
	if controllerManager == nil {
		fmt.Println("Error: Space Management Controller Manager is nil")
		return &SupplyManagementMenuState{
			manager:        manager,
			handlerContext: handler.NewHandlerContext(),
		}
	}

  controllerInstance := controllerManager.SupplyManagement
  if controllerInstance == nil{
    fmt.Println("Error: Supply Management controller is nil")
    return &SupplyManagementMenuState{
      manager: manager,
      handlerContext: handler.NewHandlerContext(),
    }
  }

  	handlerContext := handler.NewHandlerContext()

  	//Standard Handlers
	listHandler := handler.NewListHandlerStrategy[model.SupplyManagement](controllerInstance)
	getHandler := handler.NewRetrieveByIDHandlerStrategy[model.SupplyManagement](controllerInstance)
	deleteHandler := handler.NewDeleteHandlerStrategy[model.SupplyManagement](controllerInstance)
	backHandler := handler.NewChangeMenuHandlerStrategy(manager, spaceManagementMenu)

  	//Custom Handlers
	updateHandler := spaceManagementHandler.NewUpdateSupplyManagementStrategy(controllerInstance)
  	getByRoomIDhandler := spaceManagementHandler.NewGetSupplyManagementByRoomIdStrategy(controllerInstance)
  	insertHandler := spaceManagementHandler.NewInsertSupplyManagementStrategy(controllerInstance)

  	handlerContext.AddHandler("1", "Get all Supply Managements", listHandler)
	handlerContext.AddHandler("2", "Get Supply Management by ID", getHandler)
	handlerContext.AddHandler("3", "Get Supply Management by RoomID", getByRoomIDhandler)
	handlerContext.AddHandler("4", "Create an Supply Management", insertHandler)
	handlerContext.AddHandler("5", "Update an Supply Management", updateHandler)
	handlerContext.AddHandler("6", "Delete an Supply Management", deleteHandler)
	handlerContext.AddHandler("back", "Back to main menu", backHandler)


  return &SupplyManagementMenuState{
    manager: manager,
    handlerContext: handlerContext,
  }

}

// func printSupplyManagementOptions() {
// 	fmt.Println("========== Supply Management ==========")
// 	fmt.Println("Please select your action")
// 	fmt.Println("1. See all Supply Management")
// 	fmt.Println("2. Get the Supply Management by its ID")
// 	fmt.Println("3. Get the Supply Management by the RoomID")
// 	fmt.Println("4. Create new Supply Management")
// 	fmt.Println("5. Update the Supply Management")
// 	fmt.Println("6. Delete the Supply Management")
// 	fmt.Println("Type 'back' to return to previous menu")
// 	fmt.Println("========================================")
// }

// func SupplyManagementHandler(facade *controller.SpaceManagementControllerFacade) {
// 	inputBuffer := ""
// 	for inputBuffer != "back" {
// 		util.ClearScreen()
// 		util.PrintBanner()
// 		printSupplyManagementOptions()
// 		inputBuffer = util.GetCommandInput()
// 		switch inputBuffer {
// 		case "1":
// 			handleGetAllSupplyManagement(facade)
// 		case "2":
// 			handleGetSupplyManagementByID(facade)
// 		case "3":
// 			handleGetSupplyManagementByRoomID(facade)
// 		case "4":
// 			handleCreateSupplyManagement(facade)
// 		case "5":
// 			handleUpdateSupplyManagement(facade)
// 		case "6":
// 			handleDeleteSupplyManagement(facade)
// 		}
// 	}
// }

// func handleGetAllSupplyManagement(facade *controller.SpaceManagementControllerFacade) {
// 	fmt.Printf("Getting all Supply management records...\n")

// 	supplies, err := facade.SupplyManagement.GetAll()
// 	if err != nil {
// 		fmt.Printf("Error retrieving supply data: %v\n", err)
// 		return
// 	}

// 	fmt.Println("\n=== Supply Management List ===")
// 	for _, item := range *supplies {
// 		fmt.Printf("ID: %d | Name: %s | Room: %d | Quantity: %d |\n",
// 			item.ID, item.SupplyLabel, item.RoomID, item.Quantity)
// 	}

// 	fmt.Println("\nPress Enter to continue...")
// 	util.GetCommandInput()
// }

// func handleGetSupplyManagementByID(facade *controller.SpaceManagementControllerFacade) {
// 	fmt.Println("Please insert your Supply management ID")
// 	ID := util.GetCommandInput()

// 	idNum, err := strconv.ParseUint(ID, 10, 32)
// 	if err != nil {
// 		fmt.Printf("Error: Invalid ID format - %v\n", err)
// 		return
// 	}

// 	supply, err := facade.SupplyManagement.GetById(uint(idNum))
// 	if err != nil {
// 		fmt.Printf("Error retrieving supply data: %v\n", err)
// 		return
// 	}

// 	fmt.Println("\n=== Supply Details ===")
// 	fmt.Printf("ID: %d | Name: %s | Room: %d | Quantity: %d\n",
// 		supply.ID, supply.SupplyLabel, supply.RoomID, supply.Quantity)

// 	fmt.Println("\nPress Enter to continue...")
// 	util.GetCommandInput()
// }

// func handleGetSupplyManagementByRoomID(facade *controller.SpaceManagementControllerFacade) {
// 	fmt.Println("Please insert Room ID to get Supply management:")
// 	roomID := util.GetCommandInput()

// 	roomIDNum, err := strconv.ParseUint(roomID, 10, 32)
// 	if err != nil {
// 		fmt.Printf("Error: Invalid Room ID format - %v\n", err)
// 		return
// 	}

// 	supplies, err := facade.SupplyManagement.GetByRoomId(uint(roomIDNum))
// 	if err != nil {
// 		fmt.Printf("Error retrieving supply data for room %s: %v\n", roomID, err)
// 		return
// 	}

// 	fmt.Printf("\n=== Supplies in Room %s ===\n", roomID)
// 	if len(*supplies) == 0 {
// 		fmt.Printf("No supplies found in room %s\n", roomID)
// 		return
// 	}

// 	for _, item := range *supplies {
// 		fmt.Printf("ID: %d | Name: %s | Room: %d | Quantity: %d\n",
// 			item.ID, item.SupplyLabel, item.RoomID, item.Quantity)
// 	}

// 	fmt.Println("\nPress Enter to continue...")
// 	util.GetCommandInput()
// }

// func handleCreateSupplyManagement(facade *controller.SpaceManagementControllerFacade) {
// 	fmt.Println("Create a new Supply management")

// 	fmt.Print("Enter Room ID: ")
// 	roomID := util.GetCommandInput()
// 	roomIDNum, err := strconv.ParseUint(roomID, 10, 32)
// 	if err != nil {
// 		fmt.Printf("Error: Invalid Room ID format - %v\n", err)
// 		return
// 	}

// 	fmt.Print("Enter Supply Name: ")
// 	name := util.GetCommandInput()
// 	fmt.Print("Enter Quantity: ")
// 	quantityStr := util.GetCommandInput()

// 	quantity, err := strconv.ParseUint(quantityStr, 10, 32)
// 	if err != nil {
// 		fmt.Printf("Error: Invalid quantity format - %v\n", err)
// 		return
// 	}

// 	supply := &model.SupplyManagement{
// 		RoomID:      uint(roomIDNum),
// 		SupplyLabel: name,
// 		Quantity:    int(quantity),
// 	}

// 	err = facade.SupplyManagement.Create(supply)
// 	if err != nil {
// 		fmt.Printf("Error creating supply: %v\n", err)
// 		return
// 	}

// 	fmt.Println("\nSuccessfully created new supply!")
// 	fmt.Println("Press Enter to continue...")
// 	util.GetCommandInput()
// }

// func handleUpdateSupplyManagement(facade *controller.SpaceManagementControllerFacade) {
// 	fmt.Println("Update the Supply management")

// 	fmt.Print("Enter ID to update: ")
// 	ID := util.GetCommandInput()
// 	idNum, err := strconv.ParseUint(ID, 10, 32)
// 	if err != nil {
// 		fmt.Printf("Error: Invalid ID format - %v\n", err)
// 		return
// 	}

// 	fmt.Print("Enter new Room ID: ")
// 	roomID := util.GetCommandInput()
// 	roomIDNum, err := strconv.ParseUint(roomID, 10, 32)
// 	if err != nil {
// 		fmt.Printf("Error: Invalid Room ID format - %v\n", err)
// 		return
// 	}

// 	fmt.Print("Enter new Supply Name: ")
// 	name := util.GetCommandInput()
// 	fmt.Print("Enter new Quantity: ")
// 	quantityStr := util.GetCommandInput()

// 	quantity, err := strconv.ParseUint(quantityStr, 10, 32)
// 	if err != nil {
// 		fmt.Printf("Error: Invalid quantity format - %v\n", err)
// 		return
// 	}

// 	supply := &model.SupplyManagement{
// 		RoomID:      uint(roomIDNum),
// 		SupplyLabel: name,
// 		Quantity:    int(quantity),
// 	}

// 	err = facade.SupplyManagement.Update(uint(idNum), supply)
// 	if err != nil {
// 		fmt.Printf("Error updating supply: %v\n", err)
// 		return
// 	}

// 	fmt.Printf("\nSuccessfully updated supply with ID %s!\n", ID)
// 	fmt.Println("Press Enter to continue...")
// 	util.GetCommandInput()
// }

// func handleDeleteSupplyManagement(facade *controller.SpaceManagementControllerFacade) {
// 	fmt.Println("Delete the Supply management")

// 	fmt.Print("Enter ID to delete: ")
// 	ID := util.GetCommandInput()
// 	idNum, err := strconv.ParseUint(ID, 10, 32)
// 	if err != nil {
// 		fmt.Printf("Error: Invalid ID format - %v\n", err)
// 		return
// 	}

// 	fmt.Printf("Are you sure you want to delete supply with ID %s? (y/n): ", ID)
// 	confirm := util.GetCommandInput()
// 	if confirm != "y" {
// 		fmt.Println("Deletion cancelled")
// 		return
// 	}

// 	err = facade.SupplyManagement.Delete(uint(idNum))
// 	if err != nil {
// 		fmt.Printf("Error deleting instrument: %v\n", err)
// 		return
// 	}

// 	fmt.Printf("\nSuccessfully deleted supply with ID %s!\n", ID)
// 	fmt.Println("Press Enter to continue...")
// 	util.GetCommandInput()
// }
