// MEP-1013
package handler

import (
	model "ModEd/asset/model"
	"ModEd/asset/util"
	"fmt"
)

type ListPermanentSchedulesHandler struct {
	controller interface {
		List(condition map[string]interface{}) ([]model.PermanentSchedule, error)
	}
}

func NewListPermanentSchedulesHandler(controller interface {
	List(condition map[string]interface{}) ([]model.PermanentSchedule, error)
}) *ListPermanentSchedulesHandler {
	return &ListPermanentSchedulesHandler{
		controller: controller,
	}
}

func (h *ListPermanentSchedulesHandler) Execute() error {
	fmt.Println("===== All Permanent Schedules =====")

	schedules, err := h.controller.List(nil)
	if err != nil {
		fmt.Println("Failed to retrieve schedules:", err)
		util.PressEnterToContinue()
		return err
	}

	if len(schedules) == 0 {
		fmt.Println("No permanent schedules found.")
	} else {
		fmt.Printf("Total schedules: %d\n\n", len(schedules))
		fmt.Println("=======================================================")
		fmt.Println("ID   | Course | Class | Faculty | Department | Program")
		fmt.Println("-------------------------------------------------------")
		for _, schedule := range schedules {
			fmt.Printf("%-4d | %-6d | %-5d | %-7d | %-10d | %-7d\n",
				schedule.ID, schedule.CourseId, schedule.ClassId,
				schedule.FacultyID, schedule.DepartmentID, schedule.ProgramtypeID)
		}
		fmt.Println("=======================================================")
	}

	util.PressEnterToContinue()
	return nil
}
