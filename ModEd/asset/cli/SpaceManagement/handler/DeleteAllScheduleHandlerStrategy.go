// MEP-1013

package handler

import (
	"ModEd/asset/util"
	"fmt"
	"strings"
)

type DeleteAllSchedulesHandler struct {
	controller interface {
		DeleteAll() error
	}
}

func NewDeleteAllSchedulesHandler(controller interface {
	DeleteAll() error
}) *DeleteAllSchedulesHandler {
	return &DeleteAllSchedulesHandler{
		controller: controller,
	}
}

func (h *DeleteAllSchedulesHandler) Execute() error {
	fmt.Println("===== Delete All Schedules =====")

	fmt.Print("Are you sure you want to delete ALL schedules? This cannot be undone! (yes/no): ")
	confirmation := util.GetCommandInput()

	if strings.ToLower(confirmation) == "yes" {
		err := h.controller.DeleteAll()
		if err != nil {
			fmt.Println("Failed to delete all schedules:", err)
		} else {
			fmt.Println("Successfully deleted all schedules")
		}
	} else {
		fmt.Println("Deletion cancelled")
	}

	util.PressEnterToContinue()
	return nil
}
