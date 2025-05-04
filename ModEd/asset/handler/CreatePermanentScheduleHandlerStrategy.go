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
		NewPermanentSchedule(schedule model.PermanentSchedule) ([]model.PermanentSchedule, error)
	}
}

func NewCreatePermanentScheduleHandler(controller interface {
	CheckRoomAvailability(roomID uint, startDate, endDate time.Time) (bool, error)
	NewPermanentSchedule(schedule model.PermanentSchedule) ([]model.PermanentSchedule, error)
}) *CreatePermanentScheduleHandler {
	return &CreatePermanentScheduleHandler{
		controller: controller,
	}
}

func (h *CreatePermanentScheduleHandler) Execute() error {
	fmt.Println("===== Create New Permanent Schedule =====")

	fmt.Print("Enter Room ID: ")
	roomIDStr := util.GetCommandInput()
	roomID, err := strconv.ParseUint(roomIDStr, 10, 32)
	if err != nil {
		fmt.Println("Invalid Room ID format")
		util.PressEnterToContinue()
		return err
	}

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

	fmt.Print("Enter end time (HH:MM): ")
	endTimeStr := util.GetCommandInput()

	endDateTime := startDateStr + " " + endTimeStr
	endDate, err := time.Parse("2006-01-02 15:04", endDateTime)
	if err != nil {
		fmt.Println("Invalid time format. Please use HH:MM for time")
		util.PressEnterToContinue()
		return err
	}

	fmt.Print("Enter semester end date (YYYY-MM-DD): ")
	semesterEndStr := util.GetCommandInput()
	_, err = time.Parse("2006-01-02", semesterEndStr)
	if err != nil {
		fmt.Println("Invalid date format. Please use YYYY-MM-DD")
		util.PressEnterToContinue()
		return err
	}

	isAvailable, err := h.controller.CheckRoomAvailability(uint(roomID), startDate, endDate)
	if err != nil {
		fmt.Println("Error checking room availability:", err)
		util.PressEnterToContinue()
		return err
	}

	if !isAvailable {
		fmt.Println("Room is not available for the specified time period.")
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

	fmt.Print("Enter Class ID: ")
	classIDStr := util.GetCommandInput()
	classID, err := strconv.ParseUint(classIDStr, 10, 32)
	if err != nil {
		fmt.Println("Invalid Class ID format")
		util.PressEnterToContinue()
		return err
	}

	fmt.Print("Enter Faculty ID: ")
	facultyIDStr := util.GetCommandInput()
	facultyID, err := strconv.ParseUint(facultyIDStr, 10, 32)
	if err != nil {
		fmt.Println("Invalid Faculty ID format")
		util.PressEnterToContinue()
		return err
	}

	fmt.Print("Enter Department ID: ")
	deptIDStr := util.GetCommandInput()
	deptID, err := strconv.ParseUint(deptIDStr, 10, 32)
	if err != nil {
		fmt.Println("Invalid Department ID format")
		util.PressEnterToContinue()
		return err
	}

	fmt.Print("Enter Program Type ID: ")
	progTypeIDStr := util.GetCommandInput()
	progTypeID, err := strconv.ParseUint(progTypeIDStr, 10, 32)
	if err != nil {
		fmt.Println("Invalid Program Type ID format")
		util.PressEnterToContinue()
		return err
	}

	// timeTable := model.TimeTable{
	//     StartDate:   startDate,
	//     EndDate:     endDate,
	//     RoomID:      uint(roomID),
	//     IsAvailable: true,
	//     BookingType: model.BOOKING_PERMANENT,
	// }

	schedule := model.PermanentSchedule{
		TimeTableID:   0,
		CourseId:      uint(courseID),
		ClassId:       uint(classID),
		FacultyID:     uint(facultyID),
		DepartmentID:  uint(deptID),
		ProgramtypeID: uint(progTypeID),
	}

	schedules, err := h.controller.NewPermanentSchedule(schedule)
	if err != nil {
		fmt.Println("Failed to create schedule:", err)
		util.PressEnterToContinue()
		return err
	}

	fmt.Println("Schedule created successfully!")
	fmt.Printf("Created %d schedule entries\n", len(schedules))

	util.PressEnterToContinue()
	return nil
}
