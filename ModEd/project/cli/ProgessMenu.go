package main

import (
	"ModEd/project/controller"
	"ModEd/project/utils"
	"fmt"
)

func BuildProgressMenu(progressController *controller.ProgressController) *utils.MenuItem {
	return &utils.MenuItem{
		Title: "Track Progress",
		Children: []*utils.MenuItem{
			{
				Title: "View All Progress",
				Action: func(io *utils.MenuIO) {
					io.Println("Viewing All Progress...")

					formattedList, err := progressController.GetFormattedProgressList()
					if err != nil {
						io.Println(fmt.Sprintf("Error retrieving progress: %v", err))
						return
					}

					if len(formattedList) == 0 {
						io.Println("No progress found.")
						return
					}

					io.Println("Progress List:")
					for _, progress := range formattedList {
						io.Println(progress)
					}
				},
			},
			{
				Title: "Add New Progress",
				Action: func(io *utils.MenuIO) {
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

					err = progressController.AddNewProgress(uint(assignmentID), name)
					if err != nil {
						io.Println(fmt.Sprintf("Error adding new progress: %v", err))
					} else {
						io.Println("Progress added successfully!")
					}
				},
			},
			{
				Title: "View Progress by ID",
				Action: func(io *utils.MenuIO) {
					io.Println("Viewing Progress by ID...")
					io.Print("Enter Progress ID: ")

					progressID, err := io.ReadInputID()
					if err != nil {
						return
					}

					progress, err := progressController.RetrieveByID(progressID)
					if err != nil {
						io.Println(fmt.Sprintf("Error retrieving progress: %v", err))
						return
					}

					io.Println(fmt.Sprintf("Progress ID: %d, Assignment ID: %d, Name: %s, Completed: %t",
						progress.ID, progress.AssignmentId, progress.Name, progress.IsCompleted))
				},
			},
			{
				Title: "Update Progress Name",
				Action: func(io *utils.MenuIO) {
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

					err = progressController.UpdateProgressName(progressID, newName)
					if err != nil {
						io.Println(fmt.Sprintf("Error updating progress name: %v", err))
					} else {
						io.Println("Progress name updated successfully!")
					}
				},
			},
			{
				Title: "Delete Progress",
				Action: func(io *utils.MenuIO) {
					io.Println("Deleting Progress...")
					io.Print("Enter Progress ID to delete: ")

					progressID, err := io.ReadInputID()
					if err != nil {
						return
					}

					err = progressController.DeleteByID(uint(progressID))
					if err != nil {
						io.Println(fmt.Sprintf("Error deleting progress: %v", err))
					} else {
						io.Println("Progress deleted successfully!")
					}
				},
			},
			{
				Title: "Mark Progress as Completed",
				Action: func(io *utils.MenuIO) {
					io.Println("Marking Progress as Completed...")
					io.Print("Enter Progress ID: ")

					progressID, err := io.ReadInputID()
					if err != nil {
						return
					}

					err = progressController.MarkAsCompleted(uint(progressID))
					if err != nil {
						io.Println(fmt.Sprintf("Error marking progress as completed: %v", err))
					} else {
						io.Println("Progress marked as completed successfully!")
					}
				},
			},
			{
				Title: "Mark Progress as Incomplete",
				Action: func(io *utils.MenuIO) {
					io.Println("Marking Progress as Incomplete...")
					io.Print("Enter Progress ID: ")

					progressID, err := io.ReadInputID()
					if err != nil {
						return
					}

					err = progressController.MarkProgressAsIncomplete(uint(progressID))
					if err != nil {
						io.Println(fmt.Sprintf("Error marking progress as incomplete: %v", err))
					} else {
						io.Println("Progress marked as incomplete successfully!")
					}
				},
			},
		},
	}
}
