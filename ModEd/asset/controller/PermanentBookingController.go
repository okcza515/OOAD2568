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
	NewPermanentSchedule(schedule model.PermanentSchedule) ([]model.PermanentSchedule, error)
	RetrieveByID(id uint, preload ...string) (model.PermanentSchedule, error)
	//RetrieveByID(id uint) (model.PermanentSchedule, error)
	UpdateByID(schedule model.PermanentSchedule) error
	DeleteByID(id uint) error
	DeleteAll() error
	// List(condition map[string]interface{}) ([]model.PermanentSchedule, error)
	List(condition map[string]interface{}, preloads ...string) ([]model.PermanentSchedule, error)

	// List(condition map[string]interface{}) ([]model.PermanentSchedule, error)
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

func (controller *PermanentBookingController) NewPermanentSchedule(schedule model.PermanentSchedule) ([]model.PermanentSchedule, error) {
	var timeTable model.TimeTable
	if err := controller.db.Where("id = ?", schedule.TimeTableID).First(&timeTable).Error; err != nil {
		return nil, errors.New("time table not found")
	}

	if !timeTable.IsAvailable {
		return nil, errors.New("time slot is unavailable")
	}

	var room model.Room
	if err := controller.db.Where("id = ?", timeTable.RoomID).First(&room).Error; err != nil {
		return nil, err
	}
	if room.IsRoomOutOfService {
		return nil, errors.New("room is out of service")
	}

	startDate := timeTable.StartDate
	endDate := timeTable.EndDate
	var schedules []model.PermanentSchedule

	tx := controller.db.Begin()
	for current := startDate; !current.After(endDate); current = current.AddDate(0, 0, 7) {

		newTimeTable := model.TimeTable{
			StartDate:   current,
			EndDate:     current.Add(timeTable.EndDate.Sub(timeTable.StartDate)),
			RoomID:      timeTable.RoomID,
			IsAvailable: false,
			BookingType: model.BOOKING_PERMANENT,
		}
		if err := tx.Create(&newTimeTable).Error; err != nil {
			tx.Rollback()
			return nil, err
		}

		newSchedule := model.PermanentSchedule{
			TimeTableID:   newTimeTable.ID,
			FacultyID:     schedule.FacultyID,
			DepartmentID:  schedule.DepartmentID,
			ProgramtypeID: schedule.ProgramtypeID,
			CourseId:      schedule.CourseId,
			ClassId:       schedule.ClassId,
		}
		if err := tx.Create(&newSchedule).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
		schedules = append(schedules, newSchedule)
	}
	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return schedules, nil
}

// func (controller *PermanentBookingController) RetrieveByID(id uint) (model.PermanentSchedule, error) {
// 	return controller.baseController.RetrieveByID(id, "TimeTable", "TimeTable.Room")
// }

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

// func (controller *PermanentBookingController) List(condition map[string]interface{}) ([]model.PermanentSchedule, error) {
// 	return controller.baseController.List(condition, "TimeTable", "TimeTable.Room")
// }

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
