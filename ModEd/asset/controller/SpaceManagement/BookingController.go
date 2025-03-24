// MEP-1013
package spacemanagement

import (
	model "ModEd/asset/model/SpaceManagement"
	"fmt"
	"log"
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

func (controller *BookingController) CheckRoomAvailability(roomID uint) (bool, error) {
	var room model.Room
	if err := controller.db.First(&room, roomID).Error; err != nil {
		return false, fmt.Errorf("unable to find room with ID %d: %w", roomID, err)
	}

	if room.IsRoomOutOfService {
		return false, fmt.Errorf("room with ID %d is out of service", roomID)
	}

	return room.IsRoomOutOfService == false, nil
}

func (controller *BookingController) ResetTimeSlots() {
	resetTime := time.Now().Truncate(24 * time.Hour).Add(24 * time.Hour)

	go func() {
		for {
			now := time.Now()
			if now.After(resetTime) {
				if err := controller.resetAllBookings(); err != nil {
					log.Printf("Failed to reset bookings: %v", err)
				}
				resetTime = resetTime.Add(24 * time.Hour)
			}
			time.Sleep(time.Minute)
		}
	}()
}

func (controller *BookingController) resetAllBookings() error {
	var schedules []model.PermanentSchedule
	if err := controller.db.Where("is_available = ?", true).Find(&schedules).Error; err != nil {
		return fmt.Errorf("unable to fetch schedules: %w", err)
	}

	for _, schedule := range schedules {
		schedule.IsAvailable = true
		if err := controller.db.Save(&schedule).Error; err != nil {
			log.Printf("Failed to reset schedule ID %d: %v", schedule.ScheduleID, err)
		}
	}

	log.Println("Time slots successfully reset.")
	return nil
}

func (controller *BookingController) BookRoom(roomID uint, startDate, endDate time.Time) error {
	isAvailable, err := controller.CheckRoomAvailability(roomID)
	if err != nil || !isAvailable {
		return fmt.Errorf("room with ID %d is not available for booking: %w", roomID, err)
	}

	var room model.Room
	if err := controller.db.First(&room, roomID).Error; err != nil {
		return fmt.Errorf("unable to find room with ID %d: %w", roomID, err)
	}

	newBooking := model.PermanentSchedule{
		Classroom:   room,
		StartDate:   startDate,
		EndDate:     endDate,
		IsAvailable: false,
	}

	if err := controller.db.Create(&newBooking).Error; err != nil {
		return fmt.Errorf("unable to create new booking: %w", err)
	}

	log.Printf("Room %d successfully booked from %v to %v.", roomID, startDate, endDate)
	return nil
}

func (controller *BookingController) CancelBooking(bookingID uint) error {
	var booking model.PermanentSchedule
	if err := controller.db.First(&booking, bookingID).Error; err != nil {
		return fmt.Errorf("unable to find booking with ID %d: %w", bookingID, err)
	}

	booking.IsAvailable = true
	if err := controller.db.Save(&booking).Error; err != nil {
		return fmt.Errorf("unable to cancel booking: %w", err)
	}

	log.Printf("Booking with ID %d successfully canceled.", bookingID)
	return nil
}

func (controller *BookingController) UpdateBooking(bookingID uint, newStartDate, newEndDate time.Time) error {
	var booking model.PermanentSchedule
	if err := controller.db.First(&booking, bookingID).Error; err != nil {
		return fmt.Errorf("unable to find booking with ID %d: %w", bookingID, err)
	}

	booking.StartDate = newStartDate
	booking.EndDate = newEndDate
	if err := controller.db.Save(&booking).Error; err != nil {
		return fmt.Errorf("unable to update booking: %w", err)
	}

	log.Printf("Booking with ID %d successfully updated.", bookingID)
	return nil
}

func (controller *BookingController) GetRoomBookings(roomID uint) ([]model.PermanentSchedule, error) {
	var bookings []model.PermanentSchedule
	if err := controller.db.Where("room_id = ?", roomID).Find(&bookings).Error; err != nil {
		return nil, fmt.Errorf("unable to find bookings for room %d: %w", roomID, err)
	}
	return bookings, nil
}

func (controller *BookingController) GetAvailableRooms(startDate, endDate time.Time) ([]model.Room, error) {
	var rooms []model.Room
	if err := controller.db.Where("room_id NOT IN (SELECT room_id FROM permanent_schedules WHERE start_date < ? AND end_date > ?)", endDate, startDate).Find(&rooms).Error; err != nil {
		return nil, fmt.Errorf("unable to find available rooms: %w", err)
	}
	return rooms, nil
}

func (controller *BookingController) GetBookingDetails(bookingID uint) (*model.PermanentSchedule, error) {
	var booking model.PermanentSchedule
	if err := controller.db.First(&booking, bookingID).Error; err != nil {
		return nil, fmt.Errorf("unable to find booking with ID %d: %w", bookingID, err)
	}
	return &booking, nil
}
