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
	manager        *cli.CLIMenuStateManager
	handlerContext *handler.HandlerContext
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

	listHandler := handler.NewListHandlerStrategy[model.InstrumentManagement](controllerInstance)
	getHandler := handler.NewRetrieveByIDHandlerStrategy[model.InstrumentManagement](controllerInstance)
	deleteHandler := handler.NewDeleteHandlerStrategy[model.InstrumentManagement](controllerInstance)
	backHandler := handler.NewChangeMenuHandlerStrategy(manager, spaceManagementMenu)
	insertManyhandler := handler.NewInsertHandlerStrategy[model.InstrumentManagement](controllerInstance)

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
