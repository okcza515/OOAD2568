// MEP-1013

package handler

import (
	model "ModEd/asset/model"
	"ModEd/asset/util"
	"fmt"
	"strconv"
	"time"
)

type UpdatePermanentScheduleHandler struct {
	controller interface {
		RetrieveByID(id uint, preload ...string) (model.PermanentSchedule, error)
		UpdateByID(schedule model.PermanentSchedule) error
		CheckRoomAvailability(roomID uint, startDate, endDate time.Time) (bool, error)
	}
}

func NewUpdatePermanentScheduleHandler(controller interface {
	RetrieveByID(id uint, preload ...string) (model.PermanentSchedule, error)
	UpdateByID(schedule model.PermanentSchedule) error
	CheckRoomAvailability(roomID uint, startDate, endDate time.Time) (bool, error)
}) *UpdatePermanentScheduleHandler {
	return &UpdatePermanentScheduleHandler{
		controller: controller,
	}
}

func (h *UpdatePermanentScheduleHandler) Execute() error {
	fmt.Println("===== Update Permanent Schedule =====")

	fmt.Print("Enter Permanent Schedule ID to update: ")
	idStr := util.GetCommandInput()
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		fmt.Println("Invalid ID format")
		util.PressEnterToContinue()
		return err
	}

	schedule, err := h.controller.RetrieveByID(uint(id))
	if err != nil {
		fmt.Println("Failed to retrieve schedule:", err)
		util.PressEnterToContinue()
		return err
	}

	fmt.Println("Current details:")
	fmt.Printf("Course ID: %d, Class ID: %d\n", schedule.CourseId, schedule.ClassId)
	fmt.Printf("Faculty ID: %d, Department ID: %d, Program Type ID: %d\n",
		schedule.FacultyID, schedule.DepartmentID, schedule.ProgramtypeID)
	fmt.Printf("Room ID: %d, Start Date: %v, End Date: %v\n",
		schedule.TimeTable.RoomID,
		schedule.TimeTable.StartDate.Format("2006-01-02 15:04"),
		schedule.TimeTable.EndDate.Format("2006-01-02 15:04"))

	fmt.Println("\nEnter new values (or press Enter to keep current values):")

	fmt.Printf("Current Room ID: %d\n", schedule.TimeTable.RoomID)
	fmt.Print("New Room ID: ")
	input := util.GetCommandInput()
	if input != "" {
		roomID, err := strconv.ParseUint(input, 10, 32)
		if err != nil {
			fmt.Println("Invalid Room ID format")
			util.PressEnterToContinue()
			return err
		}
		schedule.TimeTable.RoomID = uint(roomID)
	}

	fmt.Printf("Current start date and time: %s\n", schedule.TimeTable.StartDate.Format("2006-01-02 15:04"))
	fmt.Print("New start date (YYYY-MM-DD, leave blank to skip): ")
	startDateStr := util.GetCommandInput()
	var newStartDate time.Time

	if startDateStr != "" {
		fmt.Print("New start time (HH:MM): ")
		startTimeStr := util.GetCommandInput()

		startDateTime := startDateStr + " " + startTimeStr
		newStartDate, err = time.Parse("2006-01-02 15:04", startDateTime)
		if err != nil {
			fmt.Println("Invalid date/time format. Please use YYYY-MM-DD for date and HH:MM for time")
			util.PressEnterToContinue()
			return err
		}
		schedule.TimeTable.StartDate = newStartDate
	}

	fmt.Printf("Current end date and time: %s\n", schedule.TimeTable.EndDate.Format("2006-01-02 15:04"))
	fmt.Print("New end date (YYYY-MM-DD, leave blank to skip): ")
	endDateStr := util.GetCommandInput()
	var newEndDate time.Time

	if endDateStr != "" {
		fmt.Print("New end time (HH:MM): ")
		endTimeStr := util.GetCommandInput()

		endDateTime := endDateStr + " " + endTimeStr
		newEndDate, err = time.Parse("2006-01-02 15:04", endDateTime)
		if err != nil {
			fmt.Println("Invalid date/time format. Please use YYYY-MM-DD for date and HH:MM for time")
			util.PressEnterToContinue()
			return err
		}
		schedule.TimeTable.EndDate = newEndDate
	}

	if input != "" || startDateStr != "" || endDateStr != "" {
		isAvailable, err := h.controller.CheckRoomAvailability(
			schedule.TimeTable.RoomID,
			schedule.TimeTable.StartDate,
			schedule.TimeTable.EndDate,
		)
		if err != nil {
			fmt.Println("Error checking room availability:", err)
			util.PressEnterToContinue()
			return err
		}
		if !isAvailable {
			fmt.Println("Room is not available for the specified time period")
			util.PressEnterToContinue()
			return fmt.Errorf("room not available")
		}
	}

	fmt.Printf("Current Course ID: %d\n", schedule.CourseId)
	fmt.Print("New Course ID: ")
	input = util.GetCommandInput()
	if input != "" {
		courseID, err := strconv.ParseUint(input, 10, 32)
		if err != nil {
			fmt.Println("Invalid Course ID format")
			util.PressEnterToContinue()
			return err
		}
		schedule.CourseId = uint(courseID)
	}

	fmt.Printf("Current Class ID: %d\n", schedule.ClassId)
	fmt.Print("New Class ID: ")
	input = util.GetCommandInput()
	if input != "" {
		classID, err := strconv.ParseUint(input, 10, 32)
		if err != nil {
			fmt.Println("Invalid Class ID format")
			util.PressEnterToContinue()
			return err
		}
		schedule.ClassId = uint(classID)
	}

	fmt.Printf("Current Faculty ID: %d\n", schedule.FacultyID)
	fmt.Print("New Faculty ID: ")
	input = util.GetCommandInput()
	if input != "" {
		facultyID, err := strconv.ParseUint(input, 10, 32)
		if err != nil {
			fmt.Println("Invalid Faculty ID format")
			util.PressEnterToContinue()
			return err
		}
		schedule.FacultyID = uint(facultyID)
	}

	fmt.Printf("Current Department ID: %d\n", schedule.DepartmentID)
	fmt.Print("New Department ID: ")
	input = util.GetCommandInput()
	if input != "" {
		deptID, err := strconv.ParseUint(input, 10, 32)
		if err != nil {
			fmt.Println("Invalid Department ID format")
			util.PressEnterToContinue()
			return err
		}
		schedule.DepartmentID = uint(deptID)
	}

	fmt.Printf("Current Program Type ID: %d\n", schedule.ProgramtypeID)
	fmt.Print("New Program Type ID: ")
	input = util.GetCommandInput()
	if input != "" {
		progTypeID, err := strconv.ParseUint(input, 10, 32)
		if err != nil {
			fmt.Println("Invalid Program Type ID format")
			util.PressEnterToContinue()
			return err
		}
		schedule.ProgramtypeID = uint(progTypeID)
	}

	if err := schedule.Validate(); err != nil {
		fmt.Println("Validation error:", err)
		util.PressEnterToContinue()
		return err
	}

	err = h.controller.UpdateByID(schedule)
	if err != nil {
		fmt.Println("Failed to update schedule:", err)
		util.PressEnterToContinue()
		return err
	}

	fmt.Println("Permanent Schedule updated successfully!")
	util.PressEnterToContinue()
	return nil
}
