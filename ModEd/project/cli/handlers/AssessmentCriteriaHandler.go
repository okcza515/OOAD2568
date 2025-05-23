package handlers

import (
	"ModEd/core"
	"ModEd/project/controller"
	"fmt"
)

type AssessmentCriteriaHandler struct {
	menuIO         *core.MenuIO
	instanceStorer *controller.InstanceStorer
}

func NewAssessmentCriteriaHandler(instanceStorer *controller.InstanceStorer) *AssessmentCriteriaHandler {
	return &AssessmentCriteriaHandler{
		menuIO:         core.NewMenuIO(),
		instanceStorer: instanceStorer,
	}
}

func (handler *AssessmentCriteriaHandler) Define(io *core.MenuIO) {
	io.Println("Defining Assessment Criteria...")

	io.Print("Enter Criteria Name (-1 to cancel): ")
	criteriaName, err := io.ReadInput()
	if err != nil || criteriaName == "-1" {
		io.Println("Cancelled.")
		return
	}

	err = handler.instanceStorer.AssessmentCriteria.InsertAssessmentCriteria(criteriaName)
	if err != nil {
		io.Println(fmt.Sprintf("Error adding new criteria: %v", err))
	} else {
		io.Println("Criteria added successfully!")
	}
}

func (handler *AssessmentCriteriaHandler) ListAll(io *core.MenuIO) {
	io.Println("Listing All Assessment Criteria...")

	criteriaList, err := handler.instanceStorer.AssessmentCriteria.List(map[string]interface{}{})
	if err != nil {
		io.Println(fmt.Sprintf("Error listing criteria: %v", err))
		return
	}

	if len(criteriaList) == 0 {
		io.Println("No assessment criteria found.")
		return
	}

	io.PrintTableFromSlice(criteriaList, []string{"ID", "CriteriaName"})
}

func (handler *AssessmentCriteriaHandler) Update(io *core.MenuIO) {
	io.Println("Updating Assessment Criteria...")

	criteriaList, err := handler.instanceStorer.AssessmentCriteria.List(map[string]interface{}{})
	if err != nil {
		io.Println(fmt.Sprintf("Error retrieving criteria list: %v", err))
		return
	}

	for _, c := range criteriaList {
		io.Println(fmt.Sprintf("ID: %v, Name: %v", c.ID, c.CriteriaName))
	}

	io.Print("Enter Criteria ID to update (-1 to cancel): ")
	id, err := io.ReadInputID()
	if err != nil {
		return
	}

	criteria, err := handler.instanceStorer.AssessmentCriteria.RetrieveByID(uint(id))
	if err != nil || criteria == nil {
		io.Println("Criteria not found.")
		return
	}

	io.Print("Enter New Criteria Name (-1 to cancel): ")
	newName, _ := io.ReadInput()
	if newName == "-1" {
		io.Println("Cancelled.")
		return
	}

	criteria.CriteriaName = newName
	err = handler.instanceStorer.AssessmentCriteria.UpdateByID(criteria)
	if err != nil {
		io.Println(fmt.Sprintf("Error updating criteria: %v", err))
	} else {
		io.Println("Criteria updated successfully!")
	}
}

func (handler *AssessmentCriteriaHandler) Delete(io *core.MenuIO) {
	io.Println("Deleting Assessment Criteria...")

	criteriaList, err := handler.instanceStorer.AssessmentCriteria.List(map[string]interface{}{})
	if err != nil {
		io.Println(fmt.Sprintf("Error listing criteria: %v", err))
		return
	}

	for _, c := range criteriaList {
		io.Println(fmt.Sprintf("ID: %v, Name: %v", c.ID, c.CriteriaName))
	}

	io.Print("Enter Criteria ID to delete (-1 to cancel): ")
	id, err := io.ReadInputID()
	if err != nil {
		return
	}

	err = handler.instanceStorer.AssessmentCriteria.DeleteAssessmentCriteria(id)
	if err != nil {
		io.Println(fmt.Sprintf("Error deleting criteria: %v", err))
	} else {
		io.Println("Criteria deleted successfully!")
	}
}
