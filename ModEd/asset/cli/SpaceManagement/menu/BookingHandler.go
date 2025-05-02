package menu

import (
	controller "ModEd/asset/controller"
	model "ModEd/asset/model"
	"ModEd/asset/util"
	"ModEd/core/cli"
	"ModEd/core/handler"
	"fmt"
	"strconv"
	"time"

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
	fmt.Println("8. Seed Bookings Data")
	fmt.Println("Type 'back' to return to previous menu")
	fmt.Println("=============================")
}

func (menu *BookingMenuState) HandleUserInput(input string) error {
	err := menu.handlerContext.HandleInput(input)
	if err != nil {
		fmt.Println("Error handling user input:", err)
	}
	if input == "back" {
		util.PressEnterToContinue()
	}
	return err
}

// CreateBookingHandler handles the booking creation process
type CreateBookingHandler struct {
	controller controller.BookingControllerInterface
}

func (h *CreateBookingHandler) Execute() error {
	fmt.Println("=== Create New Booking ===")
	
	var booking model.Booking
	
	fmt.Print("Enter TimeTable ID: ")
	var timeTableIDStr string
	fmt.Scanln(&timeTableIDStr)
	timeTableID, err := strconv.ParseUint(timeTableIDStr, 10, 32)
	if err != nil {
		fmt.Println("Invalid TimeTable ID")
		return err
	}
	booking.TimeTableID = uint(timeTableID)
	
	fmt.Print("Enter User ID: ")
	var userIDStr string
	fmt.Scanln(&userIDStr)
	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		fmt.Println("Invalid User ID")
		return err
	}
	booking.UserID = uint(userID)
	
	fmt.Print("Enter User Role (STUDENT/TEACHER/STAFF): ")
	var role string
	fmt.Scanln(&role)
	switch role {
	case "STUDENT":
		booking.UserRole = model.ROLE_STUDENT
	case "TEACHER":
		booking.UserRole = model.ROLE_ADVISOR
	case "STAFF":
		booking.UserRole = model.ROLE_ADMIN
	default:
		fmt.Println("Invalid role. Setting default to STUDENT")
		booking.UserRole = model.ROLE_STUDENT
	}
	
	fmt.Print("Enter Event Name: ")
	booking.EventName = util.GetCommandInput()
	
	result, err := h.controller.CreateBooking(booking)
	if err != nil {
		fmt.Println("Error creating booking:", err)
		return err
	}
	
	fmt.Println("Booking created successfully with ID:", result.ID)
	util.PressEnterToContinue()
	return nil
}

type ListBookingsHandler struct {
	controller controller.BookingControllerInterface
}

func (h *ListBookingsHandler) Execute() error {
	fmt.Println("=== List All Bookings ===")
	
	bookings, err := h.controller.ListBookings(map[string]interface{}{})
	if err != nil {
		fmt.Println("Error retrieving bookings:", err)
		return err
	}
	
	if len(bookings) == 0 {
		fmt.Println("No bookings found")
	} else {
		fmt.Println("Found", len(bookings), "booking(s):")
		for i, booking := range bookings {
			fmt.Printf("%d. Booking ID: %d, TimeTable ID: %d, User ID: %d, Role: %s, Event: %s\n",
				i+1, booking.ID, booking.TimeTableID, booking.UserID, booking.UserRole, booking.EventName)
		}
	}
	
	util.PressEnterToContinue()
	return nil
}

type GetBookingHandler struct {
	controller controller.BookingControllerInterface
}

func (h *GetBookingHandler) Execute() error {
	fmt.Println("=== Get Booking Details ===")
	
	fmt.Print("Enter Booking ID: ")
	var idStr string
	fmt.Scanln(&idStr)
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		fmt.Println("Invalid Booking ID")
		return err
	}
	
	booking, err := h.controller.GetBooking(uint(id))
	if err != nil {
		fmt.Println("Error retrieving booking:", err)
		return err
	}
	
	fmt.Println("Booking Details:")
	fmt.Println("ID:", booking.ID)
	fmt.Println("Created At:", booking.CreatedAt)
	fmt.Println("Updated At:", booking.UpdatedAt)
	fmt.Println("TimeTable ID:", booking.TimeTableID)
	fmt.Println("Time Period:", booking.TimeTable.StartDate.Format("2006-01-02 15:04"), "to", booking.TimeTable.EndDate.Format("2006-01-02 15:04"))
	fmt.Println("Room:", booking.TimeTable.Room.RoomName)
	fmt.Println("User ID:", booking.UserID)
	fmt.Println("User Role:", booking.UserRole)
	fmt.Println("Event Name:", booking.EventName)
	
	util.PressEnterToContinue()
	return nil
}

type UpdateBookingHandler struct {
	controller controller.BookingControllerInterface
}

func (h *UpdateBookingHandler) Execute() error {
	fmt.Println("=== Update Booking ===")
	
	fmt.Print("Enter Booking ID to update: ")
	var idStr string
	fmt.Scanln(&idStr)
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		fmt.Println("Invalid Booking ID")
		return err
	}
	
	booking, err := h.controller.GetBooking(uint(id))
	if err != nil {
		fmt.Println("Error retrieving booking:", err)
		return err
	}
	
	fmt.Print("Enter new Event Name (current: "+booking.EventName+") or leave empty to keep current: ")
	newEventName := util.GetCommandInput()
	if newEventName != "" {
		booking.EventName = newEventName
	}
	
	fmt.Print("Enter new User Role (STUDENT/TEACHER/STAFF) (current: "+string(booking.UserRole)+") or leave empty to keep current: ")
	newRole := util.GetCommandInput()
	if newRole != "" {
		switch newRole {
		case "STUDENT":
			booking.UserRole = model.ROLE_STUDENT
		case "TEACHER":
			booking.UserRole = model.ROLE_ADVISOR
		case "STAFF":
			booking.UserRole = model.ROLE_ADMIN
		default:
			fmt.Println("Invalid role. Keeping current role.")
		}
	}
	
	err = h.controller.UpdateBooking(booking)
	if err != nil {
		fmt.Println("Error updating booking:", err)
		return err
	}
	
	fmt.Println("Booking updated successfully")
	util.PressEnterToContinue()
	return nil
}

type DeleteBookingHandler struct {
	controller controller.BookingControllerInterface
}

func (h *DeleteBookingHandler) Execute() error {
	fmt.Println("=== Delete Booking ===")
	
	fmt.Print("Enter Booking ID to delete: ")
	var idStr string
	fmt.Scanln(&idStr)
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		fmt.Println("Invalid Booking ID")
		return err
	}
	
	err = h.controller.DeleteBooking(uint(id))
	if err != nil {
		fmt.Println("Error deleting booking:", err)
		return err
	}
	
	fmt.Println("Booking deleted successfully")
	util.PressEnterToContinue()
	return nil
}

type CheckRoomAvailabilityHandler struct {
	controller controller.BookingControllerInterface
}

func (h *CheckRoomAvailabilityHandler) Execute() error {
	fmt.Println("=== Check Room Availability ===")
	
	fmt.Print("Enter Room ID: ")
	var roomIDStr string
	fmt.Scanln(&roomIDStr)
	roomID, err := strconv.ParseUint(roomIDStr, 10, 32)
	if err != nil {
		fmt.Println("Invalid Room ID")
		return err
	}
	
	fmt.Print("Enter Start Date (YYYY-MM-DD HH:MM): ")
	startDateStr := util.GetCommandInput()
	startDate, err := time.Parse("2006-01-02 15:04", startDateStr)
	if err != nil {
		fmt.Println("Invalid date format. Please use YYYY-MM-DD HH:MM")
		return err
	}
	
	fmt.Print("Enter End Date (YYYY-MM-DD HH:MM): ")
	endDateStr := util.GetCommandInput()
	endDate, err := time.Parse("2006-01-02 15:04", endDateStr)
	if err != nil {
		fmt.Println("Invalid date format. Please use YYYY-MM-DD HH:MM")
		return err
	}
	
	available, err := h.controller.CheckRoomAvailability(uint(roomID), startDate, endDate)
	if err != nil {
		fmt.Println("Error checking availability:", err)
		return err
	}
	
	if available {
		fmt.Println("Room is available for the selected time period")
	} else {
		fmt.Println("Room is NOT available for the selected time period")
	}
	
	util.PressEnterToContinue()
	return nil
}

type GetBookingsByTimeTableHandler struct {
	controller controller.BookingControllerInterface
}

func (h *GetBookingsByTimeTableHandler) Execute() error {
	fmt.Println("=== Get Bookings by TimeTable ===")
	
	fmt.Print("Enter TimeTable ID: ")
	var timeTableIDStr string
	fmt.Scanln(&timeTableIDStr)
	timeTableID, err := strconv.ParseUint(timeTableIDStr, 10, 32)
	if err != nil {
		fmt.Println("Invalid TimeTable ID")
		return err
	}
	
	bookings, err := h.controller.GetBookingsByTimeTable(uint(timeTableID))
	if err != nil {
		fmt.Println("Error retrieving bookings:", err)
		return err
	}
	
	if len(bookings) == 0 {
		fmt.Println("No bookings found for this timetable")
	} else {
		fmt.Println("Found", len(bookings), "booking(s):")
		for i, booking := range bookings {
			fmt.Printf("%d. Booking ID: %d, User ID: %d, Role: %s, Event: %s\n",
				i+1, booking.ID, booking.UserID, booking.UserRole, booking.EventName)
		}
	}
	
	util.PressEnterToContinue()
	return nil
}

type SeedBookingsHandler struct {
	controller controller.BookingControllerInterface
}

func (h *SeedBookingsHandler) Execute() error {
	fmt.Println("=== Seed Bookings Data ===")
	
	fmt.Print("Enter path to bookings JSON file: ")
	path := util.GetCommandInput()
	
	bookings, err := h.controller.SeedBookingsDatabase(path)
	if err != nil {
		fmt.Println("Error seeding bookings database:", err)
		return err
	}
	
	fmt.Println("Successfully seeded", len(bookings), "bookings")
	util.PressEnterToContinue()
	return nil
}

func NewBookingMenuState(db *gorm.DB, manager *cli.CLIMenuStateManager, spaceManagementMenu *SpaceManagementState) *BookingMenuState {
	if db == nil {
		fmt.Println("Error: Database connection is nil")
		return &BookingMenuState{
			manager:        manager,
			handlerContext: handler.NewHandlerContext(),
		}
	}

	controllerFacade := controller.GetSpaceManagementInstance(db)
	if controllerFacade == nil {
		fmt.Println("Error: Space Management Controller Facade is nil")
		return &BookingMenuState{
			manager:        manager,
			handlerContext: handler.NewHandlerContext(),
		}
	}

	bookingController := controllerFacade.Booking
	if bookingController == nil {
		fmt.Println("Error: Booking controller is nil")
		return &BookingMenuState{
			manager:        manager,
			handlerContext: handler.NewHandlerContext(),
		}
	}

	handlerContext := handler.NewHandlerContext()
	
	createHandler := &CreateBookingHandler{controller: bookingController}
	listHandler := &ListBookingsHandler{controller: bookingController}
	getHandler := &GetBookingHandler{controller: bookingController}
	updateHandler := &UpdateBookingHandler{controller: bookingController}
	deleteHandler := &DeleteBookingHandler{controller: bookingController}
	checkAvailabilityHandler := &CheckRoomAvailabilityHandler{controller: bookingController}
	getByTimeTableHandler := &GetBookingsByTimeTableHandler{controller: bookingController}
	seedHandler := &SeedBookingsHandler{controller: bookingController}
	
	backHandler := handler.NewChangeMenuHandlerStrategy(manager, spaceManagementMenu)

	handlerContext.AddHandler("1", "Create new Booking", createHandler)
	handlerContext.AddHandler("2", "List all Bookings", listHandler)
	handlerContext.AddHandler("3", "Get detail of a Booking", getHandler)
	handlerContext.AddHandler("4", "Update a Booking", updateHandler)
	handlerContext.AddHandler("5", "Delete a Booking", deleteHandler)
	handlerContext.AddHandler("6", "Check Room Availability", checkAvailabilityHandler)
	handlerContext.AddHandler("7", "Get Bookings by TimeTable", getByTimeTableHandler)
	handlerContext.AddHandler("8", "Seed Bookings Data", seedHandler)
	handlerContext.AddHandler("back", "Back to main menu", backHandler)

	return &BookingMenuState{
		manager:        manager,
		handlerContext: handlerContext,
		controller:     bookingController,
	}
}