package main

import (
	"ModEd/project/controller"
	"ModEd/project/utils"
	"fmt"
)

func BuildAssignmentMenu(assignmentController *controller.AssignmentController) *utils.MenuItem {
	return &utils.MenuItem{
		Title: "Assignments Management",
		Children: []*utils.MenuItem{
			{
				Title: "View All Assignments",
				Action: func(io *utils.MenuIO) {
					assignments, err := assignmentController.List(map[string]interface{}{})
					if err != nil {
						io.Println(fmt.Sprintf("Error retrieving assignments: %v", err))
						return
					}
					if len(assignments) == 0 {
						io.Println("No assignments found.")
						return
					}
					for _, a := range assignments {
						io.Println(fmt.Sprintf("ID: %d | Project ID: %d | Name: %s | Due: %s", a.ID, a.SeniorProjectId, a.Name, a.DueDate.Format("2006-01-02")))
					}
				},
			},
			{
				Title: "Add New Assignment",
				Action: func(io *utils.MenuIO) {
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

					_, err = assignmentController.InsertAssignment(projectID, name, description, dueDate)
					if err != nil {
						io.Println(fmt.Sprintf("Error adding assignment: %v", err))
					} else {
						io.Println("Assignment added successfully.")
					}
				},
			},
			{
				Title: "Delete Assignment",
				Action: func(io *utils.MenuIO) {
					io.Print("Enter Assignment ID to delete: ")
					id, err := io.ReadInputID()
					if err != nil {
						return
					}

					err = assignmentController.DeleteAssignment(id)
					if err != nil {
						io.Println(fmt.Sprintf("Error deleting assignment: %v", err))
					} else {
						io.Println("Assignment deleted successfully.")
					}
				},
			},
		},
	}
}
