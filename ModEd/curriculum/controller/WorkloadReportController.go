package controller

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type WorkloadReportControllerInterface interface {
	GenerateWorkloadReport(instructorID uint) error
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
		// meetingController := NewMeetingController(b.db)
		// meetings, err := meetingController.List()
	}

	if b.includeProjects {
		// Generate project report
	}
	if b.includeStudentRequests {
		// Generate student request report
	}

	return nil
}
