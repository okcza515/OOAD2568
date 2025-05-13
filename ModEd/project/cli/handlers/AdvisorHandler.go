package handlers

import (
	"ModEd/core"
	"ModEd/project/controller"
	"fmt"
)

type AdvisorHandler struct {
	menuIO         *core.MenuIO
	instanceStorer *controller.InstanceStorer
}

func NewAdvisorHandler(instanceStorer *controller.InstanceStorer) *AdvisorHandler {
	return &AdvisorHandler{
		menuIO:         core.NewMenuIO(),
		instanceStorer: instanceStorer,
	}
}

func (handler *AdvisorHandler) AssignAdvisor(io *core.MenuIO) {
	io.Println("Assigning Advisor to Project...")

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

	io.Print("Is this a primary advisor? (yes/no): ")
	isPrimary, err := io.ReadYesOrNo()
	if err != nil {
		return
	}

	advisor, err := handler.instanceStorer.Advisor.AssignAdvisor(projectId, instructorId, isPrimary)
	if err != nil {
		io.Println(fmt.Sprintf("Error assigning advisor: %v", err))
	} else {
		io.Println(fmt.Sprintf("Advisor assigned successfully! Advisor ID: %v", advisor.ID))
	}
}

func (handler *AdvisorHandler) UpdateAdvisorRole(io *core.MenuIO) {
	io.Println("Updating Advisor Role...")

	io.Print("Enter Advisor ID (-1 to cancel): ")
	advisorId, err := io.ReadInputID()
	if err != nil {
		return
	}

	io.Print("Set as primary advisor? (yes/no): ")
	isPrimary, err := io.ReadYesOrNo()
	if err != nil {
		return
	}

	err = handler.instanceStorer.Advisor.UpdateAdvisorRole(uint(advisorId), isPrimary)
	if err != nil {
		io.Println(fmt.Sprintf("Error updating advisor role: %v", err))
	} else {
		io.Println("Advisor role updated successfully!")
	}
}

func (handler *AdvisorHandler) RemoveAdvisor(io *core.MenuIO) {
	io.Println("Removing Advisor...")

	io.Print("Enter Advisor ID (-1 to cancel): ")
	advisorId, err := io.ReadInputID()
	if err != nil {
		return
	}

	err = handler.instanceStorer.Advisor.DeleteByID(advisorId)
	if err != nil {
		io.Println(fmt.Sprintf("Error removing advisor: %v", err))
	} else {
		io.Println("Advisor removed successfully!")
	}
}

func (handler *AdvisorHandler) ListAdvisorsByProject(io *core.MenuIO) {
	io.Println("Listing Advisors by Project...")

	io.Print("Enter Project ID (-1 to cancel): ")
	projectId, err := io.ReadInputID()
	if err != nil {
		return
	}

	advisors, err := handler.instanceStorer.Advisor.List(map[string]interface{}{
		"senior_project_id": projectId,
	})
	if err != nil {
		io.Println(fmt.Sprintf("Error listing advisors: %v", err))
		return
	}

	if len(advisors) == 0 {
		io.Println("No advisors found for this project.")
		return
	}

	io.Println(fmt.Sprintf("Advisors for Project ID %v:", projectId))
	io.PrintTableFromSlice(advisors, []string{"ID", "IsPrimary", "InstructorId", "CreatedAt"})
}

func (handler *AdvisorHandler) ListProjectsByInstructor(io *core.MenuIO) {
	io.Println("Listing Projects by Instructor...")

	io.Print("Enter Instructor ID (-1 to cancel): ")
	instructorId, err := io.ReadInputID()
	if err != nil {
		return
	}

	advisors, err := handler.instanceStorer.Advisor.List(map[string]interface{}{
		"instructorId": instructorId,
	})
	if err != nil {
		io.Println(fmt.Sprintf("Error listing projects: %v", err))
		return
	}

	if len(advisors) == 0 {
		io.Println("No projects found for this instructor.")
		return
	}

	io.Println(fmt.Sprintf("Projects for Instructor ID %v:", instructorId))
	io.PrintTableFromSlice(advisors, []string{"ID", "IsPrimary", "SeniorProjectId", "CreatedAt"})
}
