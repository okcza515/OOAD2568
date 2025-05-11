//MEP-1013

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
	fmt.Println("===== Create New Permanent Schedule =====")

	var schedule model.PermanentSchedule
	schedule.TimeTable = model.TimeTable{}

	fmt.Print("Enter Room ID: ")
	roomIDStr := util.GetCommandInput()
	roomID, err := strconv.ParseUint(roomIDStr, 10, 32)
	if err != nil {
		fmt.Println("Invalid Room ID format")
		util.PressEnterToContinue()
		return err
	}
	schedule.TimeTable.RoomID = uint(roomID)

	fmt.Print("Enter start date (YYYY-MM-DD): ")
	startDateStr := util.GetCommandInput()
	fmt.Print("Enter start time (HH:MM): ")
	startTimeStr := util.GetCommandInput()

	startDateTime := startDateStr + " " + startTimeStr
	startDate, err := time.Parse("2006-01-02 15:04", startDateTime)
	if err != nil {
		fmt.Println("Invalid date/time format. Please use YYYY-MM-DD for date and HH:MM for time")
		util.PressEnterToContinue()
		return err
	}
	schedule.TimeTable.StartDate = startDate

	fmt.Print("Enter end date (YYYY-MM-DD): ")
	endDateStr := util.GetCommandInput()
	fmt.Print("Enter end time (HH:MM): ")
	endTimeStr := util.GetCommandInput()

	endDateTime := endDateStr + " " + endTimeStr
	endDate, err := time.Parse("2006-01-02 15:04", endDateTime)
	if err != nil {
		fmt.Println("Invalid date/time format. Please use YYYY-MM-DD for date and HH:MM for time")
		util.PressEnterToContinue()
		return err
	}
	schedule.TimeTable.EndDate = endDate

	fmt.Print("Enter recurrence end date (YYYY-MM-DD): ")
	recurrenceEndDateStr := util.GetCommandInput()
	recurrenceEndDate, err := time.Parse("2006-01-02", recurrenceEndDateStr)
	if err != nil {
		fmt.Println("Invalid date format. Please use YYYY-MM-DD")
		util.PressEnterToContinue()
		return err
	}

	fmt.Print("Enter Course ID: ")
	courseIDStr := util.GetCommandInput()
	courseID, err := strconv.ParseUint(courseIDStr, 10, 32)
	if err != nil {
		fmt.Println("Invalid Course ID format")
		util.PressEnterToContinue()
		return err
	}
	schedule.CourseId = uint(courseID)

	fmt.Print("Enter Class ID: ")
	classIDStr := util.GetCommandInput()
	classID, err := strconv.ParseUint(classIDStr, 10, 32)
	if err != nil {
		fmt.Println("Invalid Class ID format")
		util.PressEnterToContinue()
		return err
	}
	schedule.ClassId = uint(classID)

	fmt.Print("Enter Faculty ID: ")
	facultyIDStr := util.GetCommandInput()
	facultyID, err := strconv.ParseUint(facultyIDStr, 10, 32)
	if err != nil {
		fmt.Println("Invalid Faculty ID format")
		util.PressEnterToContinue()
		return err
	}
	schedule.FacultyID = uint(facultyID)

	fmt.Print("Enter Department ID: ")
	deptIDStr := util.GetCommandInput()
	deptID, err := strconv.ParseUint(deptIDStr, 10, 32)
	if err != nil {
		fmt.Println("Invalid Department ID format")
		util.PressEnterToContinue()
		return err
	}
	schedule.DepartmentID = uint(deptID)

	fmt.Print("Enter Program Type ID: ")
	progTypeIDStr := util.GetCommandInput()
	progTypeID, err := strconv.ParseUint(progTypeIDStr, 10, 32)
	if err != nil {
		fmt.Println("Invalid Program Type ID format")
		util.PressEnterToContinue()
		return err
	}
	schedule.ProgramtypeID = uint(progTypeID)

	for startDate.Before(recurrenceEndDate) || startDate.Equal(recurrenceEndDate) {
		isAvailable, err := handler.controller.CheckRoomAvailability(schedule.TimeTable.RoomID, startDate, endDate)
		if err != nil {
			fmt.Println("Error checking room availability:", err)
			util.PressEnterToContinue()
			return err
		}

		if !isAvailable {
			fmt.Printf("Room is not available for the time period starting on %s.\n", startDate.Format("2006-01-02 15:04"))
			util.PressEnterToContinue()
			return fmt.Errorf("room not available")
		}

		schedule.TimeTable.StartDate = startDate
		schedule.TimeTable.EndDate = endDate

		err = handler.controller.Insert(schedule)
		if err != nil {
			fmt.Println("Failed to create schedule:", err)
			util.PressEnterToContinue()
			return err
		}

		startDate = startDate.AddDate(0, 0, 7)
		endDate = endDate.AddDate(0, 0, 7)
	}

	fmt.Println("Recurring schedules created successfully!")
	util.PressEnterToContinue()
	return nil
}

// func (handler *CreatePermanentScheduleHandler) Execute() error {
// 	fmt.Println("===== Create New Permanent Schedule =====")

// 	var schedule model.PermanentSchedule
// 	schedule.TimeTable = model.TimeTable{}

// 	fmt.Print("Enter Room ID: ")
// 	roomIDStr := util.GetCommandInput()
// 	roomID, err := strconv.ParseUint(roomIDStr, 10, 32)
// 	if err != nil {
// 		fmt.Println("Invalid Room ID format")
// 		util.PressEnterToContinue()
// 		return err
// 	}
// 	schedule.TimeTable.RoomID = uint(roomID)

// 	fmt.Print("Enter start date (YYYY-MM-DD): ")
// 	startDateStr := util.GetCommandInput()
// 	fmt.Print("Enter start time (HH:MM): ")
// 	startTimeStr := util.GetCommandInput()

// 	startDateTime := startDateStr + " " + startTimeStr
// 	startDate, err := time.Parse("2006-01-02 15:04", startDateTime)
// 	if err != nil {
// 		fmt.Println("Invalid date/time format. Please use YYYY-MM-DD for date and HH:MM for time")
// 		util.PressEnterToContinue()
// 		return err
// 	}
// 	schedule.TimeTable.StartDate = startDate

// 	fmt.Print("Enter end date (YYYY-MM-DD): ")
// 	endDateStr := util.GetCommandInput()
// 	fmt.Print("Enter end time (HH:MM): ")
// 	endTimeStr := util.GetCommandInput()

// 	endDateTime := endDateStr + " " + endTimeStr
// 	endDate, err := time.Parse("2006-01-02 15:04", endDateTime)
// 	if err != nil {
// 		fmt.Println("Invalid date/time format. Please use YYYY-MM-DD for date and HH:MM for time")
// 		util.PressEnterToContinue()
// 		return err
// 	}
// 	schedule.TimeTable.EndDate = endDate

// 	isAvailable, err := handler.controller.CheckRoomAvailability(uint(roomID), startDate, endDate)
// 	if err != nil {
// 		fmt.Println("Error checking room availability:", err)
// 		util.PressEnterToContinue()
// 		return err
// 	}

// 	if !isAvailable {
// 		fmt.Println("Room is not available for the specified time period.")
// 		util.PressEnterToContinue()
// 		return fmt.Errorf("room not available")
// 	}

// 	fmt.Print("Enter Course ID: ")
// 	courseIDStr := util.GetCommandInput()
// 	courseID, err := strconv.ParseUint(courseIDStr, 10, 32)
// 	if err != nil {
// 		fmt.Println("Invalid Course ID format")
// 		util.PressEnterToContinue()
// 		return err
// 	}
// 	schedule.CourseId = uint(courseID)
// 	fmt.Print("Enter Class ID: ")
// 	classIDStr := util.GetCommandInput()
// 	classID, err := strconv.ParseUint(classIDStr, 10, 32)
// 	if err != nil {
// 		fmt.Println("Invalid Class ID format")
// 		util.PressEnterToContinue()
// 		return err
// 	}
// 	schedule.ClassId = uint(classID)
// 	fmt.Print("Enter Faculty ID: ")
// 	facultyIDStr := util.GetCommandInput()
// 	facultyID, err := strconv.ParseUint(facultyIDStr, 10, 32)
// 	if err != nil {
// 		fmt.Println("Invalid Faculty ID format")
// 		util.PressEnterToContinue()
// 		return err
// 	}
// 	schedule.FacultyID = uint(facultyID)
// 	fmt.Print("Enter Department ID: ")
// 	deptIDStr := util.GetCommandInput()
// 	deptID, err := strconv.ParseUint(deptIDStr, 10, 32)
// 	if err != nil {
// 		fmt.Println("Invalid Department ID format")
// 		util.PressEnterToContinue()
// 		return err
// 	}
// 	schedule.DepartmentID = uint(deptID)
// 	fmt.Print("Enter Program Type ID: ")
// 	progTypeIDStr := util.GetCommandInput()
// 	progTypeID, err := strconv.ParseUint(progTypeIDStr, 10, 32)
// 	if err != nil {
// 		fmt.Println("Invalid Program Type ID format")
// 		util.PressEnterToContinue()
// 		return err
// 	}
// 	schedule.ProgramtypeID = uint(progTypeID)
// 	err = handler.controller.Insert(schedule)
// 	if err != nil {
// 		fmt.Println("Failed to create schedule:", err)
// 		util.PressEnterToContinue()
// 		return err
// 	}
// 	fmt.Println("Schedule created successfully!")
// 	util.PressEnterToContinue()
// 	return nil
// }
