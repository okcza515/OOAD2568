//MEP-1013

package handler

import (
	"ModEd/asset/model"
	"ModEd/asset/util"
	"fmt"
	"strconv"
)

type UpdateScheduleHandler struct {
	controller interface {
		RetrieveByID(id uint) (model.PermanentSchedule, error)
		UpdateByID(schedule model.PermanentSchedule) error
	}
}

func NewUpdateScheduleHandler(controller interface {
	RetrieveByID(id uint) (model.PermanentSchedule, error)
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

	fmt.Println("\nEnter new values (or press Enter to keep current values):")

	fmt.Printf("Current Course ID: %d\n", schedule.CourseId)
	fmt.Print("New Course ID: ")
	input := util.GetCommandInput()
	if input != "" {
		val, err := strconv.ParseUint(input, 10, 32)
		if err == nil {
			schedule.CourseId = uint(val)
		} else {
			fmt.Println("Invalid input - keeping current value")
		}
	}

	fmt.Printf("Current Class ID: %d\n", schedule.ClassId)
	fmt.Print("New Class ID: ")
	input = util.GetCommandInput()
	if input != "" {
		val, err := strconv.ParseUint(input, 10, 32)
		if err == nil {
			schedule.ClassId = uint(val)
		} else {
			fmt.Println("Invalid input - keeping current value")
		}
	}

	fmt.Printf("Current Faculty ID: %d\n", schedule.FacultyID)
	fmt.Print("New Faculty ID: ")
	input = util.GetCommandInput()
	if input != "" {
		val, err := strconv.ParseUint(input, 10, 32)
		if err == nil {
			schedule.FacultyID = uint(val)
		} else {
			fmt.Println("Invalid input - keeping current value")
		}
	}

	fmt.Printf("Current Department ID: %d\n", schedule.DepartmentID)
	fmt.Print("New Department ID: ")
	input = util.GetCommandInput()
	if input != "" {
		val, err := strconv.ParseUint(input, 10, 32)
		if err == nil {
			schedule.DepartmentID = uint(val)
		} else {
			fmt.Println("Invalid input - keeping current value")
		}
	}

	fmt.Printf("Current Program Type ID: %d\n", schedule.ProgramtypeID)
	fmt.Print("New Program Type ID: ")
	input = util.GetCommandInput()
	if input != "" {
		val, err := strconv.ParseUint(input, 10, 32)
		if err == nil {
			schedule.ProgramtypeID = uint(val)
		} else {
			fmt.Println("Invalid input - keeping current value")
		}
	}

	err = h.controller.UpdateByID(schedule)
	if err != nil {
		fmt.Println("Failed to update schedule:", err)
	} else {
		fmt.Println("Schedule updated successfully!")
	}

	util.PressEnterToContinue()
	return nil
}
