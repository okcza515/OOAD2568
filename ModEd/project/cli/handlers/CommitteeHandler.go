package handlers

import (
	"ModEd/core"
	"ModEd/project/controller"
	"ModEd/project/model"
	"fmt"
)

type CommitteeHandler struct {
	menuIO         *core.MenuIO
	instanceStorer *controller.InstanceStorer
}

func NewCommitteeHandler(instanceStorer *controller.InstanceStorer) *CommitteeHandler {
	return &CommitteeHandler{
		menuIO:         core.NewMenuIO(),
		instanceStorer: instanceStorer,
	}
}

func (h *CommitteeHandler) Add(io *core.MenuIO) {
	io.Println("Adding Committee Member...")

	io.Print("Enter Project ID (-1 to cancel): ")
	projectId, err := io.ReadInputID()
	if err != nil {
		return
	}

	io.Print("Enter Instructor ID (-1 to cancel): ")
	instructorId, err := io.ReadInputID()
	if err != nil {
		return
	}

	committee := &model.Committee{
		SeniorProjectId: projectId,
		InstructorId:    instructorId,
	}

	err = h.instanceStorer.Committee.InsertCommittee(committee)
	if err != nil {
		io.Println(fmt.Sprintf("Error adding committee member: %v", err))
	} else {
		io.Println("Committee member added successfully!")
	}
}

func (h *CommitteeHandler) ListByProject(io *core.MenuIO) {
	io.Println("Listing Committee Members by Project...")

	io.Print("Enter Project ID (-1 to cancel): ")
	projectId, err := io.ReadInputID()
	if err != nil {
		return
	}

	committees, err := h.instanceStorer.Committee.ListCommitteesByProject(projectId)
	if err != nil {
		io.Println(fmt.Sprintf("Error listing committee members: %v", err))
		return
	}

	if len(committees) == 0 {
		io.Println("No committee members found for this project.")
		return
	}

	io.Println(fmt.Sprintf("Committee Members for Project ID %v:", projectId))
	for _, c := range committees {
		io.Println(fmt.Sprintf("Committee ID: %v, Instructor ID: %v", c.ID, c.InstructorId))
	}
}

func (h *CommitteeHandler) ListProjectsByInstructor(io *core.MenuIO) {
	io.Println("Listing Projects by Committee Member...")

	io.Print("Enter Instructor ID (-1 to cancel): ")
	instructorId, err := io.ReadInputID()
	if err != nil {
		return
	}

	committees, err := h.instanceStorer.Committee.ListCommitteesByInstructor(instructorId)
	if err != nil {
		io.Println(fmt.Sprintf("Error listing projects: %v", err))
		return
	}

	if len(committees) == 0 {
		io.Println("No projects found for this committee member.")
		return
	}

	io.Println(fmt.Sprintf("Projects for Instructor ID %v as Committee Member:", instructorId))
	io.PrintTableFromSlice(committees, []string{"ID", "SeniorProjectId"})
}

func (h *CommitteeHandler) Remove(io *core.MenuIO) {
	io.Println("Removing Committee Member...")

	io.Print("Enter Committee ID (-1 to cancel): ")
	committeeId, err := io.ReadInputID()
	if err != nil {
		return
	}

	err = h.instanceStorer.Committee.DeleteByID(committeeId)
	if err != nil {
		io.Println(fmt.Sprintf("Error removing committee member: %v", err))
	} else {
		io.Println("Committee member removed successfully!")
	}
}
