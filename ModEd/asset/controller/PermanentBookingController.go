// MEP-1013
package controller

import (
	model "ModEd/asset/model"
	"ModEd/core"
	"ModEd/utils/deserializer"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type PermanentBookingController struct {
	db                     *gorm.DB
	lastCreatedScheduleIDs []uint
	*core.BaseController[model.PermanentSchedule]
}

func NewPermanentBookingController(db *gorm.DB, BaseController *core.BaseController[model.PermanentSchedule]) *PermanentBookingController {
	return &PermanentBookingController{
		db:             db,
		BaseController: BaseController,
	}
}

func (controller *PermanentBookingController) CreateSchedule(schedule *model.PermanentSchedule) error {
	if err := controller.db.Create(schedule).Error; err != nil {
		return fmt.Errorf("failed to create schedule: %w", err)
	}
	return nil
}

func (controller *PermanentBookingController) SeedScheduleDatabase(path string) (schedule []*model.PermanentSchedule, err error) {
	deserializer, err := deserializer.NewFileDeserializer(path)
	if err != nil {
		return nil, errors.New("failed to create file deserializer")
	}
	if err := deserializer.Deserialize(&schedule); err != nil {
		return nil, errors.New("failed to deserialize schedule")
	}
	for _, schedule := range schedule {
		err := controller.CreateSchedule(schedule)
		if err != nil {
			return nil, errors.New("failed to seed Schedule DB")
		}
	}
	return schedule, nil
}

func (controller *PermanentBookingController) CheckRoomIsInService(RoomID uint) (bool, error) {
	var room model.Room
	if err := controller.db.First(&room, RoomID).Error; err != nil {
		return false, fmt.Errorf("unable to find room with ID %d: %w", RoomID, err)
	}
	if room.IsRoomOutOfService {
		return false, fmt.Errorf("room with ID %d is out of service", RoomID)
	}
	return true, nil
}

func (controller *PermanentBookingController) CreateWeeklySchedule(startDateTime, endDateTime time.Time, roomID uint, courseID uint, classID uint, facultyID uint, departmentID uint, programTypeID uint, semesterEndDate time.Time) error {
	controller.lastCreatedScheduleIDs = []uint{}

	if !startDateTime.Before(endDateTime) {
		return fmt.Errorf("start time must be before end time")
	}

	var room model.Room
	if err := controller.db.First(&room, roomID).Error; err != nil {
		return fmt.Errorf("unable to find room with ID %d: %w", roomID, err)
	}
	if room.IsRoomOutOfService {
		return fmt.Errorf("room with ID %d is out of service", roomID)
	}

	currentStart := startDateTime
	currentEnd := endDateTime

	for currentStart.Before(semesterEndDate) || currentStart.Equal(semesterEndDate) {
		timetable := model.TimeTable{
			StartDate:   currentStart,
			EndDate:     currentEnd,
			RoomID:      roomID,
			IsAvailable: false,
		}
		if err := controller.db.Create(&timetable).Error; err != nil {
			return fmt.Errorf("failed to create timetable: %w", err)
		}

		schedule := model.PermanentSchedule{
			TimeTableID:   timetable.ID,
			FacultyID:     facultyID,
			DepartmentID:  departmentID,
			ProgramtypeID: programTypeID,
			CourseId:      courseID,
			ClassId:       classID,
		}

		if err := controller.BaseController.Insert(schedule); err != nil {
			return fmt.Errorf("failed to create permanent schedule: %w", err)
		}
		controller.lastCreatedScheduleIDs = append(controller.lastCreatedScheduleIDs, schedule.ID)

		currentStart = currentStart.AddDate(0, 0, 7)
		currentEnd = currentEnd.AddDate(0, 0, 7)
	}

	return nil
}

func (controller *PermanentBookingController) GetLastCreatedScheduleIDs() []uint {
	return controller.lastCreatedScheduleIDs
}

func (controller *PermanentBookingController) GetAllPermanentBookings() ([]model.PermanentSchedule, error) {
	return controller.BaseController.List(nil, "TimeTable", "Room", "Course", "Class", "Faculty", "Department", "Programtype")
}

func (controller *PermanentBookingController) GetPermanentBookingByID(id uint) (model.PermanentSchedule, error) {
	return controller.BaseController.RetrieveByID(id, "TimeTable", "Room", "Course", "Class", "Faculty", "Department", "Programtype")
}

func (controller *PermanentBookingController) UpdatePermanentBooking(StartDate, EndDate time.Time, RoomID uint, CourseID uint, ClassID uint, FacultyID uint, DepartmentID uint, ProgramtypeID uint, ScheduleID uint) error {
	if !StartDate.Before(EndDate) {
		return fmt.Errorf("start date must be before end date")
	}

	var room model.Room
	if err := controller.db.First(&room, RoomID).Error; err != nil {
		return fmt.Errorf("unable to find room with ID %d: %w", RoomID, err)
	}
	if room.IsRoomOutOfService {
		return fmt.Errorf("room with ID %d is out of service", RoomID)
	}

	schedule, err := controller.RetrieveByID(ScheduleID)
	if err != nil {
		return fmt.Errorf("unable to find schedule with ID %d: %w", ScheduleID, err)
	}

	var timetable model.TimeTable
	if err := controller.db.First(&timetable, schedule.TimeTableID).Error; err != nil {
		return fmt.Errorf("unable to find timetable with ID %d: %w", schedule.TimeTableID, err)
	}

	timetable.StartDate = StartDate
	timetable.EndDate = EndDate
	timetable.RoomID = RoomID

	if err := controller.db.Save(&timetable).Error; err != nil {
		return fmt.Errorf("failed to update timetable: %w", err)
	}

	schedule.FacultyID = FacultyID
	schedule.DepartmentID = DepartmentID
	schedule.ProgramtypeID = ProgramtypeID
	schedule.CourseId = CourseID
	schedule.ClassId = ClassID

	if err := controller.UpdateByID(schedule); err != nil {
		return fmt.Errorf("failed to update permanent schedule: %w", err)
	}

	return nil
}

func (controller *PermanentBookingController) DeletePermanentSchedule(id uint) error {
	schedule, err := controller.BaseController.RetrieveByID(id)
	if err != nil {
		return fmt.Errorf("unable to find schedule with ID %d: %w", id, err)
	}

	if err := controller.BaseController.DeleteByID(id); err != nil {
		return fmt.Errorf("failed to delete permanent schedule: %w", err)
	}

	if err := controller.db.Delete(&model.TimeTable{}, schedule.TimeTableID).Error; err != nil {
		return fmt.Errorf("failed to delete timetable with ID %d: %w", schedule.TimeTableID, err)
	}

	return nil
}
