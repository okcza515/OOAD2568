package handler

import (
	"ModEd/curriculum/controller"
	"ModEd/curriculum/utils"
	"fmt"
)

func RunAdministrativeWorkloadHandler(controller controller.MeetingControllerService) {
	for {
		DisplayAdministrativeWorkloadModuleMenu()
		choice := utils.GetUserChoice()
		fmt.Println("choice: ", choice)

		switch choice {
		case "1":
			meetingList, err := controller.GetAll()
			if err != nil {
				println("Error fetching meetings:", err.Error())
				continue
			}
			for _, meeting := range *meetingList {
				fmt.Printf("Meeting ID: %d, Title: %s, Date: %s\n", meeting.ID, meeting.Title, meeting.Date)
			}
		case "2":
			fmt.Println("View meeting by ID...")
		default:
			println("Invalid option. Please try again.")
		}
	}
}

func DisplayAdministrativeWorkloadModuleMenu() {
	fmt.Println("\nAdministrative Workload Module Menu:")
	fmt.Println("1. View all meetings")
	fmt.Println("2. View meeting by ID")
	fmt.Println("Type 'exit' to quit")
}
