package menu

import (
	"ModEd/project/controller"
	"ModEd/project/model"
	"ModEd/project/utils"
	"fmt"
)

func BuildPresentationMenu(presentationController *controller.PresentationController) *utils.MenuItem {
	return &utils.MenuItem{
		Title: "Presentations Management",
		Children: []*utils.MenuItem{
			{
				Title: "View All Presentations",
				Action: func(io *utils.MenuIO) {
					io.Println("Viewing Presentations...")

					presentations, err := presentationController.ListAllPresentations()
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
				},
			},
			{
				Title: "Add New Presentation",
				Action: func(io *utils.MenuIO) {
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

					// Convert the string input into PresentationType
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
					_, err = presentationController.InsertPresentation(projectID, presentationType, dueDate)
					if err != nil {
						io.Println(fmt.Sprintf("Error inserting presentation: %v", err))
					} else {
						io.Println("Presentation added successfully!")
					}
				},
			},
			{
				Title: "View Presentation by ID",
				Action: func(io *utils.MenuIO) {
					io.Println("Viewing Presentation by ID...")

					io.Print("Enter Presentation ID: ")
					presentationID, err := io.ReadInputID()
					if err != nil {
						return
					}

					presentation, err := presentationController.RetrievePresentation(presentationID)
					if err != nil {
						io.Println(fmt.Sprintf("Error retrieving presentation: %v", err))
						return
					}

					io.Println(fmt.Sprintf("Presentation ID: %d\nProject ID: %d\nType: %s\nDate: %s",
						presentation.ID, presentation.SeniorProjectId, presentation.PresentationType, presentation.Date.Format("2006-01-02")))
				},
			},
			{
				Title: "Update Presentation",
				Action: func(io *utils.MenuIO) {
					io.Println("Updating Presentation...")

					io.Print("Enter Presentation ID to update: ")
					presentationID, err := io.ReadInputID()
					if err != nil {
						return
					}

					presentation, err := presentationController.RetrievePresentation(presentationID)
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

					err = presentationController.UpdatePresentation(presentation)
					if err != nil {
						io.Println(fmt.Sprintf("Error updating presentation: %v", err))
					} else {
						io.Println("Presentation updated successfully!")
					}
				},
			},
			{
				Title: "Delete Presentation",
				Action: func(io *utils.MenuIO) {
					io.Println("Deleting Presentation...")

					io.Print("Enter Presentation ID to delete: ")
					presentationID, err := io.ReadInputID()
					if err != nil {
						return
					}

					err = presentationController.DeletePresentation(presentationID)
					if err != nil {
						io.Println(fmt.Sprintf("Error deleting presentation: %v", err))
					} else {
						io.Println("Presentation deleted successfully!")
					}
				},
			},
		},
	}
}
