// MEP-1013
package controller

import (
	model "ModEd/asset/model/spacemanagement"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type PermanentBookingController struct {
	db                     *gorm.DB
	lastCreatedScheduleIDs []uint
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

func (controller *PermanentBookingController) CreateWeeklySchedule(StartDate time.Time, EndDate time.Time, RoomID uint, CourseID uint, ClassID uint, FacultyID uint, DepartmentID uint, ProgramtypeID uint) error {
	controller.lastCreatedScheduleIDs = []uint{}

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

	currentDate := StartDate
	for currentDate.Before(EndDate) || currentDate.Equal(EndDate) {
		slotStart := currentDate
		slotEnd := currentDate.Add(time.Hour * 2)

		timetable := model.TimeTable{
			StartDate:   slotStart,
			EndDate:     slotEnd,
			RoomID:      RoomID,
			IsAvailable: false,
		}
		if err := controller.db.Create(&timetable).Error; err != nil {
			return fmt.Errorf("failed to create timetable: %w", err)
		}

		schedule := model.PermanentSchedule{
			TimeTableID:   timetable.ID,
			FacultyID:     FacultyID,
			DepartmentID:  DepartmentID,
			ProgramtypeID: ProgramtypeID,
			CourseId:      CourseID,
			ClassId:       ClassID,
		}

		if err := controller.db.Create(&schedule).Error; err != nil {
			return fmt.Errorf("failed to create permanent schedule: %w", err)
		}

		controller.lastCreatedScheduleIDs = append(controller.lastCreatedScheduleIDs, schedule.ID)

		currentDate = currentDate.AddDate(0, 0, 7)
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
