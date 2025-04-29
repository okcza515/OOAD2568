package menus

import (
	"ModEd/project/controller"
	"ModEd/project/utils"
	"fmt"
	"strconv"
	"strings"
)

func BuildAdvisorMenu(
	advisorController *controller.AdvisorController,
) *utils.MenuItem {
	return &utils.MenuItem{
		Title: "Advisor Manager",
		Children: []*utils.MenuItem{
			{
				Title: "Assign Advisor to Project",
				Action: func(io *utils.MenuIO) {
					io.Println("Assigning Advisor to Project...")

					io.Print("Enter Project ID (-1 to cancel): ")
					projectIdStr, err := io.ReadInput()
					if err != nil || projectIdStr == "-1" {
						io.Println("Cancelled.")
						return
					}
					projectId, err := strconv.Atoi(projectIdStr)
					if err != nil {
						io.Println("Invalid Project ID format.")
						return
					}

					io.Print("Enter Instructor ID (-1 to cancel): ")
					instructorIdStr, err := io.ReadInput()
					if err != nil || instructorIdStr == "-1" {
						io.Println("Cancelled.")
						return
					}
					instructorId, err := strconv.Atoi(instructorIdStr)
					if err != nil {
						io.Println("Invalid Instructor ID format.")
						return
					}

					io.Print("Is this a primary advisor? (yes/no): ")
					isPrimaryStr, err := io.ReadInput()
					if err != nil {
						io.Println("Cancelled.")
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
					advisorIdStr, err := io.ReadInput()
					if err != nil || advisorIdStr == "-1" {
						io.Println("Cancelled.")
						return
					}
					advisorId, err := strconv.Atoi(advisorIdStr)
					if err != nil {
						io.Println("Invalid Advisor ID format.")
						return
					}

					io.Print("Set as primary advisor? (yes/no): ")
					isPrimaryStr, err := io.ReadInput()
					if err != nil {
						io.Println("Cancelled.")
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
					advisorIdStr, err := io.ReadInput()
					if err != nil || advisorIdStr == "-1" {
						io.Println("Cancelled.")
						return
					}
					advisorId, err := strconv.Atoi(advisorIdStr)
					if err != nil {
						io.Println("Invalid Advisor ID format.")
						return
					}

					err = advisorController.RemoveAdvisor(uint(advisorId))
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
					projectIdStr, err := io.ReadInput()
					if err != nil || projectIdStr == "-1" {
						io.Println("Cancelled.")
						return
					}
					projectId, err := strconv.Atoi(projectIdStr)
					if err != nil {
						io.Println("Invalid Project ID format.")
						return
					}

					advisors, err := advisorController.ListAdvisorsByProject(projectId)
					if err != nil {
						io.Println(fmt.Sprintf("Error listing advisors: %v", err))
						return
					}

					if len(advisors) == 0 {
						io.Println("No advisors found for this project.")
						return
					}

					io.Println(fmt.Sprintf("Advisors for Project ID %v:", projectId))
					for _, a := range advisors {
						role := "Secondary"
						if a.IsPrimary {
							role = "Primary"
						}
						io.Println(fmt.Sprintf("Advisor ID: %v, Instructor ID: %v, Role: %v", a.ID, a.InstructorId, role))
					}
				},
			},
			{
				Title: "List Projects by Instructor",
				Action: func(io *utils.MenuIO) {
					io.Println("Listing Projects by Instructor...")

					io.Print("Enter Instructor ID (-1 to cancel): ")
					instructorIdStr, err := io.ReadInput()
					if err != nil || instructorIdStr == "-1" {
						io.Println("Cancelled.")
						return
					}
					instructorId, err := strconv.Atoi(instructorIdStr)
					if err != nil {
						io.Println("Invalid Instructor ID format.")
						return
					}

					advisors, err := advisorController.ListAdvisorsByInstructor(instructorId)
					if err != nil {
						io.Println(fmt.Sprintf("Error listing projects: %v", err))
						return
					}

					if len(advisors) == 0 {
						io.Println("No projects found for this instructor.")
						return
					}

					io.Println(fmt.Sprintf("Projects for Instructor ID %v:", instructorId))
					for _, a := range advisors {
						role := "Secondary"
						if a.IsPrimary {
							role = "Primary"
						}
						io.Println(fmt.Sprintf("Advisor ID: %v, Project ID: %v, Role: %v", a.ID, a.SeniorProjectId, role))
					}
				},
			},
		},
	}
}
