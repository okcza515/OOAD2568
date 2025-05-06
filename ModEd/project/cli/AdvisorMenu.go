package main

import (
	"ModEd/project/controller"
	"ModEd/project/utils"
	"fmt"
	"strings"
)

func BuildAdvisorMenu(advisorController *controller.AdvisorController) *utils.MenuItem {
	return &utils.MenuItem{
		Title: "Advisor Manager",
		Children: []*utils.MenuItem{
			{
				Title: "Assign Advisor to Project",
				Action: func(io *utils.MenuIO) {
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
					isPrimaryStr, err := io.ReadInput()
					if err != nil {

						return
					}
					isPrimary := strings.ToLower(isPrimaryStr) == "yes" || strings.ToLower(isPrimaryStr) == "y"

					advisor, err := advisorController.AssignAdvisor(uint(projectId), uint(instructorId), isPrimary)
					if err != nil {
						io.Println(fmt.Sprintf("Error assigning advisor: %v", err))
					} else {
						io.Println(fmt.Sprintf("Advisor assigned successfully! Advisor ID: %v", advisor.ID))
					}
				},
			},
			{
				Title: "Update Advisor Role",
				Action: func(io *utils.MenuIO) {
					io.Println("Updating Advisor Role...")

					io.Print("Enter Advisor ID (-1 to cancel): ")
					advisorId, err := io.ReadInputID()
					if err != nil {
						return
					}

					io.Print("Set as primary advisor? (yes/no): ")
					isPrimaryStr, err := io.ReadInput()
					if err != nil {
						return
					}
					isPrimary := strings.ToLower(isPrimaryStr) == "yes" || strings.ToLower(isPrimaryStr) == "y"

					err = advisorController.UpdateAdvisorRole(uint(advisorId), isPrimary)
					if err != nil {
						io.Println(fmt.Sprintf("Error updating advisor role: %v", err))
					} else {
						io.Println("Advisor role updated successfully!")
					}
				},
			},
			{
				Title: "Remove Advisor",
				Action: func(io *utils.MenuIO) {
					io.Println("Removing Advisor...")

					io.Print("Enter Advisor ID (-1 to cancel): ")
					advisorId, err := io.ReadInputID()
					if err != nil {
						return
					}

					err = advisorController.DeleteByID(advisorId)
					if err != nil {
						io.Println(fmt.Sprintf("Error removing advisor: %v", err))
					} else {
						io.Println("Advisor removed successfully!")
					}
				},
			},
			{
				Title: "List Advisors by Project",
				Action: func(io *utils.MenuIO) {
					io.Println("Listing Advisors by Project...")

					io.Print("Enter Project ID (-1 to cancel): ")
					projectId, err := io.ReadInputID()
					if err != nil {
						return
					}

					advisors, err := advisorController.List(map[string]interface{}{
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
				},
			},
			{
				Title: "List Projects by Instructor",
				Action: func(io *utils.MenuIO) {
					io.Println("Listing Projects by Instructor...")

					io.Print("Enter Instructor ID (-1 to cancel): ")
					instructorId, err := io.ReadInputID()
					if err != nil {
						return
					}

					advisors, err := advisorController.List(map[string]interface{}{
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
				},
			},
		},
	}
}
