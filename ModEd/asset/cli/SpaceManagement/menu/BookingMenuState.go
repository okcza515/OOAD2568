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
	// "strconv"
	"strings"
	// "time"

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

// type CreateBookingHandler struct {
// 	controller controller.BookingControllerInterface
// }

// func (h *CreateBookingHandler) Execute() error {
// 	fmt.Println("===== Create New Booking =====")

	
// 	var booking model.Booking
	
// 	fmt.Print("Enter TimeTable ID: ")
// 	var timeTableIDStr string
// 	fmt.Scanln(&timeTableIDStr)
// 	timeTableID, err := strconv.ParseUint(strings.TrimSpace(timeTableIDStr), 10, 32)
// 	if err != nil {
// 		fmt.Println("Invalid TimeTable ID")
// 		return err
// 	}
// 	booking.TimeTableID = uint(timeTableID)
	
// 	fmt.Print("Enter User ID: ")
// 	var userIDStr string
// 	fmt.Scanln(&userIDStr)
// 	userID, err := strconv.ParseUint(strings.TrimSpace(userIDStr), 10, 32)
// 	if err != nil {
// 		fmt.Println("Invalid User ID")
// 		return err
// 	}
// 	booking.UserID = uint(userID)
	
// 	fmt.Print("Enter User Role (STUDENT/TEACHER/STAFF): ")
// 	var role string
// 	fmt.Scanln(&role)
// 	role = strings.ToUpper(strings.TrimSpace(role))
// 	switch role {
// 	case "STUDENT":
// 		booking.UserRole = model.ROLE_STUDENT
// 	case "TEACHER":
// 		booking.UserRole = model.ROLE_ADVISOR
// 	case "STAFF":
// 		booking.UserRole = model.ROLE_ADMIN
// 	default:
// 		fmt.Println("Invalid role. Setting default to STUDENT")
// 		booking.UserRole = model.ROLE_STUDENT
// 	}
	
// 	fmt.Print("Enter Event Name: ")
// 	booking.EventName = util.GetCommandInput()
	
// 	err = h.controller.Insert(booking)
// 	if err != nil {
// 		fmt.Println("Error creating booking:", err)
// 		return err
// 	}
	
// 	fmt.Println("Booking created successfully")
// 	util.PressEnterToContinue()
// 	return nil
// }

// type ListBookingsHandler struct {
// 	controller controller.BookingControllerInterface
// }

// func (h *ListBookingsHandler) Execute() error {
// 	fmt.Println("===== List All Bookings =====")
	
// 	bookings, err := h.controller.List(map[string]interface{}{})
// 	if err != nil {
// 		fmt.Println("Error retrieving bookings:", err)
// 		return err
// 	}
	
// 	if len(bookings) == 0 {
// 		fmt.Println("No bookings found")
// 	} else {
// 		fmt.Printf("Found %d booking(s):\n\n", len(bookings))
// 		fmt.Println("===============================================================")
// 		fmt.Println(" No | Booking ID | TimeTable  | User ID |  Role   | Event Name ")
// 		fmt.Println("---------------------------------------------------------------")
// 		for i, booking := range bookings {
// 			fmt.Printf(" %2d | %-9d | %-11d | %-7d | %-7s | %-21s \n", 
// 				i+1, booking.ID, booking.TimeTableID, booking.UserID, booking.UserRole, truncateString(booking.EventName, 21))
// 		}
// 		fmt.Println("===============================================================")
// 	}
	
// 	util.PressEnterToContinue()
// 	return nil
// }

// func truncateString(str string, length int) string {
// 	if len(str) <= length {
// 		return str
// 	}
// 	return str[:length-3] + "..."
// }

// type GetBookingHandler struct {
// 	controller controller.BookingControllerInterface
// }

// func (h *GetBookingHandler) Execute() error {
// 	fmt.Println("===== Booking Details =====")

	
// 	fmt.Print("Enter Booking ID: ")
// 	var idStr string
// 	fmt.Scanln(&idStr)
// 	id, err := strconv.ParseUint(strings.TrimSpace(idStr), 10, 32)
// 	if err != nil {
// 		fmt.Println("Invalid Booking ID")
// 		return err
// 	}
	
// 	booking, err := h.controller.RetrieveByID(uint(id))
// 	if err != nil {
// 		fmt.Println("Error retrieving booking:", err)
// 		return err
// 	}
	
// 	fmt.Println("==================================================================")
// 	fmt.Printf(" BOOKING #%-43d \n", booking.ID)
// 	fmt.Println("------------------------------------------------------------------")
// 	fmt.Printf(" Created:                      %-25s \n", booking.CreatedAt.Format("2006-01-02 15:04:05"))
// 	fmt.Printf(" Updated:                      %-25s \n", booking.UpdatedAt.Format("2006-01-02 15:04:05"))
// 	fmt.Printf(" TimeTable ID:                 %-25d \n", booking.TimeTableID)
// 	fmt.Println("------------------------------------------------------------------")
// 	fmt.Printf(" Start:                        %-25s \n", booking.TimeTable.StartDate.Format("2006-01-02 15:04"))
// 	fmt.Printf(" End:                          %-25s \n", booking.TimeTable.EndDate.Format("2006-01-02 15:04"))
// 	fmt.Printf(" Room:                         %-25s \n", booking.TimeTable.Room.RoomName)
// 	fmt.Println("------------------------------------------------------------------")
// 	fmt.Printf(" User ID:                      %-25d \n", booking.UserID)
// 	fmt.Printf(" User Role:                    %-25s \n", booking.UserRole)
// 	fmt.Printf(" Event Name:                   %-25s \n", booking.EventName)
// 	fmt.Println("==================================================================")
	
// 	util.PressEnterToContinue()
// 	return nil
// }

// type UpdateBookingHandler struct {
// 	controller controller.BookingControllerInterface
// }

// func (h *UpdateBookingHandler) Execute() error {
// 	fmt.Println("===== Update Booking =====")

	
// 	fmt.Print("Enter Booking ID to update: ")
// 	var idStr string
// 	fmt.Scanln(&idStr)
// 	id, err := strconv.ParseUint(strings.TrimSpace(idStr), 10, 32)
// 	if err != nil {
// 		fmt.Println("Invalid Booking ID")
// 		return err
// 	}
	
// 	booking, err := h.controller.RetrieveByID(uint(id))
// 	if err != nil {
// 		fmt.Println("Error retrieving booking:", err)
// 		return err
// 	}
	
// 	fmt.Printf("Current Event Name: %s\n", booking.EventName)
// 	fmt.Print("Enter new Event Name (or press Enter to keep current): ")
// 	newEventName := util.GetCommandInput()
// 	if newEventName != "" {
// 		booking.EventName = newEventName
// 	}
	
// 	fmt.Printf("Current User Role: %s\n", booking.UserRole)
// 	fmt.Print("Enter new User Role (STUDENT/ADVISOR/ADMIN) (or press Enter to keep current): ")
// 	newRole := util.GetCommandInput()
// 	if newRole != "" {
// 		newRole = strings.ToUpper(strings.TrimSpace(newRole))
// 		switch newRole {
// 		case "STUDENT":
// 			booking.UserRole = model.ROLE_STUDENT
// 		case "ADVISOR":
// 			booking.UserRole = model.ROLE_ADVISOR
// 		case "ADMIN":
// 			booking.UserRole = model.ROLE_ADMIN
// 		default:
// 			fmt.Println("Invalid role. Keeping current role.")
// 		}
// 	}
	
// 	err = h.controller.UpdateBooking(booking)
// 	if err != nil {
// 		fmt.Println("Error updating booking:", err)
// 		return err
// 	}
	
// 	fmt.Println("Booking updated successfully")
// 	util.PressEnterToContinue()
// 	return nil
// }

// type DeleteBookingHandler struct {
// 	controller controller.BookingControllerInterface
// }

// func (h *DeleteBookingHandler) Execute() error {
// 	fmt.Println("===== Delete Booking =====")

	
// 	fmt.Print("Enter Booking ID to delete: ")
// 	var idStr string
// 	fmt.Scanln(&idStr)
// 	id, err := strconv.ParseUint(strings.TrimSpace(idStr), 10, 32)
// 	if err != nil {
// 		fmt.Println("Invalid Booking ID")
// 		return err
// 	}
	
// 	fmt.Print("Are you sure you want to delete this booking? (y/n): ")
// 	var confirm string
// 	fmt.Scanln(&confirm)
// 	if strings.ToLower(strings.TrimSpace(confirm)) != "y" {
// 		fmt.Println("Deletion cancelled")
// 		util.PressEnterToContinue()
// 		return nil
// 	}
	
// 	err = h.controller.DeleteByID(uint(id))
// 	if err != nil {
// 		fmt.Println("Error deleting booking:", err)
// 		return err
// 	}
	
// 	fmt.Println("Booking deleted successfully")
// 	util.PressEnterToContinue()
// 	return nil
// }

// type CheckRoomAvailabilityHandler struct {
// 	controller controller.BookingControllerInterface
// }

// func (h *CheckRoomAvailabilityHandler) Execute() error {
// 	fmt.Println("===== Check Room Availability =====")

	
// 	fmt.Print("Enter Room ID: ")
// 	var roomIDStr string
// 	fmt.Scanln(&roomIDStr)
// 	roomID, err := strconv.ParseUint(strings.TrimSpace(roomIDStr), 10, 32)
// 	if err != nil {
// 		fmt.Println("Invalid Room ID")
// 		return err
// 	}
	
// 	fmt.Print("Enter Start Date (YYYY-MM-DD): ")
// 	startDateStr := strings.TrimSpace(util.GetCommandInput())
// 	fmt.Print("Enter Start Time (HH:MM): ")
// 	startTimeStr := strings.TrimSpace(util.GetCommandInput())
	
// 	startDateTime := startDateStr + " " + startTimeStr
// 	startDate, err := time.Parse("2006-01-02 15:04", startDateTime)
// 	if err != nil {
// 		fmt.Println("Invalid date/time format. Please use YYYY-MM-DD for date and HH:MM for time")
// 		return err
// 	}
	
// 	fmt.Print("Enter End Date (YYYY-MM-DD): ")
// 	endDateStr := strings.TrimSpace(util.GetCommandInput())
// 	fmt.Print("Enter End Time (HH:MM): ")
// 	endTimeStr := strings.TrimSpace(util.GetCommandInput())
	
// 	endDateTime := endDateStr + " " + endTimeStr
// 	endDate, err := time.Parse("2006-01-02 15:04", endDateTime)
// 	if err != nil {
// 		fmt.Println("Invalid date/time format. Please use YYYY-MM-DD for date and HH:MM for time")
// 		return err
// 	}
	
// 	available, err := h.controller.CheckRoomAvailability(uint(roomID), startDate, endDate)
// 	if err != nil {
// 		fmt.Println("Error checking availability:", err)
// 		return err
// 	}
	
// 	fmt.Println("==================================================================")
// 	fmt.Printf(" Room #%-5d                                                 \n", roomID)
// 	fmt.Printf(" Period: %-15s to %-15s               \n", 
// 		startDate.Format("2006-01-02 15:04"), 
// 		endDate.Format("2006-01-02 15:04"))
// 	fmt.Println("------------------------------------------------------------------")
// 	if available {
// 		fmt.Println("              ROOM IS AVAILABLE FOR THIS PERIOD                ")
// 	} else {
// 		fmt.Println("              ROOM IS NOT AVAILABLE FOR THIS PERIOD            ")
// 	}
// 	fmt.Println("==================================================================")
	
// 	util.PressEnterToContinue()
// 	return nil
// }

// type GetBookingsByTimeTableHandler struct {
// 	controller controller.BookingControllerInterface
// }

// func (h *GetBookingsByTimeTableHandler) Execute() error {
// 	fmt.Println("===== Bookings By TimeTable ID =====")

	
// 	fmt.Print("Enter TimeTable ID: ")
// 	var timeTableIDStr string
// 	fmt.Scanln(&timeTableIDStr)
// 	timeTableID, err := strconv.ParseUint(strings.TrimSpace(timeTableIDStr), 10, 32)
// 	if err != nil {
// 		fmt.Println("Invalid TimeTable ID")
// 		return err
// 	}
	
// 	bookings, err := h.controller.GetBookingsByTimeTable(uint(timeTableID))
// 	if err != nil {
// 		fmt.Println("Error retrieving bookings:", err)
// 		return err
// 	}
	
// 	if len(bookings) == 0 {
// 		fmt.Println("No bookings found for this timetable")
// 	} else {
// 		fmt.Printf("Found %d booking(s) for TimeTable ID %d:\n\n", len(bookings), timeTableID)
// 		fmt.Println("==========================================================")
// 		fmt.Println(" No | Booking ID | User ID |  Role   |      Event Name    ")
// 		fmt.Println("----------------------------------------------------------")
// 		for i, booking := range bookings {
// 			fmt.Printf(" %2d | %-9d | %-7d | %-7s | %-21s \n", 
// 				i+1, booking.ID, booking.UserID, booking.UserRole, truncateString(booking.EventName, 21))
// 		}
// 		fmt.Println("==========================================================")
// 	}
	
// 	util.PressEnterToContinue()
// 	return nil
// }


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
	
	// checkAvailabilityHandler := &CheckRoomAvailabilityHandler{controller: bookingController}
	// getByTimeTableHandler := &GetBookingsByTimeTableHandler{controller: bookingController}
	// seedHandler := &SeedBookingsHandler{controller: bookingController}
	// deleteSeedHandler := &DeleteSeedBookingsHandler{controller: bookingController}
	
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
	// handlerContext.AddHandler("8", "Seed Bookings Data", seedHandler)
	// handlerContext.AddHandler("9", "Delete Seed Bookings Data", deleteSeedHandler)
	handlerContext.AddHandler("back", "Back to main menu", backHandler)

	return &BookingMenuState{
		manager:        manager,
		handlerContext: handlerContext,
		controller:     bookingController,
	}
}