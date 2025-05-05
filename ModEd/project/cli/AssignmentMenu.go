package main

import (
	"ModEd/project/controller"
	"ModEd/project/utils"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func BuildAssignmentMenu(assignmentController *controller.AssignmentController) *utils.MenuItem {
	return &utils.MenuItem{
		Title: "Assignments Management",
		Children: []*utils.MenuItem{
			{
				Title: "View All Assignments",
				Action: func(io *utils.MenuIO) {
					assignments, err := assignmentController.ListAllAssignments()
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
					projectIDInput, _ := io.ReadInput()
					projectIDInput = strings.TrimSpace(projectIDInput)
					projectID, err := strconv.Atoi(projectIDInput)
					if err != nil {
						io.Println("Invalid Project ID.")
						return
					}

					io.Print("Enter Assignment Name: ")
					name, _ := io.ReadInput()
					name = strings.TrimSpace(name)

					io.Print("Enter Assignment Description: ")
					description, _ := io.ReadInput()
					description = strings.TrimSpace(description)

					io.Print("Enter Due Date (YYYY-MM-DD): ")
					dueInput, _ := io.ReadInput()
					dueInput = strings.TrimSpace(dueInput)
					dueDate, err := time.Parse("2006-01-02", dueInput)
					if err != nil {
						io.Println("Invalid date format.")
						return
					}

					_, err = assignmentController.InsertAssignment(uint(projectID), name, description, dueDate)
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
					idInput, _ := io.ReadInput()
					id, err := strconv.Atoi(strings.TrimSpace(idInput))
					if err != nil {
						io.Println("Invalid Assignment ID.")
						return
					}

					err = assignmentController.DeleteAssignment(uint(id))
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
