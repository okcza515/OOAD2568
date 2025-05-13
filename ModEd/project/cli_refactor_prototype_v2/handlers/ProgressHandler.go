package handlers

import (
	"ModEd/core"
	"ModEd/project/controller"
	"fmt"
)

type ProgressHandler struct {
	menuIO         *core.MenuIO
	instanceStorer *controller.InstanceStorer
}

func NewProgressHandler(instanceStorer *controller.InstanceStorer) *ProgressHandler {
	return &ProgressHandler{
		menuIO:         core.NewMenuIO(),
		instanceStorer: instanceStorer,
	}
}

func (h *ProgressHandler) ViewAll(io *core.MenuIO) {
	io.Println("Viewing All Progress...")

	progressList, err := h.instanceStorer.Progress.GetFormattedProgressList()
	if err != nil {
		io.Println(fmt.Sprintf("Error retrieving progress: %v", err))
		return
	}

	if len(progressList) == 0 {
		io.Println("No progress found.")
		return
	}

	io.Println("Progress List:")
	for _, line := range progressList {
		io.Println(line)
	}
}

func (h *ProgressHandler) Add(io *core.MenuIO) {
	io.Println("Adding New Progress...")

	io.Print("Enter Assignment ID: ")
	assignmentID, err := io.ReadInputID()
	if err != nil {
		return
	}

	io.Print("Enter Progress Name: ")
	name, err := io.ReadInput()
	if err != nil {
		io.Println(fmt.Sprintf("Error reading input: %v", err))
		return
	}

	err = h.instanceStorer.Progress.AddNewProgress(uint(assignmentID), name)
	if err != nil {
		io.Println(fmt.Sprintf("Error adding new progress: %v", err))
	} else {
		io.Println("Progress added successfully!")
	}
}

func (h *ProgressHandler) ViewByID(io *core.MenuIO) {
	io.Println("Viewing Progress by ID...")

	io.Print("Enter Progress ID: ")
	progressID, err := io.ReadInputID()
	if err != nil {
		return
	}

	progress, err := h.instanceStorer.Progress.RetrieveByID(progressID)
	if err != nil {
		io.Println(fmt.Sprintf("Error retrieving progress: %v", err))
		return
	}

	io.Println(fmt.Sprintf("Progress ID: %d, Assignment ID: %d, Name: %s, Completed: %t",
		progress.ID, progress.AssignmentId, progress.Name, progress.IsCompleted))
}

func (h *ProgressHandler) UpdateName(io *core.MenuIO) {
	io.Println("Updating Progress Name...")

	io.Print("Enter Progress ID: ")
	progressID, err := io.ReadInputID()
	if err != nil {
		return
	}

	io.Print("Enter New Progress Name: ")
	newName, err := io.ReadInput()
	if err != nil {
		io.Println(fmt.Sprintf("Error reading input: %v", err))
		return
	}

	err = h.instanceStorer.Progress.UpdateProgressName(progressID, newName)
	if err != nil {
		io.Println(fmt.Sprintf("Error updating progress name: %v", err))
	} else {
		io.Println("Progress name updated successfully!")
	}
}

func (h *ProgressHandler) Delete(io *core.MenuIO) {
	io.Println("Deleting Progress...")

	io.Print("Enter Progress ID to delete: ")
	progressID, err := io.ReadInputID()
	if err != nil {
		return
	}

	err = h.instanceStorer.Progress.DeleteByID(uint(progressID))
	if err != nil {
		io.Println(fmt.Sprintf("Error deleting progress: %v", err))
	} else {
		io.Println("Progress deleted successfully!")
	}
}

func (h *ProgressHandler) MarkCompleted(io *core.MenuIO) {
	io.Println("Marking Progress as Completed...")

	io.Print("Enter Progress ID: ")
	progressID, err := io.ReadInputID()
	if err != nil {
		return
	}

	err = h.instanceStorer.Progress.MarkAsCompleted(uint(progressID))
	if err != nil {
		io.Println(fmt.Sprintf("Error marking progress as completed: %v", err))
	} else {
		io.Println("Progress marked as completed successfully!")
	}
}

func (h *ProgressHandler) MarkIncomplete(io *core.MenuIO) {
	io.Println("Marking Progress as Incomplete...")

	io.Print("Enter Progress ID: ")
	progressID, err := io.ReadInputID()
	if err != nil {
		return
	}

	err = h.instanceStorer.Progress.MarkProgressAsIncomplete(uint(progressID))
	if err != nil {
		io.Println(fmt.Sprintf("Error marking progress as incomplete: %v", err))
	} else {
		io.Println("Progress marked as incomplete successfully!")
	}
}
