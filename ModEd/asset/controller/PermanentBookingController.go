// MEP-1013
package controller

import (
	model "ModEd/asset/model"
	"ModEd/core"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type PermanentBookingController struct {
	db                     *gorm.DB
	lastCreatedScheduleIDs []uint
	*core.BaseController[model.PermanentSchedule]
}

func NewPermanentBookingController(db *gorm.DB) *PermanentBookingController {
	return &PermanentBookingController{
		db: db,
	}
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

		if err := controller.db.Create(&schedule).Error; err != nil {
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
	var bookings []model.PermanentSchedule
	if err := controller.db.Preload("TimeTable").Find(&bookings).Error; err != nil {
		return nil, fmt.Errorf("unable to retrieve permanent bookings: %w", err)
	}
	return bookings, nil
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

	var schedule model.PermanentSchedule
	if err := controller.db.First(&schedule, ScheduleID).Error; err != nil {
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

	if err := controller.db.Save(&schedule).Error; err != nil {
		return fmt.Errorf("failed to update permanent schedule: %w", err)
	}

	return nil
}
