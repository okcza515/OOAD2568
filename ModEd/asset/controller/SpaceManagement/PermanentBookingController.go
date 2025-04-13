// MEP-1013
package spacemanagement

import (
	model "ModEd/asset/model/spacemanagement"
	//common "ModEd/common/model"
	"errors"
	"fmt"
	"strconv"
	"time"

	"gorm.io/gorm"
)

type PermanentScheduleController struct {
	db *gorm.DB
}

func (c *PermanentScheduleController) CheckRoomInService(roomID uint) (*bool, error) {
	var room model.Room

	err := c.db.First(&room, roomID).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("room not found")
	} else if err != nil {
		return nil, err
	}

	if room.IsRoomOutOfService {
		return nil, errors.New("room is out of service")
	}

	isInService := !room.IsRoomOutOfService
	return &isInService, nil
}

func (c *PermanentScheduleController) CreateSubjectSchedule(schedule *model.PermanentSchedule) error {
	if schedule.StartDate.IsZero() || schedule.EndDate.IsZero() {
		return errors.New("start date & end date are required")
	}
	if schedule.StartDate.After(schedule.EndDate) {
		return errors.New("start date can't be after end date")
	}
	if schedule.Faculty.Name == "" {
		return errors.New("faculty name is required")
	}
	if schedule.Department.Name == "" {
		return errors.New("department name is required")
	}
	if strconv.Itoa(int(schedule.ProgramType)) == "" {
		return errors.New("program type is required")
	}
	if schedule.Classroom.RoomID == 0 {
		return errors.New("classroom is required")
	}
	if schedule.Course.CourseId == 0 {
		return errors.New("course is required")
	}
	if schedule.Class.ClassId == 0 {
		return errors.New("class is required")
	}

	isInService, err := c.CheckRoomInService(schedule.Classroom.RoomID)
	if err != nil {
		return err
	}
	if !*isInService {
		return errors.New("room is unavailable for scheduled")
	}

	result := c.db.Create(schedule)
	if result.Error != nil {
		return result.Error
	}

	fmt.Println("Schedule created successfully")
	return nil
}

func (c *PermanentScheduleController) ScheduleRecurringSubject(baseSchedule *model.PermanentSchedule, recurringDays []string) error {
	if len(recurringDays) == 0 {
		return errors.New("recurring days are required")
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
