// MEP-1013
package controller

import (
	"ModEd/asset/model"
	"ModEd/core"
	"ModEd/core/migration"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type BookingControllerInterface interface {
	Insert(data model.Booking) error
	RetrieveByID(id uint, preloads ...string) (model.Booking, error)
	UpdateByID(booking *model.Booking) error
	DeleteByID(id uint) error
	List(condition map[string]interface{}, preloads ...string) ([]model.Booking, error)
	CheckRoomAvailability(roomID uint, startDate, endDate time.Time) (bool, error)
	GetBookingsByTimeTable(timeTableID uint) ([]model.Booking, error)
}

type BookingController struct {
	db             *gorm.DB
	baseController *core.BaseController[model.Booking]
}

func NewBookingController() BookingControllerInterface {
	db := migration.GetInstance().DB
	return &BookingController{
		db:             db,
		baseController: core.NewBaseController[model.Booking](db),
	}
}

func (c *BookingController) Insert(data model.Booking) error {
	timeTable := model.TimeTable{
		StartDate:   data.TimeTable.StartDate,
		EndDate:     data.TimeTable.EndDate,
		RoomID:      data.TimeTable.RoomID,
		IsAvailable: false,
		BookingType: data.TimeTable.BookingType,
	}

	data.TimeTableID = timeTable.ID

	if err := c.baseController.Insert(data); err != nil {
		c.db.Delete(&timeTable)
		return err
	}

	return nil
}

func (c *BookingController) RetrieveByID(id uint, preloads ...string) (model.Booking, error) {
	return c.baseController.RetrieveByID(id, "TimeTable", "TimeTable.Room")
}

func (c *BookingController) UpdateByID(booking *model.Booking) error {
	var existing model.Booking
	if err := c.db.Preload("TimeTable").First(&existing, booking.ID).Error; err != nil {
		return err
	}

	tx := c.db.Begin()

	if err := tx.Model(&existing.TimeTable).Updates(map[string]interface{}{
		"start_date":   booking.TimeTable.StartDate,
		"end_date":     booking.TimeTable.EndDate,
		"booking_type": booking.TimeTable.BookingType,
	}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Model(&existing).Updates(map[string]interface{}{
		"event_name": booking.EventName,
		"user_role":  booking.UserRole,
	}).Error; err != nil {
		tx.Rollback()
		return err
	}

	isAvailable, err := c.CheckRoomAvailability(
		existing.TimeTable.RoomID,
		booking.TimeTable.StartDate,
		booking.TimeTable.EndDate,
	)
	if err != nil {
		tx.Rollback()
		return err
	}
	if !isAvailable {
		tx.Rollback()
		return fmt.Errorf("room is not available for the selected time period")
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}

func (c *BookingController) DeleteByID(id uint) error {
	_, err := c.RetrieveByID(id)
	if err != nil {
		return err
	}

	if err := c.baseController.DeleteByID(id); err != nil {
		return err
	}

	return nil
}

func (c *BookingController) List(condition map[string]interface{}, preloads ...string) ([]model.Booking, error) {
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
