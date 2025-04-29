package handler

import (
	controller "ModEd/asset/controller"
	model "ModEd/asset/model"
	"ModEd/asset/util"
	"fmt"
	"strconv"
	"time"
)

func printBookingOption() {
	fmt.Println("========== Booking Management ==========")
	fmt.Println("1. Book a Room")
	fmt.Println("2. Cancel Booking")
	fmt.Println("3. Update Booking")
	fmt.Println("4. Get Room Bookings")
	fmt.Println("5. Get Booking Details")
	fmt.Println("6. Reset All Bookings")
	fmt.Println("7. Check Room Availability")
	fmt.Println("8. Get Available Rooms")
	fmt.Println("9. Reset Time Slots for a Room")
	fmt.Println("10. Seek Bookings Data")
	fmt.Println("Type 'back' to return to previous menu")
	fmt.Println("========================================")
}

func BookingHandler(facade *controller.SpaceManagementControllerFacade) {
	inputBuffer := ""

	for inputBuffer != "back" {
		util.ClearScreen()
		util.PrintBanner()
		printBookingOption()
		inputBuffer = util.GetCommandInput()

		switch inputBuffer {
		case "1":
			handleBookRoom(facade)
		case "2":
			handleCancelBooking(facade)
		case "3":
			handleUpdateBooking(facade)
		case "4":
			handleGetRoomBookings(facade)
		case "5":
			handleGetBookingDetails(facade)
		case "6":
			handleResetAllBookings(facade)
		case "7":
			handleCheckRoomAvailability(facade)
		case "8":
			handleGetAvailableRooms(facade)
		case "9":
			handleResetTimeSlots(facade)
		case "10":
			fmt.Println("Seed Bookings Data")
			data, err := facade.Booking.SeedBookingsDatabase("data/asset/Booking.json")
			if err != nil {
				fmt.Println("Failed to seed Booking data", err)
				util.PressEnterToContinue()
			} else {
				fmt.Println("Booking data seeded successfully")
				for _, booking := range data {
					fmt.Println(booking)
				}
				util.PressEnterToContinue()
			}
		case "back":
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
			util.PressEnterToContinue()
		}
	}
}

func getRoomTypeSelection() (*model.RoomTypeEnum, error) {
	fmt.Println("Select the type of room:")
	fmt.Println("1. Lecture Room")
	fmt.Println("2. Lab Room")
	fmt.Println("3. Meeting Room")
	fmt.Print("Enter your choice (1/2/3): ")
	roomTypeChoice := util.GetCommandInput()

	var roomType model.RoomTypeEnum

	switch roomTypeChoice {
	case "1":
		roomType = model.ROOM_LECTURE_ROOM
	case "2":
		roomType = model.ROOM_LAB_ROOM
	case "3":
		roomType = model.ROOM_MEETING_ROOM
	default:
		return nil, fmt.Errorf("invalid choice")
	}

	return &roomType, nil
}

func getValidRole() (model.Role, error) {
	fmt.Print("Enter your role (Student/Advisor/Admin): ")
	roleStr := util.GetCommandInput()

	var role model.Role
	switch roleStr {
	case "Student", "student":
		role = model.ROLE_STUDENT
	case "Advisor", "advisor":
		role = model.ROLE_ADVISOR
	case "Admin", "admin":
		role = model.ROLE_ADMIN
	default:
		return "", fmt.Errorf("invalid role. Please enter a valid role (Student/Advisor/Admin)")
	}

	return role, nil
}

func getValidDateTimeInput(promptPrefix string) (time.Time, error) {
	fmt.Printf("For %s time:\n", promptPrefix)
	fmt.Print("Enter date (YYYY-MM-DD): ")
	dateStr := util.GetCommandInput()
	fmt.Print("Enter time (HH:MM): ")
	timeStr := util.GetCommandInput()

	fullStr := dateStr + " " + timeStr
	dateTime, err := time.Parse("2006-01-02 15:04", fullStr)
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid date or time format")
	}

	return dateTime, nil
}

func validateDateTimeRange(startTime, endTime time.Time) error {
	if startTime.Before(time.Now()) {
		return fmt.Errorf("start time cannot be in the past")
	}

	if endTime.Before(startTime) || endTime.Equal(startTime) {
		return fmt.Errorf("end time must be after start time")
	}

	return nil
}

func handleBookRoom(facade *controller.SpaceManagementControllerFacade) {
	fmt.Println("====== Book a Room ======")

	roomType, err := getRoomTypeSelection()
	if err != nil {
		fmt.Println(err)
		util.PressEnterToContinue()
		return
	}

	fmt.Println("You have selected:", roomType.TypeRoomString())

	rooms, err := facade.Room.GetAll()
	if err != nil {
		fmt.Println("Error retrieving rooms:", err)
		util.PressEnterToContinue()
		return
	}

	fmt.Println("\nAvailable rooms of type", roomType.TypeRoomString(), ":")
	hasRooms := false
	for _, room := range *rooms {
		if room.RoomType == *roomType && !room.IsRoomOutOfService {
			fmt.Printf("ID: %d, Name: %s, Building: %s, Floor: %d, Capacity: %d\n", 
				room.ID, room.RoomName, room.Building, room.Floor, room.Capacity)
			hasRooms = true
		}
	}

	if !hasRooms {
		fmt.Println("No rooms available of the selected type.")
		util.PressEnterToContinue()
		return
	}

	fmt.Print("\nEnter room ID: ")
	roomIDInput := util.GetCommandInput()
	roomID, err := strconv.Atoi(roomIDInput)
	if err != nil {
		fmt.Println("Invalid room ID.")
		util.PressEnterToContinue()
		return
	}

	var validRoom bool
	for _, room := range *rooms {
		if room.ID == uint(roomID) && room.RoomType == *roomType {
			validRoom = true
			break
		}
	}

	if !validRoom {
		fmt.Println("The selected room ID does not exist or does not match the selected room type.")
		util.PressEnterToContinue()
		return
	}

	startTime, err := getValidDateTimeInput("start")
	if err != nil {
		fmt.Println(err)
		util.PressEnterToContinue()
		return
	}

	endTime, err := getValidDateTimeInput("end")
	if err != nil {
		fmt.Println(err)
		util.PressEnterToContinue()
		return
	}

	if err := validateDateTimeRange(startTime, endTime); err != nil {
		fmt.Println(err)
		util.PressEnterToContinue()
		return
	}

	role, err := getValidRole()
	if err != nil {
		fmt.Println(err)
		util.PressEnterToContinue()
		return
	}

	fmt.Print("Enter event name: ")
	eventName := util.GetCommandInput()
	if eventName == "" {
		fmt.Println("Event name cannot be empty.")
		util.PressEnterToContinue()
		return
	}

	fmt.Println("\nConfirm booking details:")
	fmt.Printf("Room ID: %d, User ID: %d, Role: %s, Event: %s\n", roomID, 1, role, eventName)
	fmt.Printf("Time: %s to %s\n", startTime.Format("2006-01-02 15:04"), endTime.Format("2006-01-02 15:04"))
	fmt.Print("Confirm booking? (y/n): ")
	confirm := util.GetCommandInput()
	if confirm != "y" && confirm != "Y" {
		fmt.Println("Booking canceled.")
		util.PressEnterToContinue()
		return
	}

	fmt.Println("Processing booking request...")

	booking, err := facade.Booking.BookRoom(uint(roomID), uint(1), role, eventName, startTime, endTime)

	fmt.Println("\n----- BOOKING RESULT -----")
	if err != nil {
		fmt.Println("Booking failed:", err)
	} else {
		if booking == nil {
			fmt.Println("Warning: Booking succeeded but returned nil booking object")
		} else {
			fmt.Printf("Room booked successfully! Booking ID: %d\n", booking.ID)
			fmt.Printf("Details: Room ID %d, Event '%s'\n", roomID, eventName)
			fmt.Printf("Time: %s to %s\n", startTime.Format("2006-01-02 15:04"), endTime.Format("2006-01-02 15:04"))
		}
	}
	fmt.Println("--------------------------")

	util.PressEnterToContinue()
}

func handleCancelBooking(facade *controller.SpaceManagementControllerFacade) {
	fmt.Println("====== Cancel Booking ======")
	fmt.Print("Enter booking ID: ")
	bookingIDInput := util.GetCommandInput()
	bookingID, err := strconv.Atoi(bookingIDInput)
	if err != nil {
		fmt.Println("Invalid booking ID.")
		util.PressEnterToContinue()
		return
	}

	booking, err := facade.Booking.GetBookingDetails(uint(bookingID))
	if err != nil {
		fmt.Println("Error fetching booking details:", err)
		util.PressEnterToContinue()
		return
	}

	fmt.Println("\nBooking details to cancel:")
	fmt.Printf("Booking ID: %d\n", booking.ID)
	fmt.Printf("Event: %s\n", booking.EventName)
	fmt.Printf("Start time: %s\n", booking.TimeTable.StartDate.Format("2006-01-02 15:04"))
	fmt.Printf("End time: %s\n", booking.TimeTable.EndDate.Format("2006-01-02 15:04"))

	fmt.Print("\nAre you sure you want to cancel this booking? (y/n): ")
	confirm := util.GetCommandInput()
	if confirm != "y" && confirm != "Y" {
		fmt.Println("Cancellation aborted.")
		util.PressEnterToContinue()
		return
	}

	err = facade.Booking.CancelBooking(uint(bookingID))
	if err != nil {
		fmt.Println("Cancel failed:", err)
	} else {
		fmt.Println("Booking canceled successfully.")
	}
	util.PressEnterToContinue()
}

func handleUpdateBooking(facade *controller.SpaceManagementControllerFacade) {
	fmt.Println("====== Update a Booking ======")

	fmt.Print("Enter booking ID: ")
	bookingIDInput := util.GetCommandInput()
	bookingID, err := strconv.Atoi(bookingIDInput)
	if err != nil {
		fmt.Println("Invalid booking ID.")
		util.PressEnterToContinue()
		return
	}

	currentBooking, err := facade.Booking.GetBookingDetails(uint(bookingID))
	if err != nil {
		fmt.Println("Error fetching booking details:", err)
		util.PressEnterToContinue()
		return
	}

	fmt.Println("\nCurrent booking details:")
	fmt.Printf("Event: %s\n", currentBooking.EventName)
	fmt.Printf("Start time: %s\n", currentBooking.TimeTable.StartDate.Format("2006-01-02 15:04"))
	fmt.Printf("End time: %s\n", currentBooking.TimeTable.EndDate.Format("2006-01-02 15:04"))

	fmt.Println("\nWhat would you like to update?")
	fmt.Println("1. Event name")
	fmt.Println("2. Start and end times")
	fmt.Println("3. Both")
	fmt.Print("Enter your choice (1/2/3): ")
	updateChoice := util.GetCommandInput()

	var eventNamePtr *string
	var startTimePtr, endTimePtr *time.Time

	if updateChoice == "1" || updateChoice == "3" {
		fmt.Print("Enter new event name: ")
		eventName := util.GetCommandInput()
		if eventName != "" {
			eventNamePtr = &eventName
		} else {
			fmt.Println("Event name cannot be empty. Keeping original name.")
		}
	}

	if updateChoice == "2" || updateChoice == "3" {
		startTime, err := getValidDateTimeInput("new start")
		if err != nil {
			fmt.Println(err)
			util.PressEnterToContinue()
			return
		}

		endTime, err := getValidDateTimeInput("new end")
		if err != nil {
			fmt.Println(err)
			util.PressEnterToContinue()
			return
		}

		if err := validateDateTimeRange(startTime, endTime); err != nil {
			fmt.Println(err)
			util.PressEnterToContinue()
			return
		}

		startTimePtr = &startTime
		endTimePtr = &endTime
	}

	fmt.Println("\nSummary of changes:")
	if eventNamePtr != nil {
		fmt.Printf("Event name: %s -> %s\n", currentBooking.EventName, *eventNamePtr)
	}
	if startTimePtr != nil {
		fmt.Printf("Start time: %s -> %s\n", 
			currentBooking.TimeTable.StartDate.Format("2006-01-02 15:04"), 
			startTimePtr.Format("2006-01-02 15:04"))
	}
	if endTimePtr != nil {
		fmt.Printf("End time: %s -> %s\n", 
			currentBooking.TimeTable.EndDate.Format("2006-01-02 15:04"), 
			endTimePtr.Format("2006-01-02 15:04"))
	}

	fmt.Print("\nConfirm these changes? (y/n): ")
	confirm := util.GetCommandInput()
	if confirm != "y" && confirm != "Y" {
		fmt.Println("Update canceled.")
		util.PressEnterToContinue()
		return
	}

	err = facade.Booking.UpdateBooking(uint(bookingID), eventNamePtr, startTimePtr, endTimePtr)
	if err != nil {
		fmt.Println("Failed to update booking:", err)
	} else {
		fmt.Println("Booking updated successfully.")
	}
	util.PressEnterToContinue()
}

func handleGetRoomBookings(facade *controller.SpaceManagementControllerFacade) {
	fmt.Println("====== Get Room Bookings ======")
	
	rooms, err := facade.Room.GetAll()
	if err != nil {
		fmt.Println("Error retrieving rooms:", err)
		util.PressEnterToContinue()
		return
	}

	fmt.Println("Available rooms:")
	for _, room := range *rooms {
		fmt.Printf("ID: %d, Name: %s, Type: %s, Building: %s\n", 
			room.ID, room.RoomName, room.RoomType, room.Building)
	}
	
	fmt.Print("\nEnter room ID: ")
	roomIDInput := util.GetCommandInput()
	roomID, err := strconv.Atoi(roomIDInput)
	if err != nil {
		fmt.Println("Invalid room ID.")
		util.PressEnterToContinue()
		return
	}

	bookings, err := facade.Booking.GetRoomBookings(uint(roomID))
	if err != nil {
		fmt.Println("Error retrieving bookings:", err)
		util.PressEnterToContinue()
		return
	}

	if len(bookings) == 0 {
		fmt.Println("No bookings found for Room ID", roomID)
		util.PressEnterToContinue()
		return
	}

	fmt.Printf("\nBookings for Room ID %d:\n", roomID)
	fmt.Println("--------------------------------------------------------")
	fmt.Printf("%-10s %-30s %-20s %-20s\n", "ID", "Event", "Start", "End")
	fmt.Println("--------------------------------------------------------")
	for _, booking := range bookings {
		fmt.Printf("%-10d %-30s %-20s %-20s\n", 
			booking.ID, 
			booking.EventName, 
			booking.TimeTable.StartDate.Format("2006-01-02 15:04"), 
			booking.TimeTable.EndDate.Format("2006-01-02 15:04"))
	}
	fmt.Println("--------------------------------------------------------")
	
	util.PressEnterToContinue()
}

func handleGetBookingDetails(facade *controller.SpaceManagementControllerFacade) {
	fmt.Println("====== Get Booking Details ======")
	fmt.Print("Enter booking ID: ")
	bookingIDInput := util.GetCommandInput()
	bookingID, err := strconv.Atoi(bookingIDInput)
	if err != nil {
		fmt.Println("Invalid booking ID.")
		util.PressEnterToContinue()
		return
	}

	booking, err := facade.Booking.GetBookingDetails(uint(bookingID))
	if err != nil {
		fmt.Println("Error retrieving booking details:", err)
		util.PressEnterToContinue()
		return
	}

	fmt.Println("\nBooking Details:")
	fmt.Println("--------------------------------------------------------")
	fmt.Printf("Booking ID:     %d\n", booking.ID)
	fmt.Printf("Event Name:     %s\n", booking.EventName)
	fmt.Printf("User ID:        %d\n", booking.UserID)
	fmt.Printf("User Role:      %s\n", booking.UserRole)
	fmt.Printf("Room:           %s (ID: %d)\n", booking.TimeTable.Room.RoomName, booking.TimeTable.Room.ID)
	fmt.Printf("Room Type:      %s\n", booking.TimeTable.Room.RoomType)
	fmt.Printf("Building:       %s\n", booking.TimeTable.Room.Building)
	fmt.Printf("Floor:          %d\n", booking.TimeTable.Room.Floor)
	fmt.Printf("Start Time:     %s\n", booking.TimeTable.StartDate.Format("2006-01-02 15:04"))
	fmt.Printf("End Time:       %s\n", booking.TimeTable.EndDate.Format("2006-01-02 15:04"))
	fmt.Println("--------------------------------------------------------")
	
	util.PressEnterToContinue()
}

func handleResetAllBookings(facade *controller.SpaceManagementControllerFacade) {
	fmt.Println("====== Reset ALL Bookings ======")
	fmt.Println("WARNING: This will reset ALL bookings in the system.")
	fmt.Println("This action cannot be undone!")
	fmt.Print("Type 'CONFIRM' to proceed: ")
	confirm := util.GetCommandInput()
	if confirm != "CONFIRM" {
		fmt.Println("Operation canceled.")
		util.PressEnterToContinue()
		return
	}

	fmt.Print("Are you absolutely sure? Type 'YES' to confirm: ")
	secondConfirm := util.GetCommandInput()
	if secondConfirm != "YES" {
		fmt.Println("Operation canceled.")
		util.PressEnterToContinue()
		return
	}

	err := facade.Booking.ResetAllBookings()
	if err != nil {
		fmt.Println("Error resetting all bookings:", err)
	} else {
		fmt.Println("Successfully reset all bookings.")
	}
	util.PressEnterToContinue()
}

func handleCheckRoomAvailability(facade *controller.SpaceManagementControllerFacade) {
	fmt.Println("====== Check Room Availability ======")
	
	rooms, err := facade.Room.GetAll()
	if err != nil {
		fmt.Println("Error retrieving rooms:", err)
		util.PressEnterToContinue()
		return
	}

	fmt.Println("Available rooms:")
	for _, room := range *rooms {
		if !room.IsRoomOutOfService {
			fmt.Printf("ID: %d, Name: %s, Type: %s, Building: %s\n", 
				room.ID, room.RoomName, room.RoomType, room.Building)
		}
	}
	
	fmt.Print("\nEnter room ID: ")
	roomIDInput := util.GetCommandInput()
	roomID, err := strconv.Atoi(roomIDInput)
	if err != nil {
		fmt.Println("Invalid room ID.")
		util.PressEnterToContinue()
		return
	}

	startTime, err := getValidDateTimeInput("start")
	if err != nil {
		fmt.Println(err)
		util.PressEnterToContinue()
		return
	}

	endTime, err := getValidDateTimeInput("end")
	if err != nil {
		fmt.Println(err)
		util.PressEnterToContinue()
		return
	}

	if err := validateDateTimeRange(startTime, endTime); err != nil {
		fmt.Println(err)
		util.PressEnterToContinue()
		return
	}

	isAvailable, err := facade.Booking.CheckRoomAvailability(uint(roomID), startTime, endTime)
	if err != nil {
		fmt.Println("Error checking availability:", err)
	} else if isAvailable {
		fmt.Printf("\nRoom ID %d is AVAILABLE for the period:\n", roomID)
		fmt.Printf("From: %s\n", startTime.Format("2006-01-02 15:04"))
		fmt.Printf("To:   %s\n", endTime.Format("2006-01-02 15:04"))
	} else {
		fmt.Printf("\nRoom ID %d is NOT AVAILABLE for the period:\n", roomID)
		fmt.Printf("From: %s\n", startTime.Format("2006-01-02 15:04"))
		fmt.Printf("To:   %s\n", endTime.Format("2006-01-02 15:04"))
	}
	util.PressEnterToContinue()
}

func handleGetAvailableRooms(facade *controller.SpaceManagementControllerFacade) {
	fmt.Println("====== Get Available Rooms ======")

	startTime, err := getValidDateTimeInput("start")
	if err != nil {
		fmt.Println(err)
		util.PressEnterToContinue()
		return
	}

	endTime, err := getValidDateTimeInput("end")
	if err != nil {
		fmt.Println(err)
		util.PressEnterToContinue()
		return
	}

	if err := validateDateTimeRange(startTime, endTime); err != nil {
		fmt.Println(err)
		util.PressEnterToContinue()
		return
	}

	fmt.Print("Filter by room type? (y/n): ")
	filterByType := util.GetCommandInput()
	
	var roomType *model.RoomTypeEnum
	if filterByType == "y" || filterByType == "Y" {
		typeSelection, err := getRoomTypeSelection()
		if err != nil {
			fmt.Println("No room type filter will be applied.")
		} else {
			roomType = typeSelection
		}
	}
	
	fmt.Print("Filter by minimum capacity? (y/n): ")
	filterByCapacity := util.GetCommandInput()
	
	var capacity *int
	if filterByCapacity == "y" || filterByCapacity == "Y" {
		fmt.Print("Enter minimum capacity required: ")
		capacityStr := util.GetCommandInput()
		cap, err := strconv.Atoi(capacityStr)
		if err != nil {
			fmt.Println("Invalid capacity. No capacity filter will be applied.")
		} else {
			capacity = &cap
		}
	}

	rooms, err := facade.Booking.GetAvailableRooms(startTime, endTime, roomType, capacity)
	if err != nil {
		fmt.Println("Error retrieving available rooms:", err)
		util.PressEnterToContinue()
		return
	}

	fmt.Println("\nAvailable Rooms for the period:")
	fmt.Printf("From: %s\n", startTime.Format("2006-01-02 15:04"))
	fmt.Printf("To:   %s\n", endTime.Format("2006-01-02 15:04"))
	
	if len(rooms) == 0 {
		fmt.Println("\nNo rooms available for the specified criteria.")
	} else {
		fmt.Println("\n---------------------------------------------------------------------------------")
		fmt.Printf("%-5s %-20s %-15s %-15s %-10s %-10s\n", 
			"ID", "Room Name", "Type", "Building", "Floor", "Capacity")
		fmt.Println("---------------------------------------------------------------------------------")
		for _, room := range rooms {
			fmt.Printf("%-5d %-20s %-15s %-15s %-10d %-10d\n", 
				room.ID, room.RoomName, room.RoomType, room.Building, room.Floor, room.Capacity)
		}
		fmt.Println("---------------------------------------------------------------------------------")
	}
	
	util.PressEnterToContinue()
}

func handleResetTimeSlots(facade *controller.SpaceManagementControllerFacade) {
	fmt.Println("====== Reset Time Slots for a Room ======")
	
	rooms, err := facade.Room.GetAll()
	if err != nil {
		fmt.Println("Error retrieving rooms:", err)
		util.PressEnterToContinue()
		return
	}

	fmt.Println("Available rooms:")
	for _, room := range *rooms {
		fmt.Printf("ID: %d, Name: %s, Type: %s, Building: %s\n", 
			room.ID, room.RoomName, room.RoomType, room.Building)
	}
	
	fmt.Print("\nEnter room ID: ")
	roomIDInput := util.GetCommandInput()
	roomID, err := strconv.Atoi(roomIDInput)
	if err != nil {
		fmt.Println("Invalid room ID.")
		util.PressEnterToContinue()
		return
	}

	fmt.Println("WARNING: This will reset all time slots for this room.")
	fmt.Println("All bookings for this room will be made available again.")
	fmt.Print("Type 'CONFIRM' to proceed: ")
	confirm := util.GetCommandInput()
	if confirm != "CONFIRM" {
		fmt.Println("Operation canceled.")
		util.PressEnterToContinue()
		return
	}

	err = facade.Booking.ResetTimeSlots(uint(roomID))
	if err != nil {
		fmt.Println("Error resetting time slots:", err)
	} else {
		fmt.Printf("Time slots for room ID %d reset successfully.\n", roomID)
	}
	util.PressEnterToContinue()
}