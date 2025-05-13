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

type AddBookingHandlerStrategy struct {
    controller interface {
        Insert(dataContext model.Booking) error
        RetrieveByID(id uint, preloads ...string) (model.Booking, error)
        CheckRoomAvailability(roomID uint, startDate, endDate time.Time) (bool, error)
    }
}

func NewAddBookingHandlerStrategy(
    controller interface {
        Insert(dataContext model.Booking) error
        RetrieveByID(id uint, preloads ...string) (model.Booking, error)
        CheckRoomAvailability(roomID uint, startDate, endDate time.Time) (bool, error)
    }) *AddBookingHandlerStrategy {
    return &AddBookingHandlerStrategy{controller: controller}
}

func (h *AddBookingHandlerStrategy) Execute() error {
    fmt.Println("===== Create New Booking =====")
    
    var booking model.Booking
    booking.TimeTable = model.TimeTable{}
    
    fmt.Print("Enter Room ID: ")
    var roomIDStr string
    fmt.Scanln(&roomIDStr)
    roomID, err := strconv.ParseUint(strings.TrimSpace(roomIDStr), 10, 32)
    if err != nil {
        fmt.Println("Invalid Room ID")
        return err
    }
    booking.TimeTable.RoomID = uint(roomID)
    
    fmt.Print("Enter Start Date (YYYY-MM-DD): ")
    startDateStr := util.GetCommandInput()
    fmt.Print("Enter Start Time (HH:MM): ")
    startTimeStr := util.GetCommandInput()
    
    startDateTime := startDateStr + " " + startTimeStr
    startDate, err := time.Parse("2006-01-02 15:04", startDateTime)
    if err != nil {
        fmt.Println("Invalid date/time format. Please use YYYY-MM-DD for date and HH:MM for time")
        return err
    }
    booking.TimeTable.StartDate = startDate
    
    fmt.Print("Enter End Date (YYYY-MM-DD): ")
    endDateStr := util.GetCommandInput()
    fmt.Print("Enter End Time (HH:MM): ")
    endTimeStr := util.GetCommandInput()
    
    endDateTime := endDateStr + " " + endTimeStr
    endDate, err := time.Parse("2006-01-02 15:04", endDateTime)
    if err != nil {
        fmt.Println("Invalid date/time format. Please use YYYY-MM-DD for date and HH:MM for time")
        return err
    }
    booking.TimeTable.EndDate = endDate

    fmt.Print("Enter Booking Type (TEMPORARY/PERMANENT): ")
    bookingTypeStr := strings.ToUpper(util.GetCommandInput())
    switch bookingTypeStr {
    case "TEMPORARY":
        booking.TimeTable.BookingType = model.BOOKING_TEMPORARY
    case "PERMANENT":
        booking.TimeTable.BookingType = model.BOOKING_PERMANENT
    default:
        fmt.Println("Invalid booking type. Setting default to TEMPORARY")
        booking.TimeTable.BookingType = model.BOOKING_TEMPORARY
    }

    available, err := h.controller.CheckRoomAvailability(booking.TimeTable.RoomID, startDate, endDate)
    if err != nil {
        fmt.Println("Error checking room availability:", err)
        return err
    }
    if !available {
        fmt.Println("Room is not available for the selected time period")
        return fmt.Errorf("room not available")
    }

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
    
    fmt.Printf("\nBooking Details:\n")
    fmt.Printf("Room ID: %d\n", booking.TimeTable.RoomID)
    fmt.Printf("Start Time: %v\n", booking.TimeTable.StartDate.Format("2006-01-02 15:04"))
    fmt.Printf("End Time: %v\n", booking.TimeTable.EndDate.Format("2006-01-02 15:04"))
    fmt.Printf("Booking Type: %s\n", booking.TimeTable.BookingType)
    fmt.Printf("User ID: %d\n", booking.UserID)
    fmt.Printf("User Role: %s\n", booking.UserRole)
    fmt.Printf("Event Name: %s\n", booking.EventName)
    
    fmt.Print("\nConfirm booking? (y/n): ")
    var confirm string
    fmt.Scanln(&confirm)
    if strings.ToLower(strings.TrimSpace(confirm)) != "y" {
        fmt.Println("Booking cancelled")
        return nil
    }

    if err := booking.Validate(); err != nil {
        fmt.Println("Validation error:", err)
        util.PressEnterToContinue()
        return err
    }
    
    err = h.controller.Insert(booking)
    if err != nil {
        fmt.Println("Error creating booking:", err)
        return err
    }
    
    fmt.Println("Booking created successfully")
    util.PressEnterToContinue()
    return nil
}