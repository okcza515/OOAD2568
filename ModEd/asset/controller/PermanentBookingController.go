// MEP-1013
package controller

import (
	model "ModEd/asset/model/spacemanagement"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type PermanentBookingController struct {
	db *gorm.DB
}

func NewPermanentBookingController(db *gorm.DB) *PermanentBookingController {
	return &PermanentBookingController{
		db: db,
	}
}

func (controller *PermanentBookingController) CreateWeeklySchedule(StartDate time.Time, EndDate time.Time, RoomID uint, CourseID uint, ClassID uint, FacultyID uint, DepartmentID uint, ProgramtypeID uint) error {
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
			return fmt.Errorf("Failed to creating timetable: %w", err)
		}

		schedule := model.PermanentSchedule{
			TimeTableID:   timetable.ID,
			FacultyID:     FacultyID,
			DepartmentID:  DepartmentID,
			ProgramtypeID: ProgramtypeID,
			Classroom:     room,
			CourseId:      CourseID,
			ClassId:       ClassID,
		}

		if err := controller.db.Create(&schedule).Error; err != nil {
			return fmt.Errorf("Failed to create permanent schedule: %w", err)
		}

		currentDate = currentDate.AddDate(0, 0, 7)
	}
	return nil
}
