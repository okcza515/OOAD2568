package handler

import (
	"ModEd/asset/util"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type CheckRoomAvailabilityHandlerStrategy struct {
	controller interface {
		CheckRoomAvailability(roomID uint, startDate, endDate time.Time) (bool, error)
	}
}

func NewCheckRoomAvailabilityHandlerStrategy(controller interface {
	CheckRoomAvailability(roomID uint, startDate, endDate time.Time) (bool, error)
}) *CheckRoomAvailabilityHandlerStrategy {
	return &CheckRoomAvailabilityHandlerStrategy{controller: controller}
}

func (h *CheckRoomAvailabilityHandlerStrategy) Execute() error {
	fmt.Println("===== Check Room Availability =====")

	fmt.Print("Enter Room ID: ")
	var roomIDStr string
	fmt.Scanln(&roomIDStr)
	roomID, err := strconv.ParseUint(strings.TrimSpace(roomIDStr), 10, 32)
	if err != nil {
		fmt.Println("Invalid Room ID")
		return err
	}
	
	fmt.Print("Enter Start Date (YYYY-MM-DD): ")
	startDateStr := strings.TrimSpace(util.GetCommandInput())
	fmt.Print("Enter Start Time (HH:MM): ")
	startTimeStr := strings.TrimSpace(util.GetCommandInput())
	
	startDateTime := startDateStr + " " + startTimeStr
	startDate, err := time.Parse("2006-01-02 15:04", startDateTime)
	if err != nil {
		fmt.Println("Invalid date/time format. Please use YYYY-MM-DD for date and HH:MM for time")
		return err
	}
	
	fmt.Print("Enter End Date (YYYY-MM-DD): ")
	endDateStr := strings.TrimSpace(util.GetCommandInput())
	fmt.Print("Enter End Time (HH:MM): ")
	endTimeStr := strings.TrimSpace(util.GetCommandInput())
	
	endDateTime := endDateStr + " " + endTimeStr
	endDate, err := time.Parse("2006-01-02 15:04", endDateTime)
	if err != nil {
		fmt.Println("Invalid date/time format. Please use YYYY-MM-DD for date and HH:MM for time")
		return err
	}
	
	available, err := h.controller.CheckRoomAvailability(uint(roomID), startDate, endDate)
	if err != nil {
		fmt.Println("Error checking availability:", err)
		return err
	}
	
	fmt.Println("==================================================================")
	fmt.Printf(" Room #%-5d                                                 \n", roomID)
	fmt.Printf(" Period: %-15s to %-15s               \n", 
		startDate.Format("2006-01-02 15:04"), 
		endDate.Format("2006-01-02 15:04"))
	fmt.Println("------------------------------------------------------------------")
	if available {
		fmt.Println("              ROOM IS AVAILABLE FOR THIS PERIOD                ")
	} else {
		fmt.Println("              ROOM IS NOT AVAILABLE FOR THIS PERIOD            ")
	}
	fmt.Println("==================================================================")
	
	util.PressEnterToContinue()
	return nil
}