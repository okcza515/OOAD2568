package handler

import (
	"ModEd/core/cli"
	model "ModEd/curriculum/model"
	"time"

	"ModEd/curriculum/controller"
	"fmt"
)

type AdministrativeTaskMenuStateHandler struct {
	manager *cli.CLIMenuStateManager
	wrapper *controller.InstructorWorkloadModuleWrapper

	instructorWorkloadModuleMenuStateHandler *InstructorWorkloadModuleMenuStateHandler
}

func NewAdministrativeTaskMenuStateHandler(
	manager *cli.CLIMenuStateManager,
	wrapper *controller.InstructorWorkloadModuleWrapper,
	instructorWorkloadModuleMenuStateHandler *InstructorWorkloadModuleMenuStateHandler,
) *SeniorProjectWorkloadMenuStateHandler {
	return &SeniorProjectWorkloadMenuStateHandler{
		manager:                                  manager,
		wrapper:                                  wrapper,
		instructorWorkloadModuleMenuStateHandler: instructorWorkloadModuleMenuStateHandler,
	}
}

func (menu *AdministrativeTaskMenuStateHandler) Render() {
	fmt.Println("1. View all meetings")
	fmt.Println("2. View meeting by ID")
	fmt.Println("3. Create meeting")
	fmt.Println("4. Update meeting")
	fmt.Println("5. Delete meeting")
	fmt.Println("6. Add attendee to meeting")

	fmt.Println("Type 'exit' to quit")
}

func (menu *AdministrativeTaskMenuStateHandler) HandleUserInput(input string) error {
	mockMeeting := &model.Meeting{
		Title:       "Project Kickoff",
		Description: "Kickoff meeting for the project",
		Date:        time.Time{},
		Location:    "Room 101",
		Attendees:   nil,
		StartTime:   time.Time{},
		EndTime:     time.Time{},
	}
	switch input {
	case "1":
		condition := map[string]interface{}{}
		preloads := []string{}
		meetingList, err := menu.wrapper.MeetingController.List(condition, preloads...)
		if err != nil {
			fmt.Println("Error fetching meetings:", err.Error())
			return err
		}
		for _, meeting := range meetingList {
			fmt.Printf("Meeting ID: %d, Title: %s, Date: %s\n", meeting.ID, meeting.Title, meeting.Date)
		}
	case "2":
		meeting, err := menu.wrapper.MeetingController.RetrieveByID(1, []string{"Attendees"}...)
		if err != nil {
			println("Error fetching meeting:", err.Error())
			return err
		}
		fmt.Printf("Meeting ID: %d, Title: %s, Date: %s\n", meeting.ID, meeting.Title, meeting.Date)
	case "3":
		meetingFactory := model.RegularMeetingFactory{}
		err := menu.wrapper.MeetingController.CreateMeetingByFactory(meetingFactory, *mockMeeting)
		if err != nil {
			println("Error creating meeting:", err)
		}

		externalMeetingFactory := model.ExternalMeetingFactory{CompanyName: "LineWomen Wongnok"}
		err = menu.wrapper.MeetingController.CreateMeetingByFactory(externalMeetingFactory, *mockMeeting)
		if err != nil {
			println("Error creating external meeting:", err)
		}

		onlineMeetingFactory := model.OnlineMeetingFactory{ZoomLink: "https://zoom.us/j/123456789"}
		err = menu.wrapper.MeetingController.CreateMeetingByFactory(onlineMeetingFactory, *mockMeeting)
		if err != nil {
			println("Error creating online meeting:", err)
		}
	case "exit":
		fmt.Println("Exiting...")
		return nil
	default:
		fmt.Println("Invalid option")
	}
	return nil
}
