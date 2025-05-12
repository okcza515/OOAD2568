package handlers

import (
	"ModEd/core"
	"ModEd/project/controller"
	"ModEd/project/model"
	"fmt"
)

type SeniorProjectHandler struct {
	menuIO         *core.MenuIO
	instanceStorer *controller.InstanceStorer
}

func NewSeniorProjectHandler(instanceStorer *controller.InstanceStorer) *SeniorProjectHandler {
	return &SeniorProjectHandler{
		menuIO:         core.NewMenuIO(),
		instanceStorer: instanceStorer,
	}
}

func (h *SeniorProjectHandler) Create(io *core.MenuIO) {
	io.Print("Enter the group name (-1 to cancel): ")
	groupNameStr, err := io.ReadInput()
	if err != nil || groupNameStr == "-1" {
		io.Println("Cancelled.")
		return
	}

	err = h.instanceStorer.SeniorProject.Insert(&model.SeniorProject{
		GroupName: groupNameStr,
	})
	if err != nil {
		io.Println(fmt.Sprintf("Error creating senior project: %v", err))
		return
	}

	io.Println("Senior project created successfully!")
}

func (h *SeniorProjectHandler) List(io *core.MenuIO) {
	records, err := h.instanceStorer.SeniorProject.List(map[string]interface{}{})
	if err != nil {
		io.Println(fmt.Sprintf("Error listing senior projects: %v", err))
		return
	}

	if len(records) == 0 {
		io.Println("No senior projects found.")
		return
	}

	io.PrintTableFromSlice(records, []string{"ID", "GroupName", "CreatedAt"})
}
