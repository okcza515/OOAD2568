package controller

import (
	"ModEd/asset/model"
	"ModEd/core"
	"ModEd/core/migration"
	"ModEd/utils/deserializer"
	"errors"
	"time"

	"gorm.io/gorm"
)

type BookingControllerInterface interface {
	SeedBookingsDatabase(path string) ([]*model.Booking, error)
	CreateBooking(booking model.Booking) (model.Booking, error)
	GetBooking(id uint) (model.Booking, error)
	UpdateBooking(booking model.Booking) error
	DeleteBooking(id uint) error
	ListBookings(condition map[string]interface{}) ([]model.Booking, error)
	CheckRoomAvailability(roomID uint, startDate, endDate time.Time) (bool, error)
	GetBookingsByTimeTable(timeTableID uint) ([]model.Booking, error)
}

type BookingController struct {
	db *gorm.DB
	baseController *core.BaseController[model.Booking]
	observers []SpaceManagementObserverInterface[model.Booking]
}

func NewBookingController() BookingControllerInterface {
	db := migration.GetInstance().DB
	return &BookingController{
		db: db,
		baseController: core.NewBaseController[model.Booking](db),
		observers: make([]SpaceManagementObserverInterface[model.Booking], 0),
	}
}

func (c *BookingController) RegisterObserver(observer SpaceManagementObserverInterface[model.Booking]) {
	c.observers = append(c.observers, observer)
}

func (c *BookingController) NotifyObservers(eventType string, booking model.Booking) {
	for _, observer := range c.observers {
		observer.HandleEvent(eventType, booking)
	}
}

func (c *BookingController) SeedBookingsDatabase(path string) (bookings []*model.Booking, err error) {
	deserializer, err := deserializer.NewFileDeserializer(path)
	if err != nil {
		return nil, errors.New("failed to create file deserializer")
	}
	if err := deserializer.Deserialize(&bookings); err != nil {
		return nil, errors.New("failed to deserialize bookings")
	}
	for _, booking := range bookings {
		err := c.baseController.Insert(*booking)
		if err != nil {
			return nil, errors.New("failed to seed Booking DB")
		}
	}
	return bookings, nil
}

func (c *BookingController) CreateBooking(booking model.Booking) (model.Booking, error) {
	timeTable := c.db.Model(&model.TimeTable{}).Where("id = ?", booking.TimeTableID).First(&model.TimeTable{}).Row()
	if timeTable == nil {
		return booking, errors.New("time table not found")
	}
	
	var startDate, endDate time.Time
	var roomID uint
	var isAvailable bool
	
	if err := timeTable.Scan(&startDate, &endDate, &roomID, &isAvailable); err != nil {
		return booking, err
	}
	
	if !isAvailable {
		return booking, errors.New("time slot is not available")
	}
	
	var room model.Room
	if err := c.db.Where("id = ?", roomID).First(&room).Error; err != nil {
		return booking, err
	}
	
	if room.IsRoomOutOfService {
		return booking, errors.New("room is out of service")
	}
	
	tx := c.db.Begin()
	
	if err := tx.Model(&model.TimeTable{}).Where("id = ?", booking.TimeTableID).
		Updates(map[string]interface{}{
			"is_available": false,
			"booking_type": model.BOOKING_TEMPORARY,
		}).Error; err != nil {
		tx.Rollback()
		return booking, err
	}
	
	if err := tx.Create(&booking).Error; err != nil {
		tx.Rollback()
		return booking, err
	}
	
	if err := tx.Commit().Error; err != nil {
		return booking, err
	}
	
	c.NotifyObservers("booking_created", booking)
	
	return booking, nil
}

func (c *BookingController) GetBooking(id uint) (model.Booking, error) {
	return c.baseController.RetrieveByID(id, "TimeTable", "TimeTable.Room")
}

func (c *BookingController) UpdateBooking(booking model.Booking) error {
	_, err := c.GetBooking(booking.ID)
	if err != nil {
		return err
	}
	
	err = c.baseController.UpdateByID(booking)
	if err != nil {
		return err
	}
	
	c.NotifyObservers("booking_updated", booking)
	
	return nil
}

func (c *BookingController) DeleteBooking(id uint) error {
	booking, err := c.GetBooking(id)
	if err != nil {
		return err
	}
	
	tx := c.db.Begin()
	
	if err := tx.Model(&model.TimeTable{}).Where("id = ?", booking.TimeTableID).
		Updates(map[string]interface{}{
			"is_available": true,
			"booking_type": nil,
		}).Error; err != nil {
		tx.Rollback()
		return err
	}
	
	if err := tx.Delete(&booking).Error; err != nil {
		tx.Rollback()
		return err
	}
	
	if err := tx.Commit().Error; err != nil {
		return err
	}
	
	c.NotifyObservers("booking_deleted", booking)
	
	return nil
}

func (c *BookingController) ListBookings(condition map[string]interface{}) ([]model.Booking, error) {
	return c.baseController.List(condition, "TimeTable", "TimeTable.Room")
}

func (c *BookingController) CheckRoomAvailability(roomID uint, startDate, endDate time.Time) (bool, error) {
	var room model.Room
	if err := c.db.Where("id = ?", roomID).First(&room).Error; err != nil {
		return false, err
	}
	
	if room.IsRoomOutOfService {
		return false, nil
	}
	
	var count int64
	err := c.db.Model(&model.TimeTable{}).
		Where("room_id = ? AND is_available = ? AND ((start_date <= ? AND end_date >= ?) OR (start_date <= ? AND end_date >= ?) OR (start_date >= ? AND end_date <= ?))",
			roomID, false, startDate, startDate, endDate, endDate, startDate, endDate).
		Count(&count).Error
	
	if err != nil {
		return false, err
	}
	
	return count == 0, nil
}

func (c *BookingController) GetBookingsByTimeTable(timeTableID uint) ([]model.Booking, error) {
	condition := map[string]interface{}{"time_table_id": timeTableID}
	return c.baseController.List(condition, "TimeTable", "TimeTable.Room")
}
