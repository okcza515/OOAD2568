// MEP-1013
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
		case "back":
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
			util.PressEnterToContinue()
		}
	}
}

func handleBookRoom(facade *controller.SpaceManagementControllerFacade) {
	fmt.Println("Book a Room")

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
		fmt.Println("Invalid choice. Please select a valid room type.")
		util.PressEnterToContinue()
		return
	}

	fmt.Println("You have selected:", roomType.TypeRoomString())

	fmt.Print("Enter room ID: ")
	roomIDInput := util.GetCommandInput()
	roomID, err := strconv.Atoi(roomIDInput)
	if err != nil {
		fmt.Println("Invalid room ID.")
		util.PressEnterToContinue()
		return
	}

	var startTime, endTime time.Time

	fmt.Println("For start time:")
	fmt.Print("Enter date (YYYY-MM-DD): ")
	startDateStr := util.GetCommandInput()
	fmt.Print("Enter time (HH:MM): ")
	startTimeStr := util.GetCommandInput()

	startFullStr := startDateStr + " " + startTimeStr
	startTime, err = time.Parse("2006-01-02 15:04", startFullStr)
	if err != nil {
		fmt.Println("Invalid date or time format.")
		util.PressEnterToContinue()
		return
	}

	fmt.Println("For end time:")
	fmt.Print("Enter date (YYYY-MM-DD): ")
	endDateStr := util.GetCommandInput()
	fmt.Print("Enter time (HH:MM): ")
	endTimeStr := util.GetCommandInput()

	endFullStr := endDateStr + " " + endTimeStr
	endTime, err = time.Parse("2006-01-02 15:04", endFullStr)
	if err != nil {
		fmt.Println("Invalid date or time format.")
		util.PressEnterToContinue()
		return
	}

	if endTime.Before(startTime) || endTime.Equal(startTime) {
		fmt.Println("End time must be after start time.")
		util.PressEnterToContinue()
		return
	}

	fmt.Print("Enter your role (Student/Advisor/Admin): ")
	roleStr := util.GetCommandInput()

	var role model.Role
	switch roleStr {
	case "Student":
		role = model.ROLE_STUDENT
	case "Advisor":
		role = model.ROLE_ADVISOR
	case "Admin":
		role = model.ROLE_ADMIN
	default:
		fmt.Println("Invalid role. Please enter a valid role (Student/Advisor/Admin).")
		util.PressEnterToContinue()
		return
	}

	fmt.Print("Enter event name: ")
	eventName := util.GetCommandInput()

	fmt.Println("Attempting to book room...")
	fmt.Printf("Room ID: %d, User ID: %d, Role: %s, Event: %s\n", roomID, 1, roleStr, eventName)
	fmt.Printf("Time: %s to %s\n", startTime.Format("2006-01-02 15:04"), endTime.Format("2006-01-02 15:04"))

	time.Sleep(2000 * time.Millisecond)

	booking, err := facade.Booking.BookRoom(uint(roomID), uint(1), role, eventName, startTime, endTime)

	fmt.Println("\n----- BOOKING RESULT -----")
	if err != nil {
		fmt.Println("Booking failed:", err)
		fmt.Println("Please check if the room exists and is available for the specified time.")
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

	fmt.Println("\nPress Enter to continue...")
	util.PressEnterToContinue()
}

func handleCancelBooking(facade *controller.SpaceManagementControllerFacade) {
	fmt.Println("Cancel Booking")
	fmt.Print("Enter booking ID: ")
	bookingIDInput := util.GetCommandInput()
	bookingID, err := strconv.Atoi(bookingIDInput)
	if err != nil {
		fmt.Println("Invalid booking ID.")
		util.PressEnterToContinue()
		return
	}

	err = facade.Booking.CancelBooking(uint(bookingID))
	if err != nil {
		fmt.Println("Cancel failed:", err)
	} else {
		fmt.Println("Booking canceled.")
	}
	util.PressEnterToContinue()
}

func handleUpdateBooking(facade *controller.SpaceManagementControllerFacade) {
	fmt.Println("Update a Booking")

	fmt.Print("Enter booking ID: ")
	bookingIDInput := util.GetCommandInput()
	bookingID, err := strconv.Atoi(bookingIDInput)
	if err != nil {
		fmt.Println("Invalid booking ID.")
		util.PressEnterToContinue()
		return
	}

	var startTime, endTime time.Time

	fmt.Println("For new start time:")
	fmt.Print("Enter date (YYYY-MM-DD): ")
	startDateStr := util.GetCommandInput()
	fmt.Print("Enter time (HH:MM): ")
	startTimeStr := util.GetCommandInput()

	startFullStr := startDateStr + " " + startTimeStr
	startTime, err = time.Parse("2006-01-02 15:04", startFullStr)
	if err != nil {
		fmt.Println("Invalid date or time format.")
		util.PressEnterToContinue()
		return
	}

	fmt.Println("For new end time:")
	fmt.Print("Enter date (YYYY-MM-DD): ")
	endDateStr := util.GetCommandInput()
	fmt.Print("Enter time (HH:MM): ")
	endTimeStr := util.GetCommandInput()

	endFullStr := endDateStr + " " + endTimeStr
	endTime, err = time.Parse("2006-01-02 15:04", endFullStr)
	if err != nil {
		fmt.Println("Invalid date or time format.")
		util.PressEnterToContinue()
		return
	}

	if endTime.Before(startTime) || endTime.Equal(startTime) {
		fmt.Println("End time must be after start time.")
		util.PressEnterToContinue()
		return
	}

	fmt.Print("Enter new event name: ")
	eventName := util.GetCommandInput()

	eventNamePtr := &eventName
	startTimePtr := &startTime
	endTimePtr := &endTime

	err = facade.Booking.UpdateBooking(uint(bookingID), eventNamePtr, startTimePtr, endTimePtr)
	if err != nil {
		fmt.Println("Failed to update booking:", err)
	} else {
		fmt.Println("Booking updated successfully.")
	}
	util.PressEnterToContinue()
}

func handleGetRoomBookings(facade *controller.SpaceManagementControllerFacade) {
	fmt.Println("Get Room Bookings")
	fmt.Print("Enter room ID: ")
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
	} else {
		fmt.Println("Bookings for Room ID", roomID)
		for _, booking := range bookings {
			fmt.Printf("Booking ID: %d, Event: %s, Start: %s, End: %s\n", booking.ID, booking.EventName, booking.TimeTable.StartDate, booking.TimeTable.EndDate)
		}
	}
	util.PressEnterToContinue()
}

func handleGetBookingDetails(facade *controller.SpaceManagementControllerFacade) {
	fmt.Println("Get Booking Details")
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
	} else {
		fmt.Printf("Booking ID: %d\nEvent: %s\nRoom: %s\nStart: %s\nEnd: %s\n", booking.ID, booking.EventName, booking.TimeTable.Room.RoomType, booking.TimeTable.StartDate, booking.TimeTable.EndDate)
	}
	util.PressEnterToContinue()
}

func handleResetAllBookings(facade *controller.SpaceManagementControllerFacade) {
	fmt.Println("Are you sure you want to reset ALL bookings? (y/n): ")
	confirm := util.GetCommandInput()
	if confirm != "y" && confirm != "Y" {
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
	fmt.Println("Check Room Availability")
	fmt.Print("Enter room ID: ")
	roomIDInput := util.GetCommandInput()
	roomID, err := strconv.Atoi(roomIDInput)
	if err != nil {
		fmt.Println("Invalid room ID.")
		util.PressEnterToContinue()
		return
	}

	var startTime, endTime time.Time

	fmt.Println("For start time:")
	fmt.Print("Enter date (YYYY-MM-DD): ")
	startDateStr := util.GetCommandInput()
	fmt.Print("Enter time (HH:MM): ")
	startTimeStr := util.GetCommandInput()

	startFullStr := startDateStr + " " + startTimeStr
	startTime, err = time.Parse("2006-01-02 15:04", startFullStr)
	if err != nil {
		fmt.Println("Invalid date or time format.")
		util.PressEnterToContinue()
		return
	}

	fmt.Println("For end time:")
	fmt.Print("Enter date (YYYY-MM-DD): ")
	endDateStr := util.GetCommandInput()
	fmt.Print("Enter time (HH:MM): ")
	endTimeStr := util.GetCommandInput()

	endFullStr := endDateStr + " " + endTimeStr
	endTime, err = time.Parse("2006-01-02 15:04", endFullStr)
	if err != nil {
		fmt.Println("Invalid date or time format.")
		util.PressEnterToContinue()
		return
	}

	isAvailable, err := facade.Booking.CheckRoomAvailability(uint(roomID), startTime, endTime)
	if err != nil {
		fmt.Println("Error checking availability:", err)
	} else if isAvailable {
		fmt.Println("Room is available!")
	} else {
		fmt.Println("Room is not available.")
	}
	util.PressEnterToContinue()
}

func handleGetAvailableRooms(facade *controller.SpaceManagementControllerFacade) {
	fmt.Println("Get Available Rooms")

	var startTime, endTime time.Time

	fmt.Println("For start time:")
	fmt.Print("Enter date (YYYY-MM-DD): ")
	startDateStr := util.GetCommandInput()
	fmt.Print("Enter time (HH:MM): ")
	startTimeStr := util.GetCommandInput()

	startFullStr := startDateStr + " " + startTimeStr
	startTime, err1 := time.Parse("2006-01-02 15:04", startFullStr)
	if err1 != nil {
		fmt.Println("Invalid date or time format.")
		util.PressEnterToContinue()
		return
	}

	fmt.Println("For end time:")
	fmt.Print("Enter date (YYYY-MM-DD): ")
	endDateStr := util.GetCommandInput()
	fmt.Print("Enter time (HH:MM): ")
	endTimeStr := util.GetCommandInput()

	endFullStr := endDateStr + " " + endTimeStr
	endTime, err2 := time.Parse("2006-01-02 15:04", endFullStr)
	if err2 != nil {
		fmt.Println("Invalid date or time format.")
		util.PressEnterToContinue()
		return
	}

	rooms, err := facade.Booking.GetAvailableRooms(startTime, endTime, nil, nil)
	if err != nil {
		fmt.Println("Error retrieving available rooms:", err)
	} else {
		fmt.Println("Available Rooms:")
		for _, room := range rooms {
			fmt.Println("Room ID:", room.ID, "Room:", room.RoomType)
		}
	}
	util.PressEnterToContinue()
}

func handleResetTimeSlots(facade *controller.SpaceManagementControllerFacade) {
	fmt.Println("Reset Time Slots for a Room")
	fmt.Print("Enter room ID: ")
	roomIDInput := util.GetCommandInput()
	roomID, err := strconv.Atoi(roomIDInput)
	if err != nil {
		fmt.Println("Invalid room ID.")
		util.PressEnterToContinue()
		return
	}

	fmt.Println("Are you sure you want to reset all time slots for this room? (y/n): ")
	confirm := util.GetCommandInput()
	if confirm != "y" && confirm != "Y" {
		fmt.Println("Operation canceled.")
		util.PressEnterToContinue()
		return
	}

	err = facade.Booking.ResetTimeSlots(uint(roomID))
	if err != nil {
		fmt.Println("Error resetting time slots:", err)
	} else {
		fmt.Println("Time slots reset successfully.")
	}
	util.PressEnterToContinue()
}
