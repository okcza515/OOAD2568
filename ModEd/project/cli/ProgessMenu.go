package menus

import (
	"ModEd/project/controller"
	"ModEd/project/utils"
	"fmt"
	"strconv"
)

// BuildProgressMenu creates the progress tracking menu
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
				Title: "Mark Progress as Completed",
				Action: func(io *utils.MenuIO) {
					io.Println("Marking Progress as Completed...")
					io.Print("Enter Progress ID: ")

					input, err := io.ReadInput()
					if err != nil {
						io.Println(fmt.Sprintf("Error reading input: %v", err))
						return
					}

					progressID, err := strconv.ParseUint(input, 10, 32)
					if err != nil {
						io.Println(fmt.Sprintf("Invalid Progress ID: %v", err))
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
		},
	}
}
