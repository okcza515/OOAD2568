package handler

import (
	model "ModEd/curriculum/model"
	"time"

	"ModEd/curriculum/controller"
	"ModEd/curriculum/utils"
	"fmt"
)

func RunAdministrativeWorkloadHandler(controller controller.MeetingControllerService) {
	for {
		DisplayAdministrativeWorkloadModuleMenu()
		choice := utils.GetUserChoice()
		fmt.Println("choice: ", choice)

		mockMeeting := &model.Meeting{
			Title:       "Project Kickoff",
			Description: "Kickoff meeting for the project",
			Date:        time.Time{},
			Location:    "Room 101",
			Attendees:   nil,
			StartTime:   time.Time{},
			EndTime:     time.Time{},
		}

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
			fmt.Println("Enter meeting ID:")
			var meetingID uint
			meeting, err := controller.GetByID(meetingID)
			if err != nil {
				println("Error fetching meeting:", err.Error())
				continue
			}
			fmt.Printf("Meeting ID: %d, Title: %s, Date: %s\n", meeting.ID, meeting.Title, meeting.Date)
		case "3":
			err := controller.CreateMeeting(mockMeeting)
			if err != nil {
				println("Error creating meeting:", err.Error())
				continue
			}
			fmt.Println("Meeting created successfully")
		case "4":
			meetingID := uint(1)
			err := controller.UpdateMeeting(meetingID, mockMeeting)
			if err != nil {
				println("Error updating meeting:", err.Error())
				continue
			}
			fmt.Println("Meeting updated successfully")
		case "5":
			meetingID := uint(1)
			err := controller.DeleteMeeting(meetingID)
			if err != nil {
				println("Error deleting meeting:", err.Error())
				continue
			}
			fmt.Println("Meeting deleted successfully")
		case "6":
			meetingID := uint(1)
			instructorID := uint(1)
			err := controller.AddAttendee(meetingID, instructorID)
			if err != nil {
				println("Error adding attendee:", err.Error())
				continue
			}
			fmt.Println("Attendee added successfully")
		case "exit":
			fmt.Println("Exiting Administrative Workload Module...")
		default:
			println("Invalid option. Please try again.")
		}
	}
}

func DisplayAdministrativeWorkloadModuleMenu() {
	fmt.Println("\nAdministrative Workload Module Menu:")

	fmt.Println("1. View all meetings")
	fmt.Println("2. View meeting by ID")
	fmt.Println("3. Create meeting")
	fmt.Println("4. Update meeting")
	fmt.Println("5. Delete meeting")
	fmt.Println("6. Add attendee to meeting")

	fmt.Println("Type 'exit' to quit")
}
