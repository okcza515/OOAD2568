package controller

import (
    "ModEd/asset/model"
    "ModEd/core"
    "ModEd/core/migration"
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
    observers      []SpaceManagementObserverInterface[model.Booking]
}

func NewBookingController() BookingControllerInterface {
    db := migration.GetInstance().DB
    return &BookingController{
        db:             db,
        baseController: core.NewBaseController[model.Booking](db),
        observers:      make([]SpaceManagementObserverInterface[model.Booking], 0),
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

func (c *BookingController) Insert(data model.Booking) error {
    if err := c.baseController.Insert(data); err != nil {
        return err
    }
    c.NotifyObservers("booking_created", data)
    return nil
}

func (c *BookingController) RetrieveByID(id uint, preloads ...string) (model.Booking, error) {
    return c.baseController.RetrieveByID(id, "TimeTable", "TimeTable.Room")
}

func (c *BookingController) UpdateByID(booking *model.Booking) error {
    var existing model.Booking
    if err := c.db.First(&existing, booking.ID).Error; err != nil {
        return err
    }

    if err := c.db.Model(&existing).Updates(booking).Error; err != nil {
        return err
    }

    c.NotifyObservers("booking_updated", *booking)
    return nil
}

func (c *BookingController) DeleteByID(id uint) error {
    booking, err := c.RetrieveByID(id)
    if err != nil {
        return err
    }

    if err := c.baseController.DeleteByID(id); err != nil {
        return err
    }

    c.NotifyObservers("booking_deleted", booking)
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