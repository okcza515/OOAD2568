// MEP-1013
package menu

import (
	controller "ModEd/asset/controller"
	spaceManagementHandler "ModEd/asset/handler"
	model "ModEd/asset/model"
	"ModEd/asset/util"
	"ModEd/core/cli"
	"ModEd/core/handler"
	"fmt"
	"strings"

	"gorm.io/gorm"
)

type BookingMenuState struct {
	manager        *cli.CLIMenuStateManager
	handlerContext *handler.HandlerContext
	controller     controller.BookingControllerInterface
}

func (menu *BookingMenuState) Render() {
	fmt.Println()
	fmt.Println("===== Booking Management =====")
	fmt.Println("1. Create new Booking")
	fmt.Println("2. List all Bookings")
	fmt.Println("3. Get detail of a Booking")
	fmt.Println("4. Update a Booking")
	fmt.Println("5. Delete a Booking")
	fmt.Println("6. Check Room Availability")
	fmt.Println("7. Get Bookings by TimeTable")
	fmt.Println("Type 'back' to return to previous menu")
	fmt.Println("=============================")
}

func (menu *BookingMenuState) HandleUserInput(input string) error {
	err := menu.handlerContext.HandleInput(strings.TrimSpace(input))
	if err != nil {
		fmt.Println("Error handling user input:", err)
	}
	if strings.ToLower(input) == "back" {
		util.PressEnterToContinue()
	}
	return err
}

func NewBookingMenuState(db *gorm.DB, manager *cli.CLIMenuStateManager, spaceManagementMenu *SpaceManagementState) *BookingMenuState {
	if db == nil {
		fmt.Println("Error: Database connection is nil")
		return &BookingMenuState{
			manager:        manager,
			handlerContext: handler.NewHandlerContext(),
		}
	}

	controllerManager := controller.GetSpaceManagementInstance(db)
	if controllerManager == nil {
		fmt.Println("Error: Space Management Controller Facade is nil")
		return &BookingMenuState{
			manager:        manager,
			handlerContext: handler.NewHandlerContext(),
		}
	}

	bookingController := controllerManager.Booking
	if bookingController == nil {
		fmt.Println("Error: Booking controller is nil")
		return &BookingMenuState{
			manager:        manager,
			handlerContext: handler.NewHandlerContext(),
		}
	}

	handlerContext := handler.NewHandlerContext()
	
	//Standard Handlers
	listHandler := handler.NewListHandlerStrategy[model.Booking](bookingController)
	getHandler := handler.NewRetrieveByIDHandlerStrategy[model.Booking](bookingController)
	deleteHandler := handler.NewDeleteHandlerStrategy[model.Booking](bookingController)
	backHandler := handler.NewChangeMenuHandlerStrategy(manager, spaceManagementMenu)

	//Custom Handlers
	insertHandler := spaceManagementHandler.NewAddBookingHandlerStrategy(bookingController)
	updateHandler := spaceManagementHandler.NewUpdateBookingHandlerStrategy(bookingController)
	checkAvailabilityHandler := spaceManagementHandler.NewCheckRoomAvailabilityHandlerStrategy(bookingController)
	getByTimeTableHandler := spaceManagementHandler.NewGetBookingsByTimeTableHandlerStrategy(bookingController)

	handlerContext.AddHandler("1", "Create new Booking", insertHandler)
	handlerContext.AddHandler("2", "List all Bookings", listHandler)
	handlerContext.AddHandler("3", "Get detail of a Booking", getHandler)
	handlerContext.AddHandler("4", "Update a Booking", updateHandler)
	handlerContext.AddHandler("5", "Delete a Booking", deleteHandler)
	handlerContext.AddHandler("6", "Check Room Availability", checkAvailabilityHandler)
	handlerContext.AddHandler("7", "Get Bookings by TimeTable", getByTimeTableHandler)
	handlerContext.AddHandler("back", "Back to main menu", backHandler)

	return &BookingMenuState{
		manager:        manager,
		handlerContext: handlerContext,
		controller:     bookingController,
	}
}