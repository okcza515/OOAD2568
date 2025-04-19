// MEP-1013
package spacemanagement

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

func (controller * PermanentBookingController) CheckRoomInService(RoomID uint) (bool, error) {
	var room model.Room
	if err := controller.db.First(&room, RoomID).Error; err != nil {
		return false, fmt.Errorf("Room %d is out of service: %w", RoomID, err)
	}
	return !room.IsRoomOutOfService, nil
}

func (controller *PermanentBookingController) CreatePermanentSchedule(RoomID uint, Faculty string, Department string, Programtype string, CourseId string, ClassId string, StartDate, EndDate time.Time ) (*model.PermanentSchedule, error) {
	schedule := &model.PermanentSchedule{
		RoomID:      RoomID,
		Faculty:     Faculty,
		Department:  Department,
		ProgramType: Programtype,
		CourseId:    CourseId,
		ClassId:     ClassId,
	}
	if err := controller.db.Create(schedule).Error; err != nil {
		return nil, fmt.Errorf("failed to create schedule: %w", err)
	}
	return schedule, nil
}

func (controller *PermanentBookingController) 

func (controller *PermanentBookingController) GetAll() (*[]model.PermanentSchedule, error) {
	schedules := new([]model.PermanentSchedule)
	if err := controller.db.Find(&schedules).Error; err != nil {
		return nil, fmt.Errorf("failed to get all schedules: %w", err)
	}
	return schedules, nil
}