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

func NewUpdatePermanentScheduleHandler(controller interface {
	RetrieveByID(id uint, preload ...string) (model.PermanentSchedule, error)
	UpdateByID(schedule model.PermanentSchedule) error
}) *UpdateScheduleHandler {
	return &UpdateScheduleHandler{
		controller: controller,
	}
}

func (handler *UpdateScheduleHandler) Execute() error {
	fmt.Println("------- Update Permanent Schedule -------")

	fmt.Println("Please enter the ID of the Permanent Schedule:")
	scheduleIdStr := util.GetCommandInput()
	scheduleId, err := strconv.Atoi(scheduleIdStr)
	if err != nil {
		fmt.Println("Invalid ID")
		util.PressEnterToContinue()
		return err
	}

	schedule, err := handler.controller.RetrieveByID(uint(scheduleId))
	if err != nil {
		fmt.Println("Failed to retrieve schedule:", err)
		util.PressEnterToContinue()
		return err
	}

	fmt.Printf("Current Room ID: %d\n", schedule.TimeTable.RoomID)
	fmt.Println("Please enter the new Room ID (or press Enter to keep current):")
	roomIDStr := util.GetCommandInput()
	if roomIDStr != "" {
		roomID, err := strconv.Atoi(roomIDStr)
		if err != nil {
			fmt.Println("Invalid Room ID")
			util.PressEnterToContinue()
			return err
		}
		schedule.TimeTable.RoomID = uint(roomID)
	}

	fmt.Printf("Current Start Date: %s\n", schedule.TimeTable.StartDate.Format("2006-01-02 15:04"))
	fmt.Println("Please enter the new start date (YYYY-MM-DD, or press Enter to keep current):")
	startDateStr := util.GetCommandInput()
	if startDateStr != "" {
		fmt.Println("Please enter the new start time (HH:MM):")
		startTimeStr := util.GetCommandInput()

		fullStartDateTime := startDateStr + " " + startTimeStr
		startDate, err := time.Parse("2006-01-02 15:04", fullStartDateTime)
		if err != nil {
			fmt.Println("Invalid date/time format")
			util.PressEnterToContinue()
			return err
		}
		schedule.TimeTable.StartDate = startDate
	}

	fmt.Printf("Current End Date: %s\n", schedule.TimeTable.EndDate.Format("2006-01-02 15:04"))
	fmt.Println("Please enter the new end date (YYYY-MM-DD, or press Enter to keep current):")
	endDateStr := util.GetCommandInput()
	if endDateStr != "" {
		fmt.Println("Please enter the new end time (HH:MM):")
		endTimeStr := util.GetCommandInput()

		fullEndDateTime := endDateStr + " " + endTimeStr
		endDate, err := time.Parse("2006-01-02 15:04", fullEndDateTime)
		if err != nil {
			fmt.Println("Invalid date/time format")
			util.PressEnterToContinue()
			return err
		}
		schedule.TimeTable.EndDate = endDate
	}

	fmt.Printf("Current Course ID: %d\n", schedule.CourseId)
	fmt.Println("Please enter the new Course ID (or press Enter to keep current):")
	courseIDStr := util.GetCommandInput()
	if courseIDStr != "" {
		courseID, err := strconv.Atoi(courseIDStr)
		if err != nil {
			fmt.Println("Invalid Course ID")
			util.PressEnterToContinue()
			return err
		}
		schedule.CourseId = uint(courseID)
	}

	fmt.Printf("Current Class ID: %d\n", schedule.ClassId)
	fmt.Println("Please enter the new Class ID (or press Enter to keep current):")
	classIDStr := util.GetCommandInput()
	if classIDStr != "" {
		classID, err := strconv.Atoi(classIDStr)
		if err != nil {
			fmt.Println("Invalid Class ID")
			util.PressEnterToContinue()
			return err
		}
		schedule.ClassId = uint(classID)
	}

	fmt.Printf("Current Faculty ID: %d\n", schedule.FacultyID)
	fmt.Println("Please enter the new Faculty ID (or press Enter to keep current):")
	facultyIDStr := util.GetCommandInput()
	if facultyIDStr != "" {
		facultyID, err := strconv.Atoi(facultyIDStr)
		if err != nil {
			fmt.Println("Invalid Faculty ID")
			util.PressEnterToContinue()
			return err
		}
		schedule.FacultyID = uint(facultyID)
	}

	fmt.Printf("Current Department ID: %d\n", schedule.DepartmentID)
	fmt.Println("Please enter the new Department ID (or press Enter to keep current):")
	deptIDStr := util.GetCommandInput()
	if deptIDStr != "" {
		deptID, err := strconv.Atoi(deptIDStr)
		if err != nil {
			fmt.Println("Invalid Department ID")
			util.PressEnterToContinue()
			return err
		}
		schedule.DepartmentID = uint(deptID)
	}

	fmt.Printf("Current Program Type ID: %d\n", schedule.ProgramtypeID)
	fmt.Println("Please enter the new Program Type ID (0 for Regular, 1 for International, or press Enter to keep current):")
	progTypeIDStr := util.GetCommandInput()
	if progTypeIDStr != "" {
		progTypeID, err := strconv.Atoi(progTypeIDStr)
		if err != nil {
			fmt.Println("Invalid Program Type ID")
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

	fmt.Println("-------- Updated Permanent Schedule Details --------")
	fmt.Println("Room ID:", schedule.TimeTable.RoomID)
	fmt.Println("Start Date:", schedule.TimeTable.StartDate.Format("2006-01-02 15:04"))
	fmt.Println("End Date:", schedule.TimeTable.EndDate.Format("2006-01-02 15:04"))
	fmt.Println("Course ID:", schedule.CourseId)
	fmt.Println("Class ID:", schedule.ClassId)
	fmt.Println("Faculty ID:", schedule.FacultyID)
	fmt.Println("Department ID:", schedule.DepartmentID)
	fmt.Println("Program Type ID:", schedule.ProgramtypeID)

	fmt.Println("\nDo you want to update this Permanent Schedule? (y/n)")
	confirmStr := util.GetCommandInput()
	if confirmStr != "y" {
		fmt.Println("Permanent Schedule update cancelled.")
		util.PressEnterToContinue()
		return nil
	}

	err = handler.controller.UpdateByID(schedule)
	if err != nil {
		fmt.Println("Failed to update Permanent Schedule:", err)
		util.PressEnterToContinue()
		return err
	}

	fmt.Println("Permanent Schedule updated successfully!")
	util.PressEnterToContinue()
	return nil
}
