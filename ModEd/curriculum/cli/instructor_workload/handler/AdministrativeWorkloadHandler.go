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

type CreateMeeting struct {
	db *gorm.DB
}

func (c CreateMeeting) Execute() {
	meetingController := controller.NewMeetingController(c.db)
	mockMeeting := &model.Meeting{
		Title:       "Weekly Meeting",
		Description: "Weekly meeting to discuss project progress",
		Date:        time.Now(),
		Location:    "Room 101",
		StartTime:   time.Now(),
		EndTime:     time.Now().Add(2 * time.Hour),
		Attendees:   []commonModel.Instructor{},
	}
	if err := meetingController.CreateMeeting(mockMeeting); err != nil {
		fmt.Println("Error creating meeting:", err)
		return
	}

	fmt.Println("Meeting created successfully!")
}
