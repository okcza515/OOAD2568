// MEP-1013
package spacemanagement

import (
	model "ModEd/asset/model/spacemanagement"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type BookingController struct {
	db *gorm.DB
}

func NewBookingController(db *gorm.DB) *BookingController {
	return &BookingController{
		db: db,
	}
}

func (controller *BookingController) CheckRoomAvailability(roomID uint, startDate, endDate time.Time) (bool, error) {
	var room model.Room
	if err := controller.db.First(&room, roomID).Error; err != nil {
		return false, fmt.Errorf("unable to find room with ID %d: %w", roomID, err)
	}

	if room.IsRoomOutOfService {
		return false, fmt.Errorf("room with ID %d is out of service", roomID)
	}

	var bookingCount int64
	if err := controller.db.Model(&model.TimeTable{}).
		Where("room_id = ? AND is_available = ? AND ((start_date <= ? AND end_date >= ?) OR (start_date <= ? AND end_date >= ?) OR (start_date >= ? AND end_date <= ?))",
			roomID, false, startDate, startDate, endDate, endDate, startDate, endDate).
		Count(&bookingCount).Error; err != nil {
		return false, fmt.Errorf("error checking for existing bookings: %w", err)
	}

	if bookingCount > 0 {
		return false, nil
	}

	return true, nil
}

func (controller *BookingController) ResetTimeSlots(roomID uint) error {
	if err := controller.db.Model(&model.TimeTable{}).
		Where("room_id = ? AND booking_id IS NOT NULL", roomID).
		Update("is_available", true).Error; err != nil {
		return fmt.Errorf("unable to reset time slots for room %d: %w", roomID, err)
	}

	fmt.Printf("Successfully reset time slots for room %d\n", roomID)
	return nil
}

func (controller *BookingController) ResetAllBookings() error {
	if err := controller.db.Model(&model.TimeTable{}).
		Where("booking_id IS NOT NULL").
		Update("is_available", true).Error; err != nil {
		return fmt.Errorf("unable to reset all bookings: %w", err)
	}

	fmt.Println("Successfully reset all bookings")
	return nil
}

func (controller *BookingController) BookRoom(roomID uint, userID uint, userRole model.Role, eventName string, startDate, endDate time.Time) (*model.Booking, error) {
	isAvailable, err := controller.CheckRoomAvailability(roomID, startDate, endDate)
	if err != nil {
		return nil, err
	}
	if !isAvailable {
		return nil, fmt.Errorf("room with ID %d is not available during the specified time period", roomID)
	}

	tx := controller.db.Begin()

	timeTable := model.TimeTable{
		StartDate:   startDate,
		EndDate:     endDate,
		RoomID:      roomID,
		IsAvailable: false,
	}

	if err := tx.Create(&timeTable).Error; err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("unable to create time table entry: %w", err)
	}

	booking := model.Booking{
		TimeTableID: timeTable.ID,
		UserID:      userID,
		UserRole:    userRole,
		EventName:   eventName,
	}

	if err := tx.Create(&booking).Error; err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("unable to create booking: %w", err)
	}

	bookingID := booking.ID
	timeTable.BookingID = &bookingID
	if err := tx.Save(&timeTable).Error; err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("unable to update time table with booking ID: %w", err)
	}

	if err := tx.Commit().Error; err != nil {
		return nil, fmt.Errorf("transaction failed: %w", err)
	}

	fmt.Printf("Room %d successfully booked for event '%s' from %v to %v\n", roomID, eventName, startDate, endDate)
	return &booking, nil
}

func (controller *BookingController) CancelBooking(bookingID uint) error {
	var booking model.Booking
	if err := controller.db.First(&booking, bookingID).Error; err != nil {
		return fmt.Errorf("unable to find booking with ID %d: %w", bookingID, err)
	}

	var timeTable model.TimeTable
	if err := controller.db.First(&timeTable, booking.TimeTableID).Error; err != nil {
		return fmt.Errorf("unable to find time table for booking %d: %w", bookingID, err)
	}

	tx := controller.db.Begin()

	timeTable.IsAvailable = true
	if err := tx.Save(&timeTable).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("unable to update time table: %w", err)
	}

	if err := tx.Delete(&booking).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("unable to cancel booking: %w", err)
	}

	if err := tx.Commit().Error; err != nil {
		return fmt.Errorf("transaction failed: %w", err)
	}

	fmt.Printf("Successfully canceled booking with ID %d\n", bookingID)
	return nil
}

func (controller *BookingController) UpdateBooking(bookingID uint, eventName *string, newStartDate, newEndDate *time.Time) error {
	var booking model.Booking
	if err := controller.db.First(&booking, bookingID).Error; err != nil {
		return fmt.Errorf("unable to find booking with ID %d: %w", bookingID, err)
	}

	var timeTable model.TimeTable
	if err := controller.db.First(&timeTable, booking.TimeTableID).Error; err != nil {
		return fmt.Errorf("unable to find time table for booking %d: %w", bookingID, err)
	}

	if (newStartDate != nil && !timeTable.StartDate.Equal(*newStartDate)) || (newEndDate != nil && !timeTable.EndDate.Equal(*newEndDate)) {

		startDate := timeTable.StartDate
		if newStartDate != nil {
			startDate = *newStartDate
		}

		endDate := timeTable.EndDate
		if newEndDate != nil {
			endDate = *newEndDate
		}

		originalIsAvailable := timeTable.IsAvailable
		timeTable.IsAvailable = true
		if err := controller.db.Save(&timeTable).Error; err != nil {
			return fmt.Errorf("unable to temporarily update time table: %w", err)
		}

		isAvailable, err := controller.CheckRoomAvailability(timeTable.RoomID, startDate, endDate)

		timeTable.IsAvailable = originalIsAvailable
		if err := controller.db.Save(&timeTable).Error; err != nil {
			return fmt.Errorf("unable to restore time table status: %w", err)
		}

		if err != nil {
			return err
		}
		if !isAvailable {
			return fmt.Errorf("requested time slot is already booked")
		}
	}

	tx := controller.db.Begin()

	if eventName != nil {
		booking.EventName = *eventName
		if err := tx.Save(&booking).Error; err != nil {
			tx.Rollback()
			return fmt.Errorf("unable to update booking details: %w", err)
		}
	}

	if newStartDate != nil {
		timeTable.StartDate = *newStartDate
	}
	if newEndDate != nil {
		timeTable.EndDate = *newEndDate
	}

	if newStartDate != nil || newEndDate != nil {
		if err := tx.Save(&timeTable).Error; err != nil {
			tx.Rollback()
			return fmt.Errorf("unable to update booking dates: %w", err)
		}
	}

	if err := tx.Commit().Error; err != nil {
		return fmt.Errorf("transaction failed: %w", err)
	}

	fmt.Printf("Successfully updated booking with ID %d\n", bookingID)
	return nil
}

func (controller *BookingController) GetRoomBookings(roomID uint) ([]model.Booking, error) {
	var bookings []model.Booking
	if err := controller.db.Joins("JOIN time_tables ON bookings.time_table_id = time_tables.id").
		Where("time_tables.room_id = ?", roomID).
		Preload("TimeTable").
		Find(&bookings).Error; err != nil {
		return nil, fmt.Errorf("unable to retrieve bookings for room %d: %w", roomID, err)
	}
	return bookings, nil
}

func (controller *BookingController) GetAvailableRooms(startDate, endDate time.Time, roomType *model.RoomTypeEnum, capacity *int) ([]model.Room, error) {
	subQuery := controller.db.Table("time_tables").
		Select("room_id").
		Where("is_available = ? AND ((start_date <= ? AND end_date >= ?) OR (start_date <= ? AND end_date >= ?) OR (start_date >= ? AND end_date <= ?))",
			false, startDate, startDate, endDate, endDate, startDate, endDate).
		Group("room_id")

	query := controller.db.Where("is_room_out_of_service = ?", false).
		Where("room_id NOT IN (?)", subQuery)

	if roomType != nil {
		query = query.Where("room_type = ?", *roomType)
	}

	if capacity != nil {
		query = query.Where("capacity >= ?", *capacity)
	}

	var rooms []model.Room
	if err := query.Find(&rooms).Error; err != nil {
		return nil, fmt.Errorf("unable to find available rooms: %w", err)
	}

	return rooms, nil
}

func (controller *BookingController) GetBookingDetails(bookingID uint) (*model.Booking, error) {
	var booking model.Booking
	if err := controller.db.Preload("TimeTable").
		Preload("TimeTable.Room").
		First(&booking, bookingID).Error; err != nil {
		return nil, fmt.Errorf("unable to find booking with ID %d: %w", bookingID, err)
	}
	return &booking, nil
}
