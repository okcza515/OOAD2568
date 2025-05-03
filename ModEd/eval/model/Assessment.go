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
	QuizType       AssessmentType = "quiz"
	AssignmentType AssessmentType = "assignment"
)

// AssessmentStatus defines possible assessment statuses
type AssessmentStatus string

const (
	StatusDraft     AssessmentStatus = "draft"
	StatusPublished AssessmentStatus = "published"
	StatusClosed    AssessmentStatus = "closed"
)

// Assessment represents a combined model for both quizzes and assignments
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
	observers      []AssessmentObserver
}

// AssessmentSubmission represents a student's submission for either a quiz or assignment
type AssessmentSubmission struct {
	gorm.Model
	StudentCode commonModel.Student
	FirstName   commonModel.Student
	LastName    commonModel.Student
	Email       commonModel.Student
	Answers     string
	Submitted   bool
	SubmittedAt time.Time
	Score       float64
	Feedback    string
}

// AssessmentObserver interface for implementing observer pattern
type AssessmentObserver interface {
	OnStatusChanged(assessment *Assessment, oldStatus AssessmentStatus)
}

// AssessmentBuilder interface for implementing builder pattern
type AssessmentBuilder interface {
	SetTitle(title string) AssessmentBuilder
	SetDescription(description string) AssessmentBuilder
	SetDates(startDate, dueDate time.Time) AssessmentBuilder
	SetCourse(courseId curriculumModel.Course) AssessmentBuilder
	SetInstructor(instructorCode commonModel.Instructor) AssessmentBuilder
	Build() *Assessment
}

// ConcreteAssessmentBuilder implements AssessmentBuilder
type ConcreteAssessmentBuilder struct {
	assessment *Assessment
}

// NewAssessmentBuilder creates a new builder instance
func NewAssessmentBuilder(assessmentType AssessmentType) AssessmentBuilder {
	return &ConcreteAssessmentBuilder{
		assessment: &Assessment{
			Type:   assessmentType,
			Status: StatusDraft,
		},
	}
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

// AssessmentFactory creates different types of assessments
type AssessmentFactory struct{}

func (f *AssessmentFactory) CreateAssessment(assessmentType AssessmentType) AssessmentBuilder {
	return NewAssessmentBuilder(assessmentType)
}

// AddObserver adds an observer to the assessment
func (a *Assessment) AddObserver(observer AssessmentObserver) {
	a.observers = append(a.observers, observer)
}

// SetStatus updates the assessment status and notifies observers
func (a *Assessment) SetStatus(newStatus AssessmentStatus) {
	oldStatus := a.Status
	a.Status = newStatus
	a.notifyObservers(oldStatus)
}

// notifyObservers notifies all observers about status change
func (a *Assessment) notifyObservers(oldStatus AssessmentStatus) {
	for _, observer := range a.observers {
		observer.OnStatusChanged(a, oldStatus)
	}
}

// SubmissionStrategy interface for different submission handling strategies
type SubmissionStrategy interface {
	ValidateSubmission(submission *AssessmentSubmission) error
	ProcessSubmission(submission *AssessmentSubmission) error
}

// QuizSubmissionStrategy implements SubmissionStrategy for quizzes
type QuizSubmissionStrategy struct{}

func (s *QuizSubmissionStrategy) ValidateSubmission(submission *AssessmentSubmission) error {
	// Implement quiz-specific validation
	return nil
}

func (s *QuizSubmissionStrategy) ProcessSubmission(submission *AssessmentSubmission) error {
	// Implement quiz-specific processing
	return nil
}

// AssignmentSubmissionStrategy implements SubmissionStrategy for assignments
type AssignmentSubmissionStrategy struct{}

func (s *AssignmentSubmissionStrategy) ValidateSubmission(submission *AssessmentSubmission) error {
	// Implement assignment-specific validation
	return nil
}

func (s *AssignmentSubmissionStrategy) ProcessSubmission(submission *AssessmentSubmission) error {
	// Implement assignment-specific processing
	return nil
}
