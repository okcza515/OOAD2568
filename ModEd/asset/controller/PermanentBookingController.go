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
	RetrieveByID(id uint) (model.PermanentSchedule, error)
	UpdateByID(schedule model.PermanentSchedule) error
	DeleteByID(id uint) error
	DeleteAll() error
	List(condition map[string]interface{}) ([]model.PermanentSchedule, error)
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

func (controller *PermanentBookingController) RetrieveByID(id uint) (model.PermanentSchedule, error) {
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

func (controller *PermanentBookingController) List(condition map[string]interface{}) ([]model.PermanentSchedule, error) {
	return controller.baseController.List(condition, "TimeTable", "TimeTable.Room")
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

// type BookingStrategy interface {
// 	Create(controller *PermanentBookingController, startDateTime, endDateTime, semesterEndDate time.Time, roomID, courseID, classID, facultyID, departmentID, programtypeID uint) error
// }

// type WeeklyBookingStrategy struct{}

// func (strategy *WeeklyBookingStrategy) Create(controller *PermanentBookingController, startDateTime, endDateTime, semesterEndDate time.Time, roomID, courseID, classID, facultyID, departmentID, programtypeID uint) error {

// 	controller.lastCreatedScheduleIDs = []uint{}

// 	if !startDateTime.Before(endDateTime) {
// 		return fmt.Errorf("start time must be before end time")
// 	}

// 	inService, err := controller.CheckRoomIsInService(roomID)
// 	if err != nil {
// 		return err
// 	}

// 	if !inService {
// 		return fmt.Errorf("room with ID %d is out of service", roomID)
// 	}

// 	currentStart := startDateTime
// 	currentEnd := endDateTime

// 	for currentStart.Before(semesterEndDate) || currentStart.Equal(semesterEndDate) {
// 		timetable := model.TimeTable{
// 			StartDate:   currentStart,
// 			EndDate:     currentEnd,
// 			RoomID:      roomID,
// 			IsAvailable: false,
// 		}

// 		if err := controller.db.Create(&timetable).Error; err != nil {
// 			return fmt.Errorf("failed to create timetable: %w", err)
// 		}

// 		schedule := model.PermanentSchedule{
// 			TimeTableID:   timetable.ID,
// 			FacultyID:     facultyID,
// 			DepartmentID:  departmentID,
// 			ProgramtypeID: programtypeID,
// 			CourseId:      courseID,
// 			ClassId:       classID,
// 		}

// 		if err := controller.BaseController.Insert(schedule); err != nil {
// 			return fmt.Errorf("failed to create permanent schedule: %w", err)
// 		}

// 		controller.lastCreatedScheduleIDs = append(controller.lastCreatedScheduleIDs, schedule.ID)

// 		currentStart = currentStart.AddDate(0, 0, 7)
// 		currentEnd = currentEnd.AddDate(0, 0, 7)

// 	}

// 	return nil

// }

// type PermanentBookingController struct {
// 	db                     *gorm.DB
// 	lastCreatedScheduleIDs []uint
// 	*core.BaseController[model.PermanentSchedule]
// 	bookingStrategies map[string]BookingStrategy
// }

// func NewPermanentBookingController(db *gorm.DB) *PermanentBookingController {
// 	controller := &PermanentBookingController{
// 		db:                db,
// 		BaseController:    core.NewBaseController[model.PermanentSchedule](db),
// 		bookingStrategies: make(map[string]BookingStrategy),
// 	}

// 	controller.bookingStrategies["weekly"] = &WeeklyBookingStrategy{}

// 	return controller
// }

// func (controller *PermanentBookingController) RegisterBookingStrategy(name string, strategy BookingStrategy) {
// 	controller.bookingStrategies[name] = strategy
// }

// func (controller *PermanentBookingController) CreateBookingWithStrategy(strategyName string, startDateTime, endDateTime time.Time, roomID, courseID, classID, facultyID, departmentID, programTypeID uint, semesterEndDate time.Time) error {

// 	strategy, exists := controller.bookingStrategies[strategyName]
// 	if !exists {
// 		return fmt.Errorf("booking strategy '%s' not found", strategyName)
// 	}

// 	return strategy.Create(controller, startDateTime, endDateTime, semesterEndDate, roomID, courseID, classID, facultyID, departmentID, programTypeID)
// }

// func (controller *PermanentBookingController) CreateWeeklySchedule(startDateTime, endDateTime time.Time, roomID, courseID, classID, facultyID, departmentID, programTypeID uint, semesterEndDate time.Time) error {

// 	return controller.CreateBookingWithStrategy("weekly", startDateTime, endDateTime, roomID, courseID, classID, facultyID, departmentID, programTypeID, semesterEndDate)
// }

// func (controller *PermanentBookingController) CheckRoomIsInService(roomID uint) (bool, error) {
// 	var room model.Room
// 	if err := controller.db.First(&room, roomID).Error; err != nil {
// 		return false, fmt.Errorf("unable to find room with ID %d: %w", roomID, err)
// 	}

// 	if room.IsRoomOutOfService {
// 		return false, fmt.Errorf("room with ID %d is out of service", roomID)
// 	}

// 	return true, nil
// }

// func (controller *PermanentBookingController) SeedScheduleDatabase(path string) (schedules []*model.PermanentSchedule, err error) {
// 	deserializer, err := deserializer.NewFileDeserializer(path)
// 	if err != nil {
// 		return nil, errors.New("failed to create file deserializer")
// 	}

// 	if err := deserializer.Deserialize(&schedules); err != nil {
// 		return nil, errors.New("failed to deserialize schedule")
// 	}

// 	for _, schedule := range schedules {
// 		if err := controller.BaseController.Insert(*schedule); err != nil {
// 			return nil, fmt.Errorf("failed to seed Schedule DB: %w", err)
// 		}
// 	}

// 	return schedules, nil
// }

// func (controller *PermanentBookingController) GetLastCreatedScheduleIDs() []uint {
// 	return controller.lastCreatedScheduleIDs
// }

// func (controller *PermanentBookingController) GetAllPermanentBookings() ([]model.PermanentSchedule, error) {
// 	return controller.BaseController.List(nil, "TimeTable.Room", "Course", "Class", "Faculty", "Department", "Programtype")
// }

// func (controller *PermanentBookingController) GetPermanentBookingByID(id uint) (model.PermanentSchedule, error) {
// 	return controller.BaseController.RetrieveByID(id, "TimeTable.Room", "Course", "Class", "Faculty", "Department", "Programtype")
// }

// func (controller *PermanentBookingController) UpdateByID(startDate, endDate time.Time, roomID, courseID, classID, facultyID, departmentID, programTypeID, scheduleID uint) error {
// 	if !startDate.Before(endDate) {
// 		return fmt.Errorf("start date must be before end date")
// 	}

// 	inService, err := controller.CheckRoomIsInService(roomID)
// 	if err != nil {
// 		return err
// 	}

// 	if !inService {
// 		return fmt.Errorf("room with ID %d is out of service", roomID)
// 	}

// 	schedule, err := controller.BaseController.RetrieveByID(scheduleID)
// 	if err != nil {
// 		return fmt.Errorf("unable to find schedule with ID %d: %w", scheduleID, err)
// 	}

// 	var timetable model.TimeTable
// 	if err := controller.db.First(&timetable, schedule.TimeTableID).Error; err != nil {
// 		return fmt.Errorf("unable to find timetable with ID %d: %w", schedule.TimeTableID, err)
// 	}

// 	timetable.StartDate = startDate
// 	timetable.EndDate = endDate
// 	timetable.RoomID = roomID

// 	if err := controller.db.Save(&timetable).Error; err != nil {
// 		return fmt.Errorf("failed to update timetable: %w", err)
// 	}

// 	schedule.FacultyID = facultyID
// 	schedule.DepartmentID = departmentID
// 	schedule.ProgramtypeID = programTypeID
// 	schedule.CourseId = courseID
// 	schedule.ClassId = classID

// 	if err := controller.BaseController.UpdateByID(schedule); err != nil {
// 		return fmt.Errorf("failed to update permanent schedule: %w", err)
// 	}

// 	return nil
// }

// func (controller *PermanentBookingController) DeleteByID(id uint) error {
// 	schedule, err := controller.BaseController.RetrieveByID(id)
// 	if err != nil {
// 		return fmt.Errorf("unable to find schedule with ID %d: %w", id, err)
// 	}

// 	if err := controller.BaseController.DeleteByID(id); err != nil {
// 		return fmt.Errorf("failed to delete permanent schedule: %w", err)
// 	}

// 	if err := controller.db.Delete(&model.TimeTable{}, schedule.TimeTableID).Error; err != nil {
// 		return fmt.Errorf("failed to delete timetable with ID %d: %w", schedule.TimeTableID, err)
// 	}

// 	return nil
// }
