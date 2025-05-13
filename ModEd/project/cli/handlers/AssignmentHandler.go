package handlers

import (
	"ModEd/core"
	"ModEd/project/controller"
	"fmt"
)

type AssignmentHandler struct {
	menuIO         *core.MenuIO
	instanceStorer *controller.InstanceStorer
}

func NewAssignmentHandler(instanceStorer *controller.InstanceStorer) *AssignmentHandler {
	return &AssignmentHandler{
		menuIO:         core.NewMenuIO(),
		instanceStorer: instanceStorer,
	}
}

func (h *AssignmentHandler) ViewAll(io *core.MenuIO) {
	assignments, err := h.instanceStorer.Assignment.List(map[string]interface{}{})
	if err != nil {
		io.Println(fmt.Sprintf("Error retrieving assignments: %v", err))
		return
	}
	if len(assignments) == 0 {
		io.Println("No assignments found.")
		return
	}
	io.PrintTableFromSlice(assignments, []string{"ID", "SeniorProjectId", "Name", "DueDate"})
}

func (h *AssignmentHandler) Add(io *core.MenuIO) {
	io.Print("Enter Senior Project ID: ")
	projectID, err := io.ReadInputID()
	if err != nil {
		return
	}

	io.Print("Enter Assignment Name: ")
	name, _ := io.ReadInput()

	io.Print("Enter Assignment Description: ")
	description, _ := io.ReadInput()

	io.Print("Enter Due Date (YYYY-MM-DD): ")
	dueDate, err := io.ReadInputTime()
	if err != nil {
		io.Println("Invalid date format.")
		return
	}

	_, err = h.instanceStorer.Assignment.InsertAssignment(projectID, name, description, dueDate)
	if err != nil {
		io.Println(fmt.Sprintf("Error adding assignment: %v", err))
	} else {
		io.Println("Assignment added successfully.")
	}
}

func (h *AssignmentHandler) Delete(io *core.MenuIO) {
	io.Print("Enter Assignment ID to delete: ")
	id, err := io.ReadInputID()
	if err != nil {
		return
	}

	err = h.instanceStorer.Assignment.DeleteAssignment(id)
	if err != nil {
		io.Println(fmt.Sprintf("Error deleting assignment: %v", err))
	} else {
		io.Println("Assignment deleted successfully.")
	}
}
