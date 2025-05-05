package handler

import (
	"ModEd/asset/model"
	"ModEd/asset/util"
	"fmt"
	"strconv"
	"strings"
)

type AddBookingHandlerStrategy struct {
	controller interface {
		Insert(dataContext model.Booking) error
	}
}

func NewAddBookingHandlerStrategy(
	controller interface {
		Insert(dataContext model.Booking) error
	}) *AddBookingHandlerStrategy {
	return &AddBookingHandlerStrategy{controller: controller}
}

func (h *AddBookingHandlerStrategy) Execute() error {
	fmt.Println("===== Create New Booking =====")
	
	var booking model.Booking
	
	fmt.Print("Enter TimeTable ID: ")
	var timeTableIDStr string
	fmt.Scanln(&timeTableIDStr)
	timeTableID, err := strconv.ParseUint(strings.TrimSpace(timeTableIDStr), 10, 32)
	if err != nil {
		fmt.Println("Invalid TimeTable ID")
		return err
	}
	booking.TimeTableID = uint(timeTableID)
	
	fmt.Print("Enter User ID: ")
	var userIDStr string
	fmt.Scanln(&userIDStr)
	userID, err := strconv.ParseUint(strings.TrimSpace(userIDStr), 10, 32)
	if err != nil {
		fmt.Println("Invalid User ID")
		return err
	}
	booking.UserID = uint(userID)
	
	fmt.Print("Enter User Role (STUDENT/ADVISOR/ADMIN): ")
	var role string
	fmt.Scanln(&role)
	role = strings.ToUpper(strings.TrimSpace(role))
	switch role {
	case "STUDENT":
		booking.UserRole = model.ROLE_STUDENT
	case "ADVISOR":
		booking.UserRole = model.ROLE_ADVISOR
	case "ADMIN":
		booking.UserRole = model.ROLE_ADMIN
	default:
		fmt.Println("Invalid role. Setting default to STUDENT")
		booking.UserRole = model.ROLE_STUDENT
	}
	
	fmt.Print("Enter Event Name: ")
	booking.EventName = util.GetCommandInput()

	fmt.Println(booking)
	
	err = h.controller.Insert(booking)
	if err != nil {
		fmt.Println("Error creating booking:", err)
		return err
	}
	
	fmt.Println("Booking created successfully")
	util.PressEnterToContinue()
	return nil
}