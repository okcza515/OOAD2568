// MEP-1013
package handler

import (
	controller "ModEd/asset/controller"
	"ModEd/asset/util"
	"fmt"
	"strconv"
	"time"
)

func printPermanentBookingMenuOption() {
	fmt.Println("===== Permanent Booking Management =====")
	fmt.Println("1. Create Weekly Class Schedule")
	fmt.Println("2. View All Permanent Bookings")
	fmt.Println("3. Update Permanent Booking")
	fmt.Println("4. Check Room Service Status")
	fmt.Println("Type 'back' to return to previous menu")
	fmt.Println("========================================")
}

func PermanentBookingHandler(facade *controller.SpaceManagementControllerFacade) {
	inputBuffer := ""

	for inputBuffer != "back" {
		util.ClearScreen()
		util.PrintSpaceManagementBanner()
		printPermanentBookingMenuOption()
		inputBuffer = util.GetCommandInput()

		switch inputBuffer {
		case "1":
			handleCreateWeeklySchedule(facade)
		case "2":
			handleViewAllPermanentBookings(facade)
		case "3":
			handleUpdatePermanentBooking(facade)
		case "4":
			handleCheckRoomServiceStatus(facade)
		case "back":
			return
		default:
			fmt.Println("Invalid Command")
			util.PressEnterToContinue()
		}

		util.ClearScreen()
	}

	util.ClearScreen()
}

func readTimeInput() (time.Time, error) {
	fmt.Print("Enter date (YYYY-MM-DD): ")
	dateStr := util.GetCommandInput()
	fmt.Print("Enter time (HH:MM): ")
	timeStr := util.GetCommandInput()

	dateTimeStr := dateStr + " " + timeStr
	return time.Parse("2006-01-02 15:04", dateTimeStr)
}

func handleCreateWeeklySchedule(facade *controller.SpaceManagementControllerFacade) {
	fmt.Println("Create Weekly Class Schedule")
	fmt.Println("===========================")

	roomID := readUintInput("Please enter Room ID:")

	inService, err := facade.PermanentSchedule.CheckRoomIsInService(roomID)
	if err != nil {
		fmt.Println(err)
		util.PressEnterToContinue()
		return
	}

	if !inService {
		fmt.Printf("Room with ID %d is out of service and cannot be booked.\n", roomID)
		util.PressEnterToContinue()
		return
	}

	fmt.Println("Room is available for booking. Please continue.")

	fmt.Println("For start time:")
	startDate, errStart := readTimeInput()

	fmt.Println("For end time:")
	endDate, errEnd := readTimeInput()

	if errStart != nil || errEnd != nil {
		fmt.Println("Invalid date/time format.")
		fmt.Println("Please use YYYY-MM-DD for date and HH:MM for time.")
		util.PressEnterToContinue()
		return
	}

	courseID := readUintInput("Please enter Course ID:")
	classID := readUintInput("Please enter Class ID:")
	facultyID := readUintInput("Please enter Faculty ID:")
	departmentID := readUintInput("Please enter Department ID:")
	programTypeID := readUintInput("Please enter Program Type ID:")

	err = facade.PermanentSchedule.CreateWeeklySchedule(
		startDate,
		endDate,
		roomID,
		courseID,
		classID,
		facultyID,
		departmentID,
		programTypeID,
	)

	if err != nil {
		fmt.Println("Failed to create weekly schedule:", err)
	} else {
		scheduleIDs := facade.PermanentSchedule.GetLastCreatedScheduleIDs()
		fmt.Println("Weekly class schedule created successfully!")
		fmt.Printf("Created %d schedules with IDs: %v\n", len(scheduleIDs), scheduleIDs)
	}
	util.PressEnterToContinue()
}

func handleViewAllPermanentBookings(facade *controller.SpaceManagementControllerFacade) {
	fmt.Println("All Permanent Bookings")
	fmt.Println("========================")

	bookings, err := facade.PermanentSchedule.GetAllPermanentBookings()
	if err != nil {
		fmt.Println("Failed to retrieve bookings:", err)
		util.PressEnterToContinue()
		return
	}

	if len(bookings) == 0 {
		fmt.Println("No permanent bookings found.")
	} else {
		fmt.Printf("Found %d permanent bookings:\n\n", len(bookings))
		for i, booking := range bookings {
			fmt.Printf("%d. Schedule ID: %d\n", i+1, booking.ID)
			fmt.Printf("   Course ID: %d, Class ID: %d\n", booking.CourseId, booking.ClassId)
			fmt.Printf("   Faculty ID: %d, Department ID: %d, Program Type ID: %d\n",
				booking.FacultyID, booking.DepartmentID, booking.ProgramtypeID)

			if booking.TimeTable.ID > 0 {
				fmt.Printf("   Time: %s - %s\n",
					booking.TimeTable.StartDate.Format("2006-01-02 15:04"),
					booking.TimeTable.EndDate.Format("2006-01-02 15:04"))
				fmt.Printf("   Room ID: %d\n", booking.TimeTable.RoomID)
			}
			fmt.Println("------------------------")
		}
	}

	util.PressEnterToContinue()
}

func handleUpdatePermanentBooking(facade *controller.SpaceManagementControllerFacade) {
	fmt.Println("Update Permanent Booking")
	fmt.Println("========================")

	scheduleID := readUintInput("Please enter Schedule ID to update:")
	roomID := readUintInput("Please enter new Room ID:")

	inService, err := facade.PermanentSchedule.CheckRoomIsInService(roomID)
	if err != nil {
		fmt.Println(err)
		util.PressEnterToContinue()
		return
	}

	if !inService {
		fmt.Printf("Room with ID %d is out of service and cannot be booked.\n", roomID)
		util.PressEnterToContinue()
		return
	}

	fmt.Println("Room is available for booking. Please continue.")

	fmt.Println("For new start time:")
	startDate, errStart := readTimeInput()

	fmt.Println("For new end time:")
	endDate, errEnd := readTimeInput()

	if errStart != nil || errEnd != nil {
		fmt.Println("Invalid date/time format.")
		fmt.Println("Please use YYYY-MM-DD for date and HH:MM for time.")
		util.PressEnterToContinue()
		return
	}

	courseID := readUintInput("Please enter new Course ID:")
	classID := readUintInput("Please enter new Class ID:")
	facultyID := readUintInput("Please enter new Faculty ID:")
	departmentID := readUintInput("Please enter new Department ID:")
	programTypeID := readUintInput("Please enter new Program Type ID:")

	err = facade.PermanentSchedule.UpdatePermanentBooking(
		startDate,
		endDate,
		roomID,
		courseID,
		classID,
		facultyID,
		departmentID,
		programTypeID,
		scheduleID,
	)

	if err != nil {
		fmt.Println("Failed to update permanent booking:", err)
	} else {
		fmt.Printf("Successfully updated permanent booking (ID: %d)\n", scheduleID)
	}
	util.PressEnterToContinue()
}

func handleCheckRoomServiceStatus(facade *controller.SpaceManagementControllerFacade) {
	fmt.Println("Check Room Service Status")
	fmt.Println("========================")

	roomID := readUintInput("Please enter Room ID to check:")

	inService, err := facade.PermanentSchedule.CheckRoomIsInService(roomID)

	if err != nil {
		fmt.Println(err)
	} else if inService {
		fmt.Printf("room with ID %d is in service and available for booking.\n", roomID)
	} else {
		fmt.Printf("room with ID %d is OUT OF SERVICE and cannot be booked.\n", roomID)
	}

	util.PressEnterToContinue()
}

func readUintInput(prompt string) uint {
	fmt.Println(prompt)
	input, err := strconv.Atoi(util.GetCommandInput())
	if err != nil {
		fmt.Println("Invalid input. Using default value 1.")
		return 1
	}
	return uint(input)
}
