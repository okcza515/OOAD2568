// MEP-1013
package handler

import (
	controller "ModEd/asset/controller/spacemanagement"
	"ModEd/asset/util"
	"fmt"
	"strconv"
	"time"
)

func printPermanentBookingMenuOption() {
	fmt.Println("===== Permanent Booking Management =====")
	fmt.Println("1. Create Weekly Class Schedule")
	fmt.Println("Type 'back' to return to previous menu")
	fmt.Println("========================================")
}

func PermanentBookingHandler(facade *controller.PermanentBookingController) {
	inputBuffer := ""

	for inputBuffer != "back" {
		util.ClearScreen()
		util.PrintSpaceManagementBanner()
		printPermanentBookingMenuOption()
		inputBuffer = util.GetCommandInput()

		switch inputBuffer {
		case "1":
			fmt.Println("Create Weekly Class Schedule")

			fmt.Println("Please enter the start date (YYYY-MM-DD HH:MM):")
			startStr := util.GetCommandInput()
			startDate, errStart := time.Parse("2006-01-02 15:04", startStr)

			fmt.Println("Please enter the end date (YYYY-MM-DD HH:MM):")
			endStr := util.GetCommandInput()
			endDate, errEnd := time.Parse("2006-01-02 15:04", endStr)

			if errStart != nil || errEnd != nil {
				fmt.Println("Invalid date format.")
				util.PressEnterToContinue()
				break
			}

			roomID := readUintInput("Please enter Room ID:")
			courseID := readUintInput("Please enter Course ID:")
			classID := readUintInput("Please enter Class ID:")
			facultyID := readUintInput("Please enter Faculty ID:")
			departmentID := readUintInput("Please enter Department ID:")
			programTypeID := readUintInput("Please enter Program Type ID:")

			err := facade.CreateWeeklySchedule(
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
				fmt.Println("Weekly class schedule created successfully!")
			}
			util.PressEnterToContinue()
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

func readUintInput(prompt string) uint {
	fmt.Println(prompt)
	input, err := strconv.Atoi(util.GetCommandInput())
	if err != nil {
		fmt.Println("Invalid input. Using default value 1.")
		return 1
	}
	return uint(input)
}
