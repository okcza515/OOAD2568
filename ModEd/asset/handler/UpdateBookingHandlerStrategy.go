package handler

import (
	"ModEd/asset/model"
	"ModEd/asset/util"
	"fmt"
	"strconv"
	"strings"
)

type UpdateBookingHandlerStrategy struct {
	controller interface {
		UpdateByID(data model.Booking) error
		RetrieveByID(id uint, preloads ...string) (model.Booking, error)
	}
}

func NewUpdateBookingHandlerStrategy(controller interface {
	UpdateByID(data model.Booking) error
	RetrieveByID(id uint, preloads ...string) (model.Booking, error)
}) *UpdateBookingHandlerStrategy {
	return &UpdateBookingHandlerStrategy{controller: controller}
}

func (h *UpdateBookingHandlerStrategy) Execute() error {
	fmt.Println("===== Update Booking =====")

	
	fmt.Print("Enter Booking ID to update: ")
	var idStr string
	fmt.Scanln(&idStr)
	id, err := strconv.ParseUint(strings.TrimSpace(idStr), 10, 32)
	if err != nil {
		fmt.Println("Invalid Booking ID")
		return err
	}
	
	booking, err := h.controller.RetrieveByID(uint(id))
	if err != nil {
		fmt.Println("Error retrieving booking:", err)
		return err
	}
	
	fmt.Printf("Current Event Name: %s\n", booking.EventName)
	fmt.Print("Enter new Event Name (or press Enter to keep current): ")
	newEventName := util.GetCommandInput()
	if newEventName != "" {
		booking.EventName = newEventName
	}
	
	fmt.Printf("Current User Role: %s\n", booking.UserRole)
	fmt.Print("Enter new User Role (STUDENT/ADVISOR/ADMIN) (or press Enter to keep current): ")
	newRole := util.GetCommandInput()
	if newRole != "" {
		newRole = strings.ToUpper(strings.TrimSpace(newRole))
		switch newRole {
		case "STUDENT":
			booking.UserRole = model.ROLE_STUDENT
		case "ADVISOR":
			booking.UserRole = model.ROLE_ADVISOR
		case "ADMIN":
			booking.UserRole = model.ROLE_ADMIN
		default:
			fmt.Println("Invalid role. Keeping current role.")
		}
	}
	
	err = h.controller.UpdateByID(booking)
	if err != nil {
		fmt.Println("Error updating booking:", err)
		return err
	}
	
	fmt.Println("Booking updated successfully")
	util.PressEnterToContinue()
	return nil
}