package controller

import (
	"fmt"

	"gorm.io/gorm"
)

type WorkloadReportController struct {
	Connector *gorm.DB
}

type WorkloadReportControllerInterface interface {
	GenerateDailyWorkloadReport(instructorId uint) error
	GenerateWorkloadReportWithFilter(instructorId uint, startDate, endDate string) error
}

func CreateWorkloadReportController(db *gorm.DB) *WorkloadReportController {
	return &WorkloadReportController{
		Connector: db,
	}
}

type WorkloadReportFacade struct {
	ClassController           *ClassController
	MeetingController         *MeetingController
	StudentWorkloadController *StudentRequestController
}

// GenerateDailyWorkloadReport generates a simple daily workload report.
func (f *WorkloadReportFacade) GenerateDailyWorkloadReport(instructorId uint) error {
	fmt.Println("Generating Daily Workload Report...")

	// Get classes
	classList, err := f.ClassController.GetClasses()
	if err != nil {
		return fmt.Errorf("failed to get classes: %w", err)
	}
	for _, class := range classList {
		fmt.Printf("Class ID: %d, Class Name: %s, Schedule: %s\n", class.ClassId, class.Course.Name, class.Schedule)
	}

	// Get meetings
	meetingList, err := f.MeetingController.List(nil)
	if err != nil {
		return fmt.Errorf("failed to get meetings: %w", err)
	}
	for _, meeting := range meetingList {
		fmt.Printf("Meeting Title: %s, Start: %s, End: %s, Location: %s\n",
			meeting.Title, meeting.StartTime.Format("15:04"), meeting.EndTime.Format("15:04"), meeting.Location)
	}

	// Get student workload requests
	studentWorkloadList, err := f.StudentWorkloadController.ListStudentRequest(instructorId)
	if err != nil {
		return fmt.Errorf("failed to get student workload requests: %w", err)
	}
	for _, studentRequest := range studentWorkloadList {
		fmt.Printf("Student Code: %s, Request Type: %s\n", studentRequest.StudentCode, studentRequest.RequestType)
	}

	return nil
}

// WorkloadReportBuilder is the builder
type WorkloadReportBuilder struct {
	facade                 *WorkloadReportFacade
	includeClasses         bool
	includeMeetings        bool
	includeStudentRequests bool
	startDate              string
	endDate                string
}

// NewWorkloadReportBuilder creates a new builder instance
func NewWorkloadReportBuilder(facade *WorkloadReportFacade) *WorkloadReportBuilder {
	return &WorkloadReportBuilder{
		facade: facade,
	}
}

func (b *WorkloadReportBuilder) IncludeClasses() *WorkloadReportBuilder {
	b.includeClasses = true
	return b
}

func (b *WorkloadReportBuilder) IncludeMeetings() *WorkloadReportBuilder {
	b.includeMeetings = true
	return b
}

func (b *WorkloadReportBuilder) IncludeStudentRequests() *WorkloadReportBuilder {
	b.includeStudentRequests = true
	return b
}

func (b *WorkloadReportBuilder) SetDateRange(start, end string) *WorkloadReportBuilder {
	b.startDate = start
	b.endDate = end
	return b
}

func (b *WorkloadReportBuilder) Generate(instructorId uint) error {
	if b.includeClasses {
		classList, err := b.facade.ClassController.GetClasses()
		if err != nil {
			return err
		}
		for _, class := range classList {
			fmt.Printf("Class ID: %d, Class Name: %s, Schedule: %s\n", class.ClassId, class.Course.Name, class.Schedule)
		}
	}

	if b.includeMeetings {
		meetingList, err := b.facade.MeetingController.List(nil)
		if err != nil {
			return err
		}
		for _, meeting := range meetingList {
			fmt.Printf("Meeting Title: %s, Start Time: %s, End Time: %s, Location: %s\n",
				meeting.Title, meeting.StartTime.Format("15:04"), meeting.EndTime.Format("15:04"), meeting.Location)
		}
	}

	if b.includeStudentRequests {
		studentWorkloadList, err := b.facade.StudentWorkloadController.ListStudentRequest(instructorId)
		if err != nil {
			return err
		}
		for _, studentRequest := range studentWorkloadList {
			fmt.Printf("Student Code: %s, Request Type: %s\n", studentRequest.StudentCode, studentRequest.RequestType)
		}
	}

	// (Optional) Handle Date Range filtering logic here if needed

	return nil
}
