// MEP-1013

package handler

import (
	model "ModEd/asset/model"
	"ModEd/asset/util"
	"fmt"
	"strconv"
	"time"
)

type CreatePermanentScheduleHandler struct {
	controller interface {
		CheckRoomAvailability(roomID uint, startDate, endDate time.Time) (bool, error)
		Insert(dataContext model.PermanentSchedule) error
	}
}

func NewCreatePermanentScheduleHandler(controller interface {
	CheckRoomAvailability(roomID uint, startDate, endDate time.Time) (bool, error)
	Insert(dataContext model.PermanentSchedule) error
}) *CreatePermanentScheduleHandler {
	return &CreatePermanentScheduleHandler{
		controller: controller,
	}
}

func (handler *CreatePermanentScheduleHandler) Execute() error {
	fmt.Println("------- Create New Permanent Schedule -------")

	fmt.Println("Please enter Room ID:")
	roomIDStr := util.GetCommandInput()
	roomID, err := strconv.ParseUint(roomIDStr, 10, 32)
	if err != nil {
		fmt.Println("Invalid Room ID format")
		util.PressEnterToContinue()
		return err
	}

	fmt.Println("Please enter start date (YYYY-MM-DD):")
	startDateStr := util.GetCommandInput()
	fmt.Println("Please enter start time (HH:MM):")
	startTimeStr := util.GetCommandInput()

	startDateTime := startDateStr + " " + startTimeStr
	startDate, err := time.Parse("2006-01-02 15:04", startDateTime)
	if err != nil {
		fmt.Println("Invalid date/time format. Please use YYYY-MM-DD for date and HH:MM for time")
		util.PressEnterToContinue()
		return err
	}

	fmt.Println("Please enter end date (YYYY-MM-DD):")
	endDateStr := util.GetCommandInput()
	fmt.Println("Please enter end time (HH:MM):")
	endTimeStr := util.GetCommandInput()

	endDateTime := endDateStr + " " + endTimeStr
	endDate, err := time.Parse("2006-01-02 15:04", endDateTime)
	if err != nil {
		fmt.Println("Invalid date/time format. Please use YYYY-MM-DD for date and HH:MM for time")
		util.PressEnterToContinue()
		return err
	}

	fmt.Println("Please enter recurrence end date (YYYY-MM-DD):")
	recurrenceEndDateStr := util.GetCommandInput()
	recurrenceEndDate, err := time.Parse("2006-01-02", recurrenceEndDateStr)
	if err != nil {
		fmt.Println("Invalid date format. Please use YYYY-MM-DD")
		util.PressEnterToContinue()
		return err
	}

	fmt.Println("Please enter Course ID:")
	courseIDStr := util.GetCommandInput()
	courseID, err := strconv.ParseUint(courseIDStr, 10, 32)
	if err != nil {
		fmt.Println("Invalid Course ID format")
		util.PressEnterToContinue()
		return err
	}

	fmt.Println("Please enter Class ID:")
	classIDStr := util.GetCommandInput()
	classID, err := strconv.ParseUint(classIDStr, 10, 32)
	if err != nil {
		fmt.Println("Invalid Class ID format")
		util.PressEnterToContinue()
		return err
	}

	fmt.Println("Please enter Faculty ID:")
	facultyIDStr := util.GetCommandInput()
	facultyID, err := strconv.ParseUint(facultyIDStr, 10, 32)
	if err != nil {
		fmt.Println("Invalid Faculty ID format")
		util.PressEnterToContinue()
		return err
	}

	fmt.Println("Please enter Department ID:")
	deptIDStr := util.GetCommandInput()
	deptID, err := strconv.ParseUint(deptIDStr, 10, 32)
	if err != nil {
		fmt.Println("Invalid Department ID format")
		util.PressEnterToContinue()
		return err
	}

	fmt.Println("Please enter Program Type ID (0 for Regular, 1 for International):")
	progTypeIDStr := util.GetCommandInput()
	progTypeID, err := strconv.ParseUint(progTypeIDStr, 10, 32)
	if err != nil {
		fmt.Println("Invalid Program Type ID format")
		util.PressEnterToContinue()
		return err
	}

	schedule := &model.PermanentSchedule{
		TimeTable: model.TimeTable{
			RoomID:    uint(roomID),
			StartDate: startDate,
			EndDate:   endDate,
		},
		CourseId:      uint(courseID),
		ClassId:       uint(classID),
		FacultyID:     uint(facultyID),
		DepartmentID:  uint(deptID),
		ProgramtypeID: uint(progTypeID),
	}

	if err := schedule.Validate(); err != nil {
		fmt.Println("Validation error:", err)
		util.PressEnterToContinue()
		return err
	}

	fmt.Println("-------- Permanent Schedule Details --------")
	fmt.Println("Room ID:", schedule.TimeTable.RoomID)
	fmt.Println("Start Date:", schedule.TimeTable.StartDate.Format("2006-01-02 15:04"))
	fmt.Println("End Date:", schedule.TimeTable.EndDate.Format("2006-01-02 15:04"))
	fmt.Println("Recurrence End Date:", recurrenceEndDateStr)
	fmt.Println("Course ID:", schedule.CourseId)
	fmt.Println("Class ID:", schedule.ClassId)
	fmt.Println("Faculty ID:", schedule.FacultyID)
	fmt.Println("Department ID:", schedule.DepartmentID)
	fmt.Println("Program Type ID:", schedule.ProgramtypeID)

	fmt.Println("\nDo you want to create this Permanent Schedule? (y/n)")
	confirmStr := util.GetCommandInput()
	if confirmStr != "y" {
		fmt.Println("Permanent Schedule creation cancelled.")
		util.PressEnterToContinue()
		return nil
	}

	currentStartDate := startDate
	currentEndDate := endDate
	for currentStartDate.Before(recurrenceEndDate) || currentStartDate.Equal(recurrenceEndDate) {

		isAvailable, err := handler.controller.CheckRoomAvailability(schedule.TimeTable.RoomID, currentStartDate, currentEndDate)
		if err != nil {
			fmt.Println("Error checking room availability:", err)
			util.PressEnterToContinue()
			return err
		}

		if !isAvailable {
			fmt.Printf("Room is not available for the time period starting on %s.\n", currentStartDate.Format("2006-01-02 15:04"))
			util.PressEnterToContinue()
			return fmt.Errorf("room not available")
		}

		schedule.TimeTable.StartDate = currentStartDate
		schedule.TimeTable.EndDate = currentEndDate

		err = handler.controller.Insert(*schedule)
		if err != nil {
			fmt.Println("Failed to create schedule:", err)
			util.PressEnterToContinue()
			return err
		}

		currentStartDate = currentStartDate.AddDate(0, 0, 7)
		currentEndDate = currentEndDate.AddDate(0, 0, 7)
	}

	fmt.Println("Permanent Schedule created successfully!")
	util.PressEnterToContinue()
	return nil
}
