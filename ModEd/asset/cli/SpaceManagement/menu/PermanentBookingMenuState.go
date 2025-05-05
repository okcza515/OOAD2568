//MEP-1013

package menu

import (
	controller "ModEd/asset/controller"
	spaceManagementHandler "ModEd/asset/handler"
	"ModEd/asset/model"
	"ModEd/asset/util"
	"ModEd/core/cli"
	"ModEd/core/handler"
	"fmt"

	"gorm.io/gorm"
)

type PermanentBookingState struct {
	manager        *cli.CLIMenuStateManager
	handlerContext *handler.HandlerContext
}

func (menu *PermanentBookingState) Render() {
	fmt.Println()
	fmt.Println("===== Permanent Booking Schedule =====")
	fmt.Println("1. Create New Schedule")
	fmt.Println("2. List All Schedule")
	fmt.Println("3. Get Schedule Details")
	fmt.Println("4. Update Schedule")
	fmt.Println("5. Delete a Schedule")
	fmt.Println("6. Delete All Schedules")
	//fmt.Println("7. Seed All Schedules")
	fmt.Println("Type 'back' to return to previous menu")
	fmt.Println("======================================")
}

func (menu *PermanentBookingState) HandleUserInput(input string) error {
	err := menu.handlerContext.HandleInput(input)
	if err != nil {
		fmt.Println("Error handling user input:", err)
	}
	if input == "back" {
		util.PressEnterToContinue()
	}
	return err
}

func NewPermanentScheduleState(db *gorm.DB, manager *cli.CLIMenuStateManager, spaceManagementMenu *SpaceManagementState) *PermanentBookingState {
	if db == nil {
		fmt.Println("Error: Database connection is nil")
		return &PermanentBookingState{
			manager:        manager,
			handlerContext: handler.NewHandlerContext(),
		}
	}

	controllerFacade := controller.GetSpaceManagementInstance(db)
	if controllerFacade == nil {
		fmt.Println("Error: Space Management Controller Facade is nil")
		return &PermanentBookingState{
			manager:        manager,
			handlerContext: handler.NewHandlerContext(),
		}
	}

	controllerInstance := controllerFacade.PermanentSchedule
	if controllerInstance == nil {
		fmt.Println("Error: Permanent Booking Controller is nil")
		return &PermanentBookingState{
			manager:        manager,
			handlerContext: handler.NewHandlerContext(),
		}
	}

	handlerContext := handler.NewHandlerContext()

	createHandler := spaceManagementHandler.NewCreatePermanentScheduleHandler(controllerInstance)
	listHandler := handler.NewListHandlerStrategy[model.PermanentSchedule](controllerInstance)
	// listHandler := spaceManagementHandler.NewListPermanentSchedulesHandler(controllerInstance)
	// getHandler := spaceManagementHandler.NewGetScheduleDetailsHandler(controllerInstance)
	getHandler := handler.NewRetrieveByIDHandlerStrategy[model.PermanentSchedule](controllerInstance)
	updateHandler := spaceManagementHandler.NewUpdateScheduleHandler(controllerInstance)
	deleteHandler := spaceManagementHandler.NewDeleteScheduleHandler(controllerInstance)
	//seedHandler := spaceManagementHandler.NewSeedAllSchedulesHandler(controllerInstance)
	deleteAllHandler := spaceManagementHandler.NewDeleteAllSchedulesHandler(controllerInstance)

	backHandler := handler.NewChangeMenuHandlerStrategy(manager, spaceManagementMenu)

	handlerContext.AddHandler("1", "Create New Schedule", createHandler)
	handlerContext.AddHandler("2", "List All Schedule", listHandler)
	handlerContext.AddHandler("3", "Get Schedule Information", getHandler)
	handlerContext.AddHandler("4", "Update Schedule", updateHandler)
	handlerContext.AddHandler("5", "Delete a Schedule", deleteHandler)
	handlerContext.AddHandler("6", "Delete All Schedules", deleteAllHandler)
	handlerContext.AddHandler("back", "Back to main menu", backHandler)

	return &PermanentBookingState{
		manager:        manager,
		handlerContext: handlerContext,
	}
}
