package controller

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type WorkloadReportControllerInterface interface {
	GenerateWorkloadReport(instructorID uint) error
}

type InstructorPerformanceReport struct {
	InstructorID       uint
	TotalClasses       int
	TotalTeachingHours float64
	TotalProjects      int
	TotalMeetings      int
}

type InstructorReportFacade struct {
	db                       *gorm.DB
	InstructorID             uint
	ClassController          ClassControllerInterface
	CourseController         CourseControllerInterface
	CurriculumController     CurriculumControllerInterface
	ClassMaterialController  ClassMaterialControllerInterface
	CoursePlanController     CoursePlanControllerInterface
	MeetingController        MeetingControllerInterface
	ProjectController        ProjectControllerInterface
	StudentRequestController StudentRequestControllerInterface
}

type WorkloadReportBuilder struct {
	db                       *gorm.DB
	InstructorID             uint
	Header                   string
	ClassController          ClassControllerInterface
	CourseController         CourseControllerInterface
	CurriculumController     CurriculumControllerInterface
	ClassMaterialController  ClassMaterialControllerInterface
	CoursePlanController     CoursePlanControllerInterface
	MeetingController        MeetingControllerInterface
	ProjectController        ProjectControllerInterface
	StudentRequestController StudentRequestControllerInterface

	startDate *time.Time
	endDate   *time.Time

	includeClasses         bool
	includeCourses         bool
	includeCurriculums     bool
	includeClassMaterials  bool
	includeCoursePlans     bool
	includeMeetings        bool
	includeProjects        bool
	includeStudentRequests bool
}

func NewInstructorReportFacade(db *gorm.DB, instructorID uint) *InstructorReportFacade {
	return &InstructorReportFacade{
		db:           db,
		InstructorID: instructorID,
	}
}

func NewWorkloadReportBuilder(db *gorm.DB, instructorID uint) *WorkloadReportBuilder {
	return &WorkloadReportBuilder{
		db:           db,
		InstructorID: instructorID,
	}
}

func (b *WorkloadReportBuilder) SetHeader(header string) *WorkloadReportBuilder {
	b.Header = header
	return b
}

func (b *WorkloadReportBuilder) WithClasses() *WorkloadReportBuilder {
	b.includeClasses = true
	return b
}

func (b *WorkloadReportBuilder) WithCourses() *WorkloadReportBuilder {
	b.includeCourses = true
	return b
}

func (b *WorkloadReportBuilder) WithCurriculums() *WorkloadReportBuilder {
	b.includeCurriculums = true
	return b
}

func (b *WorkloadReportBuilder) WithClassMaterials() *WorkloadReportBuilder {
	b.includeClassMaterials = true
	return b
}

func (b *WorkloadReportBuilder) WithCoursePlans() *WorkloadReportBuilder {
	b.includeCoursePlans = true
	return b
}

func (b *WorkloadReportBuilder) WithMeetings() *WorkloadReportBuilder {
	b.includeMeetings = true
	return b
}

func (b *WorkloadReportBuilder) WithProjects() *WorkloadReportBuilder {
	b.includeProjects = true
	return b
}

func (b *WorkloadReportBuilder) WithStudentRequests() *WorkloadReportBuilder {
	b.includeStudentRequests = true
	return b
}

func (b *WorkloadReportBuilder) SetDateRange(start, end *time.Time) *WorkloadReportBuilder {
	b.startDate = start
	b.endDate = end
	return b
}

func (b *WorkloadReportBuilder) Generate() error {
	if b.Header != "" {
		fmt.Printf("--------------------%s--------------------\n", b.Header)
	}

	if b.includeClasses {
		classController := NewClassController(b.db)
		classes, err := classController.GetClasses()
		if err != nil {
			fmt.Println("Error fetching classes:", err)
			return err
		}

		for _, class := range classes {
			fmt.Printf("Class ID: %d, Course ID: %d, Section: %d, Schedule: %s\n",
				class.ClassId, class.CourseId, class.Section, class.Schedule.Format(time.RFC1123))
		}
	}
	if b.includeCourses {
		courseController := NewCourseController(b.db)
		courses, err := courseController.GetCourses()
		if err != nil {
			fmt.Println("Error fetching courses:", err)
		}

		for _, course := range courses {
			fmt.Printf("Course ID: %d, Name: %s, Description: %s\n",
				course.CourseId, course.Name, course.Description)
		}
	}
	if b.includeCurriculums {
		// Generate curriculum report
	}
	if b.includeClassMaterials {
		// Generate class material report
	}
	if b.includeCoursePlans {
		// Generate course plan report
	}

	if b.includeMeetings {
		meetingController := NewMeetingController(b.db)
		meetings, err := meetingController.List(nil)
		if err != nil {
			fmt.Println("Error fetching meetings:", err)
		}

		for _, meeting := range meetings {
			fmt.Printf("Meeting ID: %d, Title: %s, Start Time: %s - End Time: %s, Location: %s\n",
				meeting.GetID(), meeting.GetTitle(), meeting.StartTime.Format(time.RFC1123), meeting.EndTime.Format(time.RFC1123), meeting.GetLocation())
		}
	}

	if b.includeProjects {
		// projectCommitteeController := seniorProjectController.CommitteeController{b.db}
		// projectCommitteeController.ListCommitteesByInstructor(1)
	}
	if b.includeStudentRequests {
		// Generate student request report
	}

	return nil
}

func (f *InstructorReportFacade) GeneratePerformanceReport(instructorID uint, startDate, endDate *time.Time) (*InstructorPerformanceReport, error) {
	report := &InstructorPerformanceReport{InstructorID: instructorID}
	fmt.Println("---------------------Instructor Performance Report--------------------")
	// 1. Count Classes
	classController := NewClassController(f.db)
	classes, err := classController.GetClasses()
	if err != nil {
		return nil, fmt.Errorf("fetching classes: %w", err)
	}
	report.TotalClasses = len(classes)

	// Sum teaching hours (example assumes Schedule duration or custom hours field)
	const classDuration = 2.0 // hours

	var totalHours float64
	for range classes {
		totalHours += classDuration
	}
	report.TotalTeachingHours = totalHours

	// // 2. Count Projects
	// projectController := NewProjectController(f.db)
	// projectCount, err := projectController.CountProjectsByInstructor(instructorID, startDate, endDate)
	// if err != nil {
	// 	return nil, fmt.Errorf("fetching projects: %w", err)
	// }
	// report.TotalProjects = projectCount

	// 3. Count Meetings
	meetingController := NewMeetingController(f.db)
	meetings, err := meetingController.List(nil)
	if err != nil {
		return nil, fmt.Errorf("fetching meetings: %w", err)
	}
	report.TotalMeetings = len(meetings)

	return report, nil
}
