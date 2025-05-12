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

type SupplyManagementMenuState struct {
	manager        *cli.CLIMenuStateManager
	handlerContext *handler.HandlerContext
}

func (menu *SupplyManagementMenuState) Render() {
	fmt.Println("========== Instrument Management ==========")
	fmt.Println("Please select your action")
	fmt.Println("1. See all Supply Managements")
	fmt.Println("2. Get the Supply Management by its ID")
	fmt.Println("3. Get the Supply Managements by the RoomID")
	fmt.Println("4. Create new Supply Management")
	fmt.Println("5. Update the Supply Management")
	fmt.Println("6. Delete the Supply Management")
	fmt.Println("7. Seed Supply Managements data")
	fmt.Println("Type 'back' to return to previous menu")
	fmt.Println("========================================")
}

func (menu *SupplyManagementMenuState) HandleUserInput(input string) error {
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
	if controllerInstance == nil {
		fmt.Println("Error: Supply Management controller is nil")
		return &SupplyManagementMenuState{
			manager:        manager,
			handlerContext: handler.NewHandlerContext(),
		}
	}

	handlerContext := handler.NewHandlerContext()

	//Standard Handlers
	listHandler := handler.NewListHandlerStrategy[model.SupplyManagement](controllerInstance)
	getHandler := handler.NewRetrieveByIDHandlerStrategy[model.SupplyManagement](controllerInstance)
	deleteHandler := handler.NewDeleteHandlerStrategy[model.SupplyManagement](controllerInstance)
	backHandler := handler.NewChangeMenuHandlerStrategy(manager, spaceManagementMenu)
	insertManyhandler := handler.NewInsertHandlerStrategy[model.SupplyManagement](controllerInstance)

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
	handlerContext.AddHandler("7", "Seed Supply Managements data", insertManyhandler)
	handlerContext.AddHandler("back", "Back to main menu", backHandler)

	return &SupplyManagementMenuState{
		manager:        manager,
		handlerContext: handlerContext,
	}

}
