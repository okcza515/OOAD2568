// MEP-1013
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

    // Get booking ID
    fmt.Print("Enter Booking ID to update: ")
    var idStr string
    fmt.Scanln(&idStr)
    id, err := strconv.ParseUint(strings.TrimSpace(idStr), 10, 32)
    if err != nil {
        fmt.Println("Invalid Booking ID")
        return err
    }

    // Retrieve existing booking
    booking, err := h.controller.RetrieveByID(uint(id))
    if err != nil {
        fmt.Println("Error retrieving booking:", err)
        return err
    }

    // Store original booking for comparison
    originalBooking := booking

    // Update Event Name
    fmt.Printf("Current Event Name: %s\n", booking.EventName)
    fmt.Print("Enter new Event Name (or press Enter to keep current): ")
    newEventName := util.GetCommandInput()
    if newEventName != "" {
        booking.EventName = newEventName
    }

    // Update User Role
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

    // Update Start Date/Time
    fmt.Printf("Current Start Date/Time: %s\n", booking.TimeTable.StartDate.Format("2006-01-02 15:04"))
    fmt.Print("Update start date/time? (y/n): ")
    var updateStart string
    fmt.Scanln(&updateStart)
    if strings.ToLower(strings.TrimSpace(updateStart)) == "y" {
        fmt.Print("Enter new Start Date (YYYY-MM-DD): ")
        startDateStr := util.GetCommandInput()
        fmt.Print("Enter new Start Time (HH:MM): ")
        startTimeStr := util.GetCommandInput()

        startDateTime := startDateStr + " " + startTimeStr
        newStartDate, err := time.Parse("2006-01-02 15:04", startDateTime)
        if err != nil {
            fmt.Println("Invalid date/time format. Please use YYYY-MM-DD for date and HH:MM for time")
            return err
        }
        booking.TimeTable.StartDate = newStartDate
    }

    // Update End Date/Time
    fmt.Printf("Current End Date/Time: %s\n", booking.TimeTable.EndDate.Format("2006-01-02 15:04"))
    fmt.Print("Update end date/time? (y/n): ")
    var updateEnd string
    fmt.Scanln(&updateEnd)
    if strings.ToLower(strings.TrimSpace(updateEnd)) == "y" {
        fmt.Print("Enter new End Date (YYYY-MM-DD): ")
        endDateStr := util.GetCommandInput()
        fmt.Print("Enter new End Time (HH:MM): ")
        endTimeStr := util.GetCommandInput()

        endDateTime := endDateStr + " " + endTimeStr
        newEndDate, err := time.Parse("2006-01-02 15:04", endDateTime)
        if err != nil {
            fmt.Println("Invalid date/time format. Please use YYYY-MM-DD for date and HH:MM for time")
            return err
        }
        booking.TimeTable.EndDate = newEndDate
    }

    // Validate date range
    if booking.TimeTable.EndDate.Before(booking.TimeTable.StartDate) {
        fmt.Println("Error: End date/time must be after start date/time")
        return fmt.Errorf("invalid date range")
    }

    // Display changes and confirm
    fmt.Println("\nChanges to be made:")
    fmt.Println("==================================================================")
    if booking.EventName != originalBooking.EventName {
        fmt.Printf("Event Name: %s -> %s\n", originalBooking.EventName, booking.EventName)
    }
    if booking.UserRole != originalBooking.UserRole {
        fmt.Printf("User Role: %s -> %s\n", originalBooking.UserRole, booking.UserRole)
    }
    if !booking.TimeTable.StartDate.Equal(originalBooking.TimeTable.StartDate) {
        fmt.Printf("Start Date/Time: %s -> %s\n",
            originalBooking.TimeTable.StartDate.Format("2006-01-02 15:04"),
            booking.TimeTable.StartDate.Format("2006-01-02 15:04"))
    }
    if !booking.TimeTable.EndDate.Equal(originalBooking.TimeTable.EndDate) {
        fmt.Printf("End Date/Time: %s -> %s\n",
            originalBooking.TimeTable.EndDate.Format("2006-01-02 15:04"),
            booking.TimeTable.EndDate.Format("2006-01-02 15:04"))
    }
    fmt.Println("==================================================================")

    fmt.Print("\nConfirm these changes? (y/n): ")
    var confirm string
    fmt.Scanln(&confirm)
    if strings.ToLower(strings.TrimSpace(confirm)) != "y" {
        fmt.Println("Update cancelled")
        return nil
    }

    err = h.controller.UpdateByID(&booking)
    if err != nil {
        fmt.Println("Error updating booking:", err)
        return err
    }

    fmt.Println("Booking updated successfully")
    util.PressEnterToContinue()
    return nil
}