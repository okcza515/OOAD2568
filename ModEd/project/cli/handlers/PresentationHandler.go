package handlers

import (
	"ModEd/core"
	"ModEd/project/controller"
	"ModEd/project/model"
	"fmt"
)

type PresentationHandler struct {
	menuIO         *core.MenuIO
	instanceStorer *controller.InstanceStorer
}

func NewPresentationHandler(instanceStorer *controller.InstanceStorer) *PresentationHandler {
	return &PresentationHandler{
		menuIO:         core.NewMenuIO(),
		instanceStorer: instanceStorer,
	}
}

func (h *PresentationHandler) ViewAll(io *core.MenuIO) {
	io.Println("Viewing Presentations...")

	presentations, err := h.instanceStorer.Presentation.ListAllPresentations()
	if err != nil {
		io.Println(fmt.Sprintf("Error retrieving presentations: %v", err))
		return
	}

	if len(presentations) == 0 {
		io.Println("No presentations found.")
		return
	}

	io.Println("Presentations (Based on Date):")
	io.PrintTableFromSlice(presentations, []string{"ID", "SeniorProjectId", "PresentationType", "Date"})
}

func (h *PresentationHandler) Add(io *core.MenuIO) {
	io.Println("Adding New Presentation...")

	io.Print("Enter Senior Project ID: ")
	projectID, err := io.ReadInputID()
	if err != nil {
		return
	}

	io.Print("Enter Presentation Type (Proposal, Midterm, Final): ")
	presentationTypeInput, err := io.ReadInput()
	if err != nil {
		io.Println(fmt.Sprintf("Error reading input: %v", err))
		return
	}

	presentationType := model.PresentationType(presentationTypeInput)
	if !presentationType.IsValid() {
		io.Println("Invalid presentation type.")
		return
	}

	io.Print("Enter Date (YYYY-MM-DD): ")
	dueDate, err := io.ReadInputTime()
	if err != nil {
		io.Println(fmt.Sprintf("Invalid Date format: %v", err))
		return
	}

	_, err = h.instanceStorer.Presentation.InsertPresentation(projectID, presentationType, dueDate)
	if err != nil {
		io.Println(fmt.Sprintf("Error inserting presentation: %v", err))
	} else {
		io.Println("Presentation added successfully!")
	}
}

func (h *PresentationHandler) ViewByID(io *core.MenuIO) {
	io.Println("Viewing Presentation by ID...")

	io.Print("Enter Presentation ID: ")
	presentationID, err := io.ReadInputID()
	if err != nil {
		return
	}

	presentation, err := h.instanceStorer.Presentation.RetrievePresentation(presentationID)
	if err != nil {
		io.Println(fmt.Sprintf("Error retrieving presentation: %v", err))
		return
	}

	io.Println(fmt.Sprintf("Presentation ID: %d\nProject ID: %d\nType: %s\nDate: %s",
		presentation.ID, presentation.SeniorProjectId, presentation.PresentationType, presentation.Date.Format("2006-01-02")))
}

func (h *PresentationHandler) Update(io *core.MenuIO) {
	io.Println("Updating Presentation...")

	io.Print("Enter Presentation ID to update: ")
	presentationID, err := io.ReadInputID()
	if err != nil {
		return
	}

	presentation, err := h.instanceStorer.Presentation.RetrievePresentation(presentationID)
	if err != nil {
		io.Println(fmt.Sprintf("Error retrieving presentation: %v", err))
		return
	}

	io.Println(fmt.Sprintf("Current Type (%s): ", presentation.PresentationType))
	newTypeInput, _ := io.ReadInput()
	if newTypeInput != "" {
		newType := model.PresentationType(newTypeInput)
		if !newType.IsValid() {
			io.Println("Invalid Type.")
			return
		}
		presentation.PresentationType = newType
	}

	io.Println(fmt.Sprintf("Current Date (%s): ", presentation.Date.Format("2006-01-02")))
	newDateInput, err := io.ReadInputTime()
	if err != nil {
		return
	}
	presentation.Date = newDateInput

	err = h.instanceStorer.Presentation.UpdatePresentation(presentation)
	if err != nil {
		io.Println(fmt.Sprintf("Error updating presentation: %v", err))
	} else {
		io.Println("Presentation updated successfully!")
	}
}

func (h *PresentationHandler) Delete(io *core.MenuIO) {
	io.Println("Deleting Presentation...")

	io.Print("Enter Presentation ID to delete: ")
	presentationID, err := io.ReadInputID()
	if err != nil {
		return
	}

	err = h.instanceStorer.Presentation.DeletePresentation(presentationID)
	if err != nil {
		io.Println(fmt.Sprintf("Error deleting presentation: %v", err))
	} else {
		io.Println("Presentation deleted successfully!")
	}
}
