// MEP-1013

package handler

import (
	"ModEd/asset/model"
	"ModEd/asset/util"
	"fmt"
	"strconv"
	"time"
)

type UpdateScheduleHandler struct {
	controller interface {
		RetrieveByID(id uint, preload ...string) (model.PermanentSchedule, error)
		UpdateByID(schedule model.PermanentSchedule) error
	}
}

func NewUpdateScheduleHandler(controller interface {
	RetrieveByID(id uint, preload ...string) (model.PermanentSchedule, error)
	UpdateByID(schedule model.PermanentSchedule) error
}) *UpdateScheduleHandler {
	return &UpdateScheduleHandler{
		controller: controller,
	}
}

func (h *UpdateScheduleHandler) Execute() error {
	fmt.Println("===== Update Schedule =====")

	fmt.Print("Enter Schedule ID to update: ")
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

	fmt.Print("New start date (YYYY-MM-DD): ")
	startDateStr := util.GetCommandInput()

	if startDateStr != "" {
		fmt.Print("New start time (HH:MM): ")
		startTimeStr := util.GetCommandInput()

		startDateTime := startDateStr + " " + startTimeStr
		startDate, err := time.Parse("2006-01-02 15:04", startDateTime)
		if err != nil {
			fmt.Println("Invalid date/time format. Please use YYYY-MM-DD for date and HH:MM for time")
			util.PressEnterToContinue()
			return err
		}
		schedule.TimeTable.StartDate = startDate
	}

	fmt.Printf("Current end date and time: %s\n", schedule.TimeTable.EndDate.Format("2006-01-02 15:04"))

	fmt.Print("New end date (YYYY-MM-DD): ")
	endDateStr := util.GetCommandInput()

	if endDateStr != "" {
		fmt.Print("New end time (HH:MM): ")
		endTimeStr := util.GetCommandInput()

		endDateTime := endDateStr + " " + endTimeStr
		endDate, err := time.Parse("2006-01-02 15:04", endDateTime)
		if err != nil {
			fmt.Println("Invalid date/time format. Please use YYYY-MM-DD for date and HH:MM for time")
			util.PressEnterToContinue()
			return err
		}
		schedule.TimeTable.EndDate = endDate
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

	err = h.controller.UpdateByID(schedule)
	if err != nil {
		fmt.Println("Failed to update schedule:", err)
		util.PressEnterToContinue()
		return err
	}

	fmt.Println("Schedule updated successfully!")
	util.PressEnterToContinue()
	return nil
}
