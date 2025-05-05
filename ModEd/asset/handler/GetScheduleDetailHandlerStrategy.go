//MEP-1013

package handler

import (
	model "ModEd/asset/model"
	"ModEd/asset/util"
	"fmt"
	"strconv"
)

type GetScheduleDetailsHandler struct {
	controller interface {
		RetrieveByID(id uint) (model.PermanentSchedule, error)
	}
}

func NewGetScheduleDetailsHandler(controller interface {
	RetrieveByID(id uint) (model.PermanentSchedule, error)
}) *GetScheduleDetailsHandler {
	return &GetScheduleDetailsHandler{
		controller: controller,
	}
}

func (h *GetScheduleDetailsHandler) Execute() error {
	fmt.Println("===== Get Schedule Details =====")

	fmt.Print("Enter Schedule ID: ")
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

	fmt.Println("Schedule Details:")
	fmt.Println("=======================================================")
	fmt.Printf("Schedule ID:    %d\n", schedule.ID)
	fmt.Printf("Course ID:      %d\n", schedule.CourseId)
	fmt.Printf("Class ID:       %d\n", schedule.ClassId)
	fmt.Printf("Faculty ID:     %d\n", schedule.FacultyID)
	fmt.Printf("Department ID:  %d\n", schedule.DepartmentID)
	fmt.Printf("Program Type:   %d\n", schedule.ProgramtypeID)

	if schedule.TimeTable.ID > 0 {
		fmt.Println("\nTime Table Details:")
		fmt.Printf("Time Table ID:   %d\n", schedule.TimeTable.ID)
		fmt.Printf("Room ID:         %d\n", schedule.TimeTable.RoomID)
		fmt.Printf("Start Date/Time: %s\n", schedule.TimeTable.StartDate.Format("2006-01-02 15:04"))
		fmt.Printf("End Date/Time:   %s\n", schedule.TimeTable.EndDate.Format("2006-01-02 15:04"))
		fmt.Printf("Booking Type:    %s\n", schedule.TimeTable.BookingType)
		fmt.Printf("Is Available:    %t\n", schedule.TimeTable.IsAvailable)
	} else {
		fmt.Println("\nNo time table information available")
	}
	fmt.Println("=======================================================")

	util.PressEnterToContinue()
	return nil
}
