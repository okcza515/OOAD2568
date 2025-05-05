package handler

import (
    "ModEd/asset/model"
    "ModEd/asset/util"
    "fmt"
    "strconv"
    "strings"
    "time"
)

type UpdateBookingHandlerStrategy struct {
    controller interface {
        UpdateByID(*model.Booking) error
        RetrieveByID(id uint, preloads ...string) (model.Booking, error)
    }
}

func NewUpdateBookingHandlerStrategy(controller interface {
    UpdateByID(*model.Booking) error
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
    
    bookingPtr := &booking
    
    fmt.Printf("Current Event Name: %s\n", bookingPtr.EventName)
    fmt.Print("Enter new Event Name (or press Enter to keep current): ")
    newEventName := util.GetCommandInput()
    if newEventName != "" {
        bookingPtr.EventName = newEventName
    }
    
    fmt.Printf("Current User Role: %s\n", bookingPtr.UserRole)
    fmt.Print("Enter new User Role (STUDENT/ADVISOR/ADMIN) (or press Enter to keep current): ")
    newRole := util.GetCommandInput()
    if newRole != "" {
        newRole = strings.ToUpper(strings.TrimSpace(newRole))
        switch newRole {
        case "STUDENT":
            bookingPtr.UserRole = model.ROLE_STUDENT
        case "ADVISOR":
            bookingPtr.UserRole = model.ROLE_ADVISOR
        case "ADMIN":
            bookingPtr.UserRole = model.ROLE_ADMIN
        default:
            fmt.Println("Invalid role. Keeping current role.")
        }
    }

    fmt.Printf("Current Start Date: %v\n", bookingPtr.TimeTable.StartDate.Format("2006-01-02 15:04"))
    fmt.Print("Enter new Start Date (YYYY-MM-DD) (or press Enter to keep current): ")
    newStartDateStr := util.GetCommandInput()
    
    if newStartDateStr != "" {
        fmt.Print("Enter new Start Time (HH:MM): ")
        newStartTimeStr := util.GetCommandInput()
        
        startDateTime := newStartDateStr + " " + newStartTimeStr
        newStartDate, err := time.Parse("2006-01-02 15:04", startDateTime)
        if err != nil {
            fmt.Println("Invalid date/time format. Please use YYYY-MM-DD for date and HH:MM for time")
            return err
        }
        bookingPtr.TimeTable.StartDate = newStartDate
    }
    
    fmt.Printf("Current End Date: %v\n", bookingPtr.TimeTable.EndDate.Format("2006-01-02 15:04"))
    fmt.Print("Enter new End Date (YYYY-MM-DD) (or press Enter to keep current): ")
    newEndDateStr := util.GetCommandInput()
    
    if newEndDateStr != "" {
        fmt.Print("Enter new End Time (HH:MM): ")
        newEndTimeStr := util.GetCommandInput()
        
        endDateTime := newEndDateStr + " " + newEndTimeStr
        newEndDate, err := time.Parse("2006-01-02 15:04", endDateTime)
        if err != nil {
            fmt.Println("Invalid date/time format. Please use YYYY-MM-DD for date and HH:MM for time")
            return err
        }
        bookingPtr.TimeTable.EndDate = newEndDate
    }

    // Validate date range
    if bookingPtr.TimeTable.EndDate.Before(bookingPtr.TimeTable.StartDate) {
        fmt.Println("Error: End date/time must be after start date/time")
        return fmt.Errorf("invalid date range")
    }

    // Show updated booking details
    fmt.Printf("\nUpdated Booking Details:\n")
    fmt.Printf("Event Name: %s\n", bookingPtr.EventName)
    fmt.Printf("User Role: %s\n", bookingPtr.UserRole)
    fmt.Printf("Start Date: %v\n", bookingPtr.TimeTable.StartDate.Format("2006-01-02 15:04"))
    fmt.Printf("End Date: %v\n", bookingPtr.TimeTable.EndDate.Format("2006-01-02 15:04"))

    fmt.Print("\nConfirm update? (y/n): ")
    var confirm string
    fmt.Scanln(&confirm)
    if strings.ToLower(strings.TrimSpace(confirm)) != "y" {
        fmt.Println("Update cancelled")
        return nil
    }
    
    err = h.controller.UpdateByID(bookingPtr)
    if err != nil {
        fmt.Println("Error updating booking:", err)
        return err
    }
    
    fmt.Println("Booking updated successfully")
    util.PressEnterToContinue()
    return nil
}