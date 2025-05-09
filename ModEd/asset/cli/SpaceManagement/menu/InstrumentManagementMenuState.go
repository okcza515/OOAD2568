// MEP-1013
package menu

import (
	spaceManagementHandler "ModEd/asset/cli/spacemanagement/handler"
	controller "ModEd/asset/controller"
	model "ModEd/asset/model"
	"ModEd/asset/util"
	"ModEd/core/cli"
	"ModEd/core/handler"
	"fmt"

	"gorm.io/gorm"
)

type InstrumentManagementMenuState struct {
	manager        *cli.CLIMenuStateManager //changing the state(handler), encapsulate menu transition
	handlerContext *handler.HandlerContext  //selecting the func(strategy), hold command handlers
}

func (menu *InstrumentManagementMenuState) Render() {
	fmt.Println("========== Instrument Management ==========")
	fmt.Println("Please select your action")
	fmt.Println("1. See all Instrument Managements")
	fmt.Println("2. Get the Instrument Management by its ID")
	fmt.Println("3. Get the Instrument Managements by the RoomID")
	fmt.Println("4. Create new Instrument Management")
	fmt.Println("5. Update the Instrument Management")
	fmt.Println("6. Delete the Instrument Management")
	fmt.Println("7. Seed Instrument Managements data")
	fmt.Println("Type 'back' to return to previous menu")
	fmt.Println("========================================")
}

func (menu *InstrumentManagementMenuState) HandleUserInput(input string) error {
	err := menu.handlerContext.HandleInput(input)
	if err != nil {
		fmt.Println("Error handling user input", err)
	}
	if input == "back" {

		util.PressEnterToContinue()

	}
	return err
}

func NewInstrumentMenuState(db *gorm.DB, manager *cli.CLIMenuStateManager, spaceManagementMenu *SpaceManagementState) *InstrumentManagementMenuState {
	if db == nil {
		fmt.Println("Error: Database connection is nil")
		return &InstrumentManagementMenuState{
			manager:        manager,
			handlerContext: handler.NewHandlerContext(),
		}
	}

	controllerManager := controller.GetSpaceManagementInstance(db)
	if controllerManager == nil {
		fmt.Println("Error: Space Management Controller Manager is nil")
		return &InstrumentManagementMenuState{
			manager:        manager,
			handlerContext: handler.NewHandlerContext(),
		}
	}

	controllerInstance := controllerManager.InstrumentManagement
	if controllerInstance == nil {
		fmt.Println("Error: Instrument Management controller is nil")
		return &InstrumentManagementMenuState{
			manager:        manager,
			handlerContext: handler.NewHandlerContext(),
		}
	}

	handlerContext := handler.NewHandlerContext()

	//handler strategy select each handler function
	//Standard Handlers
	listHandler := handler.NewListHandlerStrategy[model.InstrumentManagement](controllerInstance)
	getHandler := handler.NewRetrieveByIDHandlerStrategy[model.InstrumentManagement](controllerInstance)
	deleteHandler := handler.NewDeleteHandlerStrategy[model.InstrumentManagement](controllerInstance)
	backHandler := handler.NewChangeMenuHandlerStrategy(manager, spaceManagementMenu)
	insertManyhandler := handler.NewInsertHandlerStrategy[model.InstrumentManagement](controllerInstance)

	//Custom Handlers
	updateHandler := spaceManagementHandler.NewUpdateInstrumentManagementStrategy(controllerInstance)
	getByRoomIDhandler := spaceManagementHandler.NewGetInstrumentManagementByRoomIdStrategy(controllerInstance)
	insertHandler := spaceManagementHandler.NewInsertInstrumentManagementStrategy(controllerInstance)

	handlerContext.AddHandler("1", "Get all Instrument Managements", listHandler)
	handlerContext.AddHandler("2", "Get Instrument Management by ID", getHandler)
	handlerContext.AddHandler("3", "Get Instrument Management by RoomID", getByRoomIDhandler)
	handlerContext.AddHandler("4", "Create an Instrument Management", insertHandler)
	handlerContext.AddHandler("5", "Update an Instrument Management", updateHandler)
	handlerContext.AddHandler("6", "Delete an Instrument Management", deleteHandler)
	handlerContext.AddHandler("7", "Seed Instrument Managements data", insertManyhandler)
	handlerContext.AddHandler("back", "Back to main menu", backHandler)

	return &InstrumentManagementMenuState{
		manager:        manager,
		handlerContext: handlerContext,
	}

}

// func InstrumentManagementHandler(facade *controller.SpaceManagementControllerFacade) {
// 	inputBuffer := ""
// 	for inputBuffer != "back" {
// 		util.ClearScreen()
// 		util.PrintBanner()
// 		printInstrumentManagementOptions()
// 		inputBuffer = util.GetCommandInput()
// 		switch inputBuffer {
// 		case "1":
// 			handleGetAllInstrumentManagement(facade)
// 		case "2":
// 			handleGetInstrumentManagementByID(facade)
// 		case "3":
// 			handleGetInstrumentManagementByRoomID(facade)
// 		case "4":
// 			handleCreateInstrumentManagement(facade)
// 		case "5":
// 			handleUpdateInstrumentManagement(facade)
// 		case "6":
// 			handleDeleteInstrumentManagement(facade)
// 		}
// 	}
// }

// func handleGetAllInstrumentManagement(facade *controller.SpaceManagementControllerFacade) {
// 	fmt.Printf("Getting all Instrument management records...\n")

// 	// Get data through the facade
// 	instruments, err := facade.InstrumentManagement.GetAll()
// 	if err != nil {
// 		fmt.Printf("Error retrieving instrument data: %v\n", err)
// 		return
// 	}

// 	// Display results
// 	fmt.Println("\n=== Instrument Management List ===")
// 	for _, item := range *instruments {
// 		fmt.Printf("ID: %d | Name: %s | Room: %d |\n",
// 			item.ID, item.InstrumentLabel, item.RoomID)
// 	}

// 	fmt.Println("\nPress Enter to continue...")
// 	util.GetCommandInput()
// }

// func handleGetInstrumentManagementByID(facade *controller.SpaceManagementControllerFacade) {
// 	fmt.Println("Please insert your Instrument management ID")
// 	ID := util.GetCommandInput()

// 	idNum, err := strconv.ParseUint(ID, 10, 32)
// 	if err != nil {
// 		fmt.Printf("Error: Invalid ID format - %v\n", err)
// 		return
// 	}

// 	instrument, err := facade.InstrumentManagement.GetById(uint(idNum))
// 	if err != nil {
// 		fmt.Printf("Error retrieving instrument data: %v\n", err)
// 		return
// 	}

// 	fmt.Println("\n=== Instrument Details ===")
// 	// Dereference the pointer to access the struct fields
// 	fmt.Printf("ID: %d | Name: %s | Room: %d\n",
// 		instrument.ID, instrument.InstrumentLabel, instrument.RoomID)

// 	fmt.Println("\nPress Enter to continue...")
// 	util.GetCommandInput()
// }

// func handleGetInstrumentManagementByRoomID(facade *controller.SpaceManagementControllerFacade) {
// 	fmt.Println("Please insert Room ID to get Instrument management:")
// 	roomID := util.GetCommandInput()

// 	roomIDNum, err := strconv.ParseUint(roomID, 10, 32)
// 	if err != nil {
// 		fmt.Printf("Error: Invalid Room ID format - %v\n", err)
// 		return
// 	}

// 	instruments, err := facade.InstrumentManagement.GetByRoomId(uint(roomIDNum))
// 	if err != nil {
// 		fmt.Printf("Error retrieving instrument data for room %s: %v\n", roomID, err)
// 		return
// 	}

// 	fmt.Printf("\n=== Instruments in Room %s ===\n", roomID)
// 	if len(*instruments) == 0 {
// 		fmt.Printf("No instruments found in room %s\n", roomID)
// 		return
// 	}

// 	for _, item := range *instruments {
// 		fmt.Printf("ID: %d | Name: %s | Room: %d\n",
// 			item.ID, item.InstrumentLabel, item.RoomID)
// 	}

// 	fmt.Println("\nPress Enter to continue...")
// 	util.GetCommandInput()
// }

// func handleCreateInstrumentManagement(facade *controller.SpaceManagementControllerFacade) {
// 	fmt.Println("Create a new Instrument management")

// 	fmt.Print("Enter Room ID: ")
// 	roomID := util.GetCommandInput()
// 	roomIDNum, err := strconv.ParseUint(roomID, 10, 32)
// 	if err != nil {
// 		fmt.Printf("Error: Invalid Room ID format - %v\n", err)
// 		return
// 	}

// 	fmt.Print("Enter Instrument Name: ")
// 	name := util.GetCommandInput()

// 	instrument := &model.InstrumentManagement{
// 		RoomID:          uint(roomIDNum),
// 		InstrumentLabel: name,
// 	}

// 	err = facade.InstrumentManagement.Create(instrument)
// 	if err != nil {
// 		fmt.Printf("Error creating instrument: %v\n", err)
// 		return
// 	}

// 	fmt.Println("\nSuccessfully created new instrument!")
// 	fmt.Println("Press Enter to continue...")
// 	util.GetCommandInput()
// }

// func handleUpdateInstrumentManagement(facade *controller.SpaceManagementControllerFacade) {
// 	fmt.Println("Update the Instrument management")

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

// 	fmt.Print("Enter new Instrument Name: ")
// 	name := util.GetCommandInput()

// 	instrument := &model.InstrumentManagement{
// 		RoomID:          uint(roomIDNum),
// 		InstrumentLabel: name,
// 	}

// 	err = facade.InstrumentManagement.Update(uint(idNum), instrument)
// 	if err != nil {
// 		fmt.Printf("Error updating instrument: %v\n", err)
// 		return
// 	}

// 	fmt.Printf("\nSuccessfully updated instrument with ID %s!\n", ID)
// 	fmt.Println("Press Enter to continue...")
// 	util.GetCommandInput()
// }

// func handleDeleteInstrumentManagement(facade *controller.SpaceManagementControllerFacade) {
// 	fmt.Println("Delete the Instrument management")

// 	fmt.Print("Enter ID to delete: ")
// 	ID := util.GetCommandInput()
// 	idNum, err := strconv.ParseUint(ID, 10, 32)
// 	if err != nil {
// 		fmt.Printf("Error: Invalid ID format - %v\n", err)
// 		return
// 	}

// 	fmt.Printf("Are you sure you want to delete instrument with ID %s? (y/n): ", ID)
// 	confirm := util.GetCommandInput()
// 	if confirm != "y" {
// 		fmt.Println("Deletion cancelled")
// 		return
// 	}

// 	err = facade.InstrumentManagement.Delete(uint(idNum))
// 	if err != nil {
// 		fmt.Printf("Error deleting instrument: %v\n", err)
// 		return
// 	}

// 	fmt.Printf("\nSuccessfully deleted instrument with ID %s!\n", ID)
// 	fmt.Println("Press Enter to continue...")
// 	util.GetCommandInput()
// }
