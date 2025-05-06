package main

import (
	"ModEd/project/controller"
	"ModEd/project/utils"
	"fmt"
	"strconv"
)

func BuildAssessmentCriteriaManagerMenu(
	assessmentCriteriaController *controller.AssessmentCriteriaController,
) *utils.MenuItem {
	return &utils.MenuItem{
		Title: "Assessment Criteria Manager",
		Children: []*utils.MenuItem{
			{
				Title:  "Define Assessment Criteria",
				Action: defineAssessmentCriteria(assessmentCriteriaController),
			},
			{
				Title:  "List All Assessment Criteria",
				Action: listAllAssessmentCriteria(assessmentCriteriaController),
			},
			{
				Title:  "Update Assessment Criteria Name",
				Action: updateAssessmentCriteria(assessmentCriteriaController),
			},
			{
				Title:  "Delete Assessment Criteria",
				Action: deleteAssessmentCriteria(assessmentCriteriaController),
			},
		},
	}
}

func defineAssessmentCriteria(ctrl *controller.AssessmentCriteriaController) func(*utils.MenuIO) {
	return func(io *utils.MenuIO) {
		io.Println("Defining Assessment Criteria...")
		io.Print("Enter Criteria Name (-1 to cancel): ")
		criteriaName, err := io.ReadInput()
		if err != nil || criteriaName == "-1" {
			io.Println("Cancelled.")
			return
		}
		err = ctrl.InsertAssessmentCriteria(criteriaName)
		if err != nil {
			io.Println(fmt.Sprintf("Error adding new criteria: %v", err))
		} else {
			io.Println("Criteria added successfully!")
		}
	}
}

func listAllAssessmentCriteria(ctrl *controller.AssessmentCriteriaController) func(*utils.MenuIO) {
	return func(io *utils.MenuIO) {
		io.Println("Listing All Assessment Criteria...")
		criteriaList, err := ctrl.ListAllAssessmentCriterias()
		if err != nil {
			io.Println(fmt.Sprintf("Error listing criteria: %v", err))
			return
		}
		if len(criteriaList) == 0 {
			io.Println("No assessment criteria found.")
			return
		}
		for _, c := range criteriaList {
			io.Println(fmt.Sprintf("ID: %v, Name: %v", c.ID, c.CriteriaName))
		}
	}
}

func updateAssessmentCriteria(ctrl *controller.AssessmentCriteriaController) func(*utils.MenuIO) {
	return func(io *utils.MenuIO) {
		io.Println("All Criteria:")
		criteriaList, err := ctrl.ListAllAssessmentCriterias()
		if err != nil {
			io.Println(fmt.Sprintf("Error retrieving criteria list: %v", err))
			return
		}
		for _, c := range criteriaList {
			io.Println(fmt.Sprintf("ID: %v, Name: %v", c.ID, c.CriteriaName))
		}

		io.Print("Enter Criteria ID to update (-1 to cancel): ")
		input, _ := io.ReadInput()
		if input == "-1" {
			io.Println("Cancelled.")
			return
		}
		id, err := strconv.Atoi(input)
		if err != nil {
			io.Println("Invalid ID.")
			return
		}

		criteria, err := ctrl.RetrieveByID(uint(id))
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
		err = ctrl.UpdateByID(criteria)
		if err != nil {
			io.Println(fmt.Sprintf("Error updating criteria: %v", err))
		} else {
			io.Println("Criteria updated successfully!")
		}
	}
}

func deleteAssessmentCriteria(ctrl *controller.AssessmentCriteriaController) func(*utils.MenuIO) {
	return func(io *utils.MenuIO) {
		io.Println("All Criteria:")
		criteriaList, err := ctrl.ListAllAssessmentCriterias()
		if err != nil {
			io.Println(fmt.Sprintf("Error listing criteria: %v", err))
			return
		}
		for _, c := range criteriaList {
			io.Println(fmt.Sprintf("ID: %v, Name: %v", c.ID, c.CriteriaName))
		}

		io.Print("Enter Criteria ID to delete (-1 to cancel): ")
		input, _ := io.ReadInput()
		if input == "-1" {
			io.Println("Cancelled.")
			return
		}
		id, err := strconv.Atoi(input)
		if err != nil {
			io.Println("Invalid ID.")
			return
		}

		err = ctrl.DeleteAssessmentCriteria(uint(id))
		if err != nil {
			io.Println(fmt.Sprintf("Error deleting criteria: %v", err))
		} else {
			io.Println("Criteria deleted successfully!")
		}
	}
}
