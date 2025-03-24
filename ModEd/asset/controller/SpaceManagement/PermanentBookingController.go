// MEP-1013
package spacemanagement

import (
	model "ModEd/asset/model/SpaceManagement"
	//common "ModEd/common/model"
	"errors"
	"fmt"
	"time"
	"strconv"

	"gorm.io/gorm"
)

type PermanentScheduleController struct {
	DB *gorm.DB
}

func (c *PermanentScheduleController) CheckRoomInService(roomID uint) (*bool, error) {
	var room model.Room

	err := c.DB.First(&room, roomID).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("Room not found")
	} else if err != nil {
		return nil, err
	}

	if room.IsRoomOutOfService {
		return nil, errors.New("Room is out of service")
	}

	isInService := !room.IsRoomOutOfService
	return &isInService, nil
}

func (c *PermanentScheduleController) CreateSubjectSchedule(schedule *model.PermanentSchedule) error {
	if schedule.StartDate.IsZero() || schedule.EndDate.IsZero() {
		return errors.New("Start date & end date are required")
	}
	if schedule.StartDate.After(schedule.EndDate) {
		return errors.New("Start date can't be after end date")
	}
	if schedule.Faculty.Name == "" {
		return errors.New("Faculty name is required")
	}
	if schedule.Department.Name == "" {
		return errors.New("Department name is required")
	}
	if strconv.Itoa(int(schedule.ProgramType)) == "" {
		return errors.New("Program type is required")
	}
	if schedule.Classroom.RoomID == 0 {
		return errors.New("Classroom is required")
	}
	if schedule.Course.CourseId == 0 {
		return errors.New("Course is required")
	}
	if schedule.Class.ClassId == 0 {
		return errors.New("Class is required")
	}

	isInService, err := c.CheckRoomInService(schedule.Classroom.RoomID)
	if err != nil {
		return err
	}
	if !*isInService {
		return errors.New("Room is unavailable for scheduled")
	}

	result := c.DB.Create(schedule)
	if result.Error != nil {
		return result.Error
	}

	fmt.Println("Schedule created successfully")
	return nil
}

func (c *PermanentScheduleController) ScheduleRecurringSubject(baseSchedule *model.PermanentSchedule, recurringDays []string) error {
	if len(recurringDays) == 0 {
		return errors.New("Recurring days are required")
	}

	dayToInt := map[string]time.Weekday{
		"Sunday":    time.Sunday,
		"Monday":    time.Monday,
		"Tuesday":   time.Tuesday,
		"Wednesday": time.Wednesday,
		"Thursday":  time.Thursday,
		"Friday":    time.Friday,
		"Saturday":  time.Saturday,
	}

	for _, day := range recurringDays {
		weekday, valid := dayToInt[day]
		if !valid {
			return errors.New("Invalid recurring day: " + day)
		}

		for baseSchedule.StartDate.Weekday() != weekday {
			baseSchedule.StartDate = baseSchedule.StartDate.AddDate(0, 0, 1)
			baseSchedule.EndDate = baseSchedule.EndDate.AddDate(0, 0, 1)
		}

		newSchedule := *baseSchedule
		err := c.CreateSubjectSchedule(&newSchedule)
		if err != nil {
			return err
		}

		baseSchedule.StartDate = baseSchedule.StartDate.AddDate(0, 0, 7)
		baseSchedule.EndDate = baseSchedule.EndDate.AddDate(0, 0, 7)

	}
	return nil
}
