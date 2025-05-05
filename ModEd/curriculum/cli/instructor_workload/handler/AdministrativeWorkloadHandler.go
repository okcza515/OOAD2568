// MEP-1008
package handler

import (
	commonModel "ModEd/common/model"
	controller "ModEd/curriculum/controller"
	model "ModEd/curriculum/model"

	"fmt"

	"gorm.io/gorm"

	"time"
)

type AdminstrativeWorkloadHandler struct {
	db *gorm.DB
}

func NewAdminstrativeWorkloadHandler(db *gorm.DB) AdminstrativeWorkloadHandler {
	return AdminstrativeWorkloadHandler{db: db}
}

// Adminstrative Menu
func (a AdminstrativeWorkloadHandler) Execute() {
	adminstrativeMenu := NewMenuHandler("Adminstrative Workload Menu", true)
	adminstrativeMenu.Add("Meeting", NewMeetingWorkloadHandler(a.db))
	adminstrativeMenu.Add("Student Request", nil)
	adminstrativeMenu.SetBackHandler(Back{})
	adminstrativeMenu.SetDefaultHandler(UnknownCommand{})
	adminstrativeMenu.Execute()
}

// Submenu for MeetingHandler
func (m MeetingWorkloadHandler) Execute() {
	meetingMenu := NewMenuHandler("Academic Workload Menu", true)
	meetingMenu.Add("Create Meeting", CreateMeeting{db: m.db})
	meetingMenu.Add("Create External Meeting", CreateExternalMeeting{db: m.db})
	meetingMenu.Add("Create Online Meeting", CreateOnlineMeeting{db: m.db})
	meetingMenu.Add("Add Attendee", AddAttendee{db: m.db})
	meetingMenu.SetBackHandler(Back{})
	meetingMenu.SetDefaultHandler(UnknownCommand{})
	meetingMenu.Execute()
}

type MeetingWorkloadHandler struct {
	db *gorm.DB
}

func NewMeetingWorkloadHandler(db *gorm.DB) MeetingWorkloadHandler {
	return MeetingWorkloadHandler{db: db}
}

type CreateMeeting struct{ db *gorm.DB }
type CreateExternalMeeting struct{ db *gorm.DB }
type CreateOnlineMeeting struct{ db *gorm.DB }
type AddAttendee struct{ db *gorm.DB }

var mockMeeting = &model.Meeting{
	Title:       "Weekly Meeting",
	Description: "Weekly meeting to discuss project progress",
	Date:        time.Now(),
	Location:    "Room 101",
	StartTime:   time.Now(),
	EndTime:     time.Now().Add(2 * time.Hour),
	Attendees:   nil,
}

func (c CreateMeeting) Execute() {
	meetingController := controller.NewMeetingController(c.db)
	meetingFactory := model.RegularMeetingFactory{}
	err := meetingController.CreateMeetingByFactory(meetingFactory, *mockMeeting)
	if err != nil {
		println("Error creating meeting:", err)
	}
	fmt.Println("Meeting created successfully!")
}

func (c CreateExternalMeeting) Execute() {
	meetingController := controller.NewMeetingController(c.db)
	externalMeetingFactory := model.ExternalMeetingFactory{CompanyName: "LineWomen Wongnok"}
	err := meetingController.CreateMeetingByFactory(externalMeetingFactory, *mockMeeting)
	if err != nil {
		println("Error creating external meeting:", err)
	} else {
		fmt.Println("External Meeting created successfully!")
	}
}

func (c CreateOnlineMeeting) Execute() {
	meetingController := controller.NewMeetingController(c.db)
	onlineMeetingFactory := model.OnlineMeetingFactory{ZoomLink: "https://zoom.us/j/123456789"}
	err := meetingController.CreateMeetingByFactory(onlineMeetingFactory, *mockMeeting)
	if err != nil {
		println("Error creating online meeting:", err)
	} else {
		fmt.Println("Online Meeting created successfully!")
	}
}

func (c AddAttendee) Execute() {
	mockMeetingId := uint(1)
	meetingController := controller.NewMeetingController(c.db)
	mockAttendees := []model.AttendeeAdapter{
		model.InstructorAdapter{Instructor: commonModel.Instructor{
			InstructorCode: "INS001",
			FirstName:      "John",
			LastName:       "Doe",
			Email:          "johndoe@gmail.com",
			StartDate:      nil,
			Department:     nil,
		}},
		model.StudentAdapter{Student: commonModel.Student{
			StudentCode: "STU001",
			FirstName:   "Jane",
			LastName:    "Smith",
			Email:       "",
			StartDate:   time.Time{},
			Program:     commonModel.REGULAR,
			Department:  "Computer Engineering",
			Status:      nil,
		}},
	}
	for _, attendee := range mockAttendees {
		err := meetingController.AddAttendee(mockMeetingId, attendee)
		if err != nil {
			println("Error adding attendee:", err)
		} else {
			fmt.Println("Attendee added successfully!")
		}
	}
}
