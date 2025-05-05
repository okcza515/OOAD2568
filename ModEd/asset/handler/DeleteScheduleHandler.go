// MEP-1013
package handler

import (
	"ModEd/asset/model"
	"ModEd/asset/util"
	"fmt"
	"strconv"
	"strings"
)

type DeleteScheduleHandler struct {
	controller interface {
		RetrieveByID(id uint, preload ...string) (model.PermanentSchedule, error)
		DeleteByID(id uint) error
	}
}

func NewDeleteScheduleHandler(controller interface {
	RetrieveByID(id uint, preload ...string) (model.PermanentSchedule, error)
	DeleteByID(id uint) error
}) *DeleteScheduleHandler {
	return &DeleteScheduleHandler{
		controller: controller,
	}
}

func (h *DeleteScheduleHandler) Execute() error {
	fmt.Println("===== Delete Schedule =====")

	fmt.Print("Enter Schedule ID to delete: ")
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

	fmt.Println("Schedule to delete:")
	fmt.Printf("ID: %d, Course: %d, Class: %d\n",
		schedule.ID, schedule.CourseId, schedule.ClassId)

	fmt.Print("Are you sure you want to delete this schedule? (yes/no): ")
	confirmation := util.GetCommandInput()

	if strings.ToLower(confirmation) == "yes" {
		err = h.controller.DeleteByID(uint(id))
		if err != nil {
			fmt.Println("Failed to delete schedule:", err)
		} else {
			fmt.Println("Schedule deleted successfully!")
		}
	} else {
		fmt.Println("Deletion cancelled")
	}

	util.PressEnterToContinue()
	return nil
}
