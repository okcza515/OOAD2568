package model

import (
	"time"

	commonModel "ModEd/common/model"
	curriculumModel "ModEd/curriculum/model"

	"gorm.io/gorm"
)

// AssessmentType defines the type of assessment
type AssessmentType string

const (
	AssessmentTypeQuiz       AssessmentType = "quiz"
	AssessmentTypeAssignment AssessmentType = "assignment"
)

// AssessmentStatus defines the status of an assessment
type AssessmentStatus string

const (
	AssessmentStatusDraft     AssessmentStatus = "draft"
	AssessmentStatusReleased  AssessmentStatus = "released"
	AssessmentStatusClosed    AssessmentStatus = "closed"
	AssessmentStatusGraded    AssessmentStatus = "graded"
)

// Assessment is the base interface for all types of assessments
type Assessment interface {
	GetID() uint
	GetTitle() string
	GetDescription() string
	GetType() AssessmentType
	GetStatus() AssessmentStatus
	GetStartTime() time.Time
	GetEndTime() time.Time
	GetCourseID() curriculumModel.Course
	GetInstructor() commonModel.Instructor
	GetSubmissions() []Submission
	Validate() error
}

// BaseAssessment is the concrete implementation of the Assessment interface
type BaseAssessment struct {
	gorm.Model
	Title       string
	Description string
	Type        AssessmentType
	Status      AssessmentStatus
	StartTime   time.Time
	EndTime     time.Time
	CourseID    curriculumModel.Course
	Instructor  commonModel.Instructor
	Submissions []Submission
}

func (a *BaseAssessment) GetID() uint {
	return a.ID
}

func (a *BaseAssessment) GetTitle() string {
	return a.Title
}

func (a *BaseAssessment) GetDescription() string {
	return a.Description
}

func (a *BaseAssessment) GetType() AssessmentType {
	return a.Type
}

func (a *BaseAssessment) GetStatus() AssessmentStatus {
	return a.Status
}

func (a *BaseAssessment) GetStartTime() time.Time {
	return a.StartTime
}

func (a *BaseAssessment) GetEndTime() time.Time {
	return a.EndTime
}

func (a *BaseAssessment) GetCourseID() curriculumModel.Course {
	return a.CourseID
}

func (a *BaseAssessment) GetInstructor() commonModel.Instructor {
	return a.Instructor
}

func (a *BaseAssessment) GetSubmissions() []Submission {
	return a.Submissions
}

func (a *BaseAssessment) Validate() error {
	if a.Title == "" {
		return fmt.Errorf("title cannot be empty")
	}
	if a.Description == "" {
		return fmt.Errorf("description cannot be empty")
	}
	if a.StartTime.After(a.EndTime) {
		return fmt.Errorf("start time cannot be after end time")
	}
	return nil
}

// SubmissionStatus defines the status of a submission
type SubmissionStatus string

const (
	SubmissionStatusDraft     SubmissionStatus = "draft"
	SubmissionStatusSubmitted SubmissionStatus = "submitted"
	SubmissionStatusGraded    SubmissionStatus = "graded"
)

// Submission represents a student's submission for an assessment
type Submission struct {
	gorm.Model
	Student     commonModel.Student
	Assessment  BaseAssessment
	Content     string
	Status      SubmissionStatus
	SubmittedAt time.Time
	Score       *float64
	Feedback    string
	GradedAt    *time.Time
	GradedBy    *commonModel.Instructor
} 