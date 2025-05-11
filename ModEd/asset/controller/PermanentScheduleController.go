// MEP-1013
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

type PermanentBookingControllerInterface interface {
	SeedPermanentBookingSchedule(path string) ([]*model.PermanentSchedule, error)
	//NewPermanentSchedule(schedule model.PermanentSchedule) ([]model.PermanentSchedule, error)
	Insert(data model.PermanentSchedule) error
	RetrieveByID(id uint, preload ...string) (model.PermanentSchedule, error)
	UpdateByID(schedule model.PermanentSchedule) error
	DeleteByID(id uint) error
	DeleteAll() error
	List(condition map[string]interface{}, preloads ...string) ([]model.PermanentSchedule, error)
	CheckRoomAvailability(roomID uint, startDate, endDate time.Time) (bool, error)
}
type PermanentBookingController struct {
	db             *gorm.DB
	baseController *core.BaseController[model.PermanentSchedule]
}

func NewPermanentBookingController() PermanentBookingControllerInterface {
	db := migration.GetInstance().DB
	return &PermanentBookingController{
		db:             db,
		baseController: core.NewBaseController[model.PermanentSchedule](db),
	}
}
func (controller *PermanentBookingController) SeedPermanentBookingSchedule(path string) (schedule []*model.PermanentSchedule, err error) {
	deserializer, err := deserializer.NewFileDeserializer(path)
	if err != nil {
		return nil, errors.New("failed to create file deserializer")
	}
	if err := deserializer.Deserialize(&schedule); err != nil {
		return nil, errors.New("failed to deserialize schedule")
	}
	for _, schedule := range schedule {
		err := controller.baseController.Insert(*schedule)
		if err != nil {
			return nil, errors.New("failed to seed Schedule DB")
		}
	}
	return schedule, nil
}

func (controller *PermanentBookingController) Insert(data model.PermanentSchedule) error {
	timeTable := model.TimeTable{
		StartDate:   data.TimeTable.StartDate,
		EndDate:     data.TimeTable.EndDate,
		RoomID:      data.TimeTable.RoomID,
		IsAvailable: false,
	}

	data.TimeTableID = timeTable.ID

	if err := controller.baseController.Insert(data); err != nil {
		controller.db.Delete(&timeTable)
		return err
	}

	return nil
}

func (controller *PermanentBookingController) RetrieveByID(id uint, preloads ...string) (model.PermanentSchedule, error) {
	return controller.baseController.RetrieveByID(id, "TimeTable", "TimeTable.Room")
}

func (controller *PermanentBookingController) UpdateByID(schedule model.PermanentSchedule) error {
	_, err := controller.RetrieveByID(schedule.ID)
	if err != nil {
		return err
	}

	err = controller.baseController.UpdateByID(schedule)
	if err != nil {
		return err
	}

	return nil
}

func (controller *PermanentBookingController) DeleteByID(id uint) error {
	schedule, err := controller.RetrieveByID(id)
	if err != nil {
		return err
	}

	tx := controller.db.Begin()

	if err := tx.Model(&model.TimeTable{}).Where("id = ?", schedule.TimeTableID).
		Updates(map[string]interface{}{
			"is_available": true,
			"booking_type": nil,
		}).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Delete(&schedule).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}

func (c *PermanentBookingController) List(condition map[string]interface{}, preloads ...string) ([]model.PermanentSchedule, error) {
	return c.baseController.List(condition, "TimeTable", "TimeTable.Room")
}

func (controller *PermanentBookingController) DeleteAll() error {
	records, err := controller.List(nil)
	if err != nil {
		return err
	}

	for _, record := range records {
		if err := controller.DeleteByID(record.ID); err != nil {
			return err
		}
	}

	return nil
}

func (controller *PermanentBookingController) CheckRoomAvailability(roomID uint, startDate, endDate time.Time) (bool, error) {
	var room model.Room
	if err := controller.db.Where("id = ?", roomID).First(&room).Error; err != nil {
		return false, err
	}

	if room.IsRoomOutOfService {
		return false, nil
	}

	var count int64
	err := controller.db.Model(&model.TimeTable{}).
		Where("room_id = ? AND is_available = ? AND ((start_date <= ? AND end_date >= ?) OR (start_date <= ? AND end_date >= ?) OR (start_date >= ? AND end_date <= ?))",
			roomID, false, startDate, startDate, endDate, endDate, startDate, endDate).
		Count(&count).Error

	if err != nil {
		return false, err
	}

	return count == 0, nil
}
