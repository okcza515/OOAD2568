// MEP-1008
package handler

import (
	commonModel "ModEd/common/model"
	controller "ModEd/curriculum/controller"
	model "ModEd/curriculum/model"
	utils "ModEd/curriculum/utils"
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
	adminstrativeMenu.Add(string(MENU_MEETING ), NewMeetingWorkloadHandler(a.db))
	adminstrativeMenu.Add(string(MENU_STUDENT_REQUEST), nil)
	adminstrativeMenu.SetBackHandler(Back{})
	adminstrativeMenu.SetDefaultHandler(UnknownCommand{})
	adminstrativeMenu.Execute()
}

// Submenu for MeetingHandler
func (m MeetingWorkloadHandler) Execute() {
	meetingMenu := NewMenuHandler("Academic Workload Menu", true)
	meetingMenu.Add(string(MENU_MEETING_LIST), ListAllMeeting{db: m.db})
	meetingMenu.Add(string(MENU_MEETING_CREATE), CreateMeeting{db: m.db})
	meetingMenu.Add(string(MENU_MEETING_CREATE_EXTERNAL), CreateExternalMeeting{db: m.db})
	meetingMenu.Add(string(MENU_MEETING_CREATE_ONLINE), CreateOnlineMeeting{db: m.db})
	meetingMenu.Add(string(MENU_MEETING_RETRIEVE), RetrieveMeetingById{db: m.db})
	meetingMenu.Add(string(MENU_MEETING_ADD_ATTENDEE), AddAttendee{db: m.db})
	meetingMenu.Add(string(MENU_MEETING_UPDATE), UpdateMeetingById{db: m.db})
	meetingMenu.Add(string(MENU_MEETING_DELETE), DeleteMeetingById{db: m.db})
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
type RetrieveMeetingById struct{ db *gorm.DB }
type UpdateMeetingById struct{ db *gorm.DB }
type DeleteMeetingById struct{ db *gorm.DB }
type ListAllMeeting struct{ db *gorm.DB }

var mockMeeting = &model.Meeting{
	Title:       "Weekly Meeting",
	Description: "Weekly meeting to discuss project progress",
	Date:        time.Now(),
	Location:    "Room 101",
	StartTime:   time.Now(),
	EndTime:     time.Now().Add(2 * time.Hour),
	Attendees:   nil,
}

var mockInstructor = &commonModel.Instructor{
	InstructorCode: "INS001",
	FirstName:      "John",
	LastName:       "Doe",
	Email:          "johndoe@gmail.com",
	StartDate:      nil,
	Department:     nil,
}

var mockStudent = &commonModel.Student{
	StudentCode: "STU001",
	FirstName:   "Jane",
	LastName:    "Smith",
	Email:       "",
	StartDate:   time.Time{},
	Program:     commonModel.REGULAR,
	Department:  "Computer Engineering",
	Status:      nil,
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
	id := utils.GetUserInputUint("Enter MeetingID: ")
	meetingController := controller.NewMeetingController(c.db)
	mockAttendees := []model.AttendeeAdapter{
		model.InstructorAdapter{Instructor: *mockInstructor},
		model.StudentAdapter{Student: *mockStudent},
	}
	for _, attendee := range mockAttendees {
		err := meetingController.AddAttendee(id, attendee)
		if err != nil {
			println("Error adding attendee:", err)
		} else {
			fmt.Println("Attendee added successfully!")
		}
	}
}

func (c RetrieveMeetingById) Execute() {
	id := utils.GetUserInputUint("Enter MeetingID: ")
	meetingController := controller.NewMeetingController(c.db)
	meeting, err := meetingController.RetrieveByID(id)
	if err != nil {
		println("Error retrieving meeting:", err)
		return
	}
	fmt.Printf("Meeting ID: %d, Title: %s, Location: %s, Date: %s\n", meeting.ID, meeting.Title, meeting.Location, meeting.Date.Format("2006-01-02"))
}

func (c UpdateMeetingById) Execute() {
	id := utils.GetUserInputUint("Enter MeetingID: ")
	meetingController := controller.NewMeetingController(c.db)
	mockMeeting.ID = id
	mockMeeting.Title = "Daily Meeting"
	mockMeeting.Location = "Room 102"
	err := meetingController.UpdateByID(mockMeeting)
	if err != nil {
		println("Error updating meeting:", err)
	} else {
		fmt.Println("Meeting updated successfully!")
	}
}

func (c DeleteMeetingById) Execute() {
	id := utils.GetUserInputUint("Enter MeetingID: ")
	meetingController := controller.NewMeetingController(c.db)
	err := meetingController.DeleteByID(id)
	if err != nil {
		println("Error deleting meeting:", err)
	} else {
		fmt.Println("Meeting deleted successfully!")
	}
}

func (c ListAllMeeting) Execute() {
	meetingController := controller.NewMeetingController(c.db)
	meetings, err := meetingController.List(nil)
	if err != nil {
		println("Error listing meetings:", err)
		return
	}
	for _, meeting := range meetings {
		fmt.Printf("Meeting ID: %d, Title: %s, Location: %s, Date: %s\n", meeting.ID, meeting.Title, meeting.Location, meeting.Date.Format("2006-01-02"))
	}
}
