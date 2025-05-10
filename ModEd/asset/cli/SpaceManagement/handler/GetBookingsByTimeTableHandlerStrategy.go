package handler

import (
	"ModEd/asset/model"
	"ModEd/asset/util"
	"fmt"
	"strconv"
	"strings"
)

type GetBookingsByTimeTableHandler struct {
	controller interface {
		GetBookingsByTimeTable(timeTableID uint) ([]model.Booking, error)
	}
}

func NewGetBookingsByTimeTableHandlerStrategy(controller interface {
	GetBookingsByTimeTable(timeTableID uint) ([]model.Booking, error)
}) *GetBookingsByTimeTableHandler {
	return &GetBookingsByTimeTableHandler{controller: controller}
}

func (h *GetBookingsByTimeTableHandler) Execute() error {
	fmt.Println("===== Bookings By TimeTable ID =====")

	fmt.Print("Enter TimeTable ID: ")
	var timeTableIDStr string
	fmt.Scanln(&timeTableIDStr)
	timeTableID, err := strconv.ParseUint(strings.TrimSpace(timeTableIDStr), 10, 32)
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
		fmt.Printf("Found %d booking(s) for TimeTable ID %d:\n\n", len(bookings), timeTableID)
		fmt.Println("==========================================================")
		fmt.Println(" No | Booking ID | User ID |  Role   |      Event Name    ")
		fmt.Println("----------------------------------------------------------")
		for i, booking := range bookings {
			fmt.Printf(" %2d | %-9d | %-7d | %-7s | %-21s \n", 
				i+1, booking.ID, booking.UserID, booking.UserRole, truncateString(booking.EventName, 21))
		}
		fmt.Println("==========================================================")
	}
	
	util.PressEnterToContinue()
	return nil
}

func truncateString(str string, length int) string {
	if len(str) <= length {
		return str
	}
	return str[:length-3] + "..."
}