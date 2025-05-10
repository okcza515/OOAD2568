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

type RoomMenuState struct {
	manager        *cli.CLIMenuStateManager
	handlerContext *handler.HandlerContext
}

func (menu *RoomMenuState) Render() {
	fmt.Println()
	fmt.Println("===== Room Management =====")
	fmt.Println("1. Add new Room")
	fmt.Println("2. List all Rooms")
	fmt.Println("3. Get detail of a Room")
	fmt.Println("4. Update a Room")
	fmt.Println("5. Delete a Room")
	fmt.Println("6. Delete all Rooms")
	// fmt.Println("7. Seed Rooms Data")
	fmt.Println("Type 'back' to return to previous menu")
	fmt.Println("===========================")
}
func (menu *RoomMenuState) HandleUserInput(input string) error {
	err := menu.handlerContext.HandleInput(input)
	if err != nil {
		fmt.Println("Error handling user input", err)
	}
	if input == "back" {

		util.PressEnterToContinue()

	}
	return err
}

func NewRoomMenuState(db *gorm.DB, manager *cli.CLIMenuStateManager, spaceManagementMenu *SpaceManagementState) *RoomMenuState {
	// Check if db is nil
	if db == nil {
		fmt.Println("Error: Database connection is nil")
		return &RoomMenuState{
			manager:        manager,
			handlerContext: handler.NewHandlerContext(),
		}
	}

	// Get controller instance and check if it's nil
	controllerFacade := controller.GetSpaceManagementInstance(db)
	if controllerFacade == nil {
		fmt.Println("Error: Space Management Controller Facade is nil")
		return &RoomMenuState{
			manager:        manager,
			handlerContext: handler.NewHandlerContext(),
		}
	}

	// Access Room controller safely
	controllerInstance := controllerFacade.Room
	if controllerInstance == nil {
		fmt.Println("Error: Room controller is nil")
		return &RoomMenuState{
			manager:        manager,
			handlerContext: handler.NewHandlerContext(),
		}
	}

	handlerContext := handler.NewHandlerContext()

	//Handler is not working yet!
	// insertHandler := handler.NewInsertHandlerStrategy[model.Room](controllerInstance)
	listHandler := handler.NewListHandlerStrategy[model.Room](controllerInstance)
	getHandler := handler.NewRetrieveByIDHandlerStrategy[model.Room](controllerInstance)
	// updateHandler := handler.NewUpdateHandlerStrategy[model.Room](controllerInstance)
	deleteHandler := handler.NewDeleteHandlerStrategy[model.Room](controllerInstance)
	backHandler := handler.NewChangeMenuHandlerStrategy(manager, spaceManagementMenu)

	//DIY
	insertHandler := spaceManagementHandler.NewAddRoomHandlerStrategy(controllerInstance)
	updateHandler := spaceManagementHandler.NewUpdateRoomHandlerStrategy(controllerInstance)
	deleteAllHandler := spaceManagementHandler.NewDeleteAllRoomStrategy(controllerInstance)

	handlerContext.AddHandler("1", "Add New Room", insertHandler)
	handlerContext.AddHandler("2", "List all Rooms", listHandler)
	handlerContext.AddHandler("3", "Get full detail of a Room", getHandler)
	handlerContext.AddHandler("4", "Update a Room", updateHandler)
	handlerContext.AddHandler("5", "Delete a Room", deleteHandler)
	handlerContext.AddHandler("6", "Delete all Rooms", deleteAllHandler)
	// handlerContext.AddHandler("7", "Seed Rooms Data", nil)
	handlerContext.AddHandler("back", "Back to main menu", backHandler)

	return &RoomMenuState{
		manager:        manager,
		handlerContext: handlerContext,
	}
}
