package model

import (
	"fmt"
	"time"

	commonModel "ModEd/common/model"
	curriculumModel "ModEd/curriculum/model"

	"gorm.io/gorm"
)

type AssessmentType string

const (
	QuizType       AssessmentType = "quiz"
	AssignmentType AssessmentType = "assignment"
)

type AssessmentStatus string

const (
	StatusDraft     AssessmentStatus = "draft"
	StatusPublished AssessmentStatus = "published"
	StatusClosed    AssessmentStatus = "closed"
)

type Assessment struct {
	gorm.Model
	AssessmentId   uint
	Type           AssessmentType
	Title          string
	Description    string
	Released       bool
	StartDate      time.Time
	DueDate        time.Time
	Status         AssessmentStatus
	CourseId       curriculumModel.Course
	InstructorCode commonModel.Instructor
	FirstName      commonModel.Instructor
	LastName       commonModel.Instructor
	Submission     []AssessmentSubmission
	Observers      []AssessmentObserver
	State          AssessmentState
}

type AssessmentState interface {
	HandleSubmission(assessment *Assessment, submission *AssessmentSubmission) error
	HandleStatusChange(assessment *Assessment, newStatus AssessmentStatus) error
}

type AssessmentSubmission struct {
	gorm.Model
	StudentCode string
	FirstName   string
	LastName    string
	Email       string
	Answers     string
	Submitted   bool
	SubmittedAt time.Time
	Score       float64
	Feedback    string
}

type AssessmentObserver interface {
	OnStatusChanged(assessment *Assessment, oldStatus AssessmentStatus)
}

type AssessmentBuilder interface {
	SetTitle(title string) AssessmentBuilder
	SetDescription(description string) AssessmentBuilder
	SetDates(startDate, dueDate time.Time) AssessmentBuilder
	SetCourse(courseId curriculumModel.Course) AssessmentBuilder
	SetInstructor(instructorCode commonModel.Instructor) AssessmentBuilder
	Build() *Assessment
}

type ConcreteAssessmentBuilder struct {
	assessment *Assessment
}

func (b *ConcreteAssessmentBuilder) SetTitle(title string) AssessmentBuilder {
	b.assessment.Title = title
	return b
}

func (b *ConcreteAssessmentBuilder) SetDescription(description string) AssessmentBuilder {
	b.assessment.Description = description
	return b
}

func (b *ConcreteAssessmentBuilder) SetDates(startDate, dueDate time.Time) AssessmentBuilder {
	b.assessment.StartDate = startDate
	b.assessment.DueDate = dueDate
	return b
}

func (b *ConcreteAssessmentBuilder) SetCourse(courseId curriculumModel.Course) AssessmentBuilder {
	b.assessment.CourseId = courseId
	return b
}

func (b *ConcreteAssessmentBuilder) SetInstructor(instructorCode commonModel.Instructor) AssessmentBuilder {
	b.assessment.InstructorCode = instructorCode
	return b
}

func (b *ConcreteAssessmentBuilder) Build() *Assessment {
	return b.assessment
}

func NewAssessmentBuilder(assessmentType AssessmentType) AssessmentBuilder {
	return &ConcreteAssessmentBuilder{
		assessment: &Assessment{
			Type:      assessmentType,
			Status:    StatusDraft,
			Observers: make([]AssessmentObserver, 0),
		},
	}
}

type SubmissionStrategy interface {
	ValidateSubmission(submission *AssessmentSubmission) error
	ProcessSubmission(submission *AssessmentSubmission) error
}

type QuizSubmissionStrategy struct{}

type AssignmentSubmissionStrategy struct{}

type AssessmentDecorator interface {
	GetAssessment() *Assessment
	GetDescription() string
}

type TimedAssessmentDecorator struct {
	assessment *Assessment
	timeLimit  time.Duration
}

func NewTimedAssessmentDecorator(assessment *Assessment, timeLimit time.Duration) AssessmentDecorator {
	return &TimedAssessmentDecorator{
		assessment: assessment,
		timeLimit:  timeLimit,
	}
}

func (d *TimedAssessmentDecorator) GetAssessment() *Assessment {
	return d.assessment
}

func (d *TimedAssessmentDecorator) GetDescription() string {
	return fmt.Sprintf("%s (Time Limit: %v)", d.assessment.Description, d.timeLimit)
}
