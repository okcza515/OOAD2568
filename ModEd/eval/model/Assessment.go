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
	observers      []AssessmentObserver
	state          AssessmentState
}

type AssessmentState interface {
	HandleSubmission(assessment *Assessment, submission *AssessmentSubmission) error
	HandleStatusChange(assessment *Assessment, newStatus AssessmentStatus) error
}

type DraftState struct{}

func (s *DraftState) HandleSubmission(assessment *Assessment, submission *AssessmentSubmission) error {
	return fmt.Errorf("cannot submit to a draft assessment")
}

func (s *DraftState) HandleStatusChange(assessment *Assessment, newStatus AssessmentStatus) error {
	if newStatus == StatusPublished {
		assessment.Status = newStatus
		assessment.state = &PublishedState{}
		return nil
	}
	return fmt.Errorf("invalid status transition from draft")
}

type PublishedState struct{}

func (s *PublishedState) HandleSubmission(assessment *Assessment, submission *AssessmentSubmission) error {
	if time.Now().After(assessment.DueDate) {
		return fmt.Errorf("submission deadline has passed")
	}
	return nil
}

func (s *PublishedState) HandleStatusChange(assessment *Assessment, newStatus AssessmentStatus) error {
	if newStatus == StatusClosed {
		assessment.Status = newStatus
		assessment.state = &ClosedState{}
		return nil
	}
	return fmt.Errorf("invalid status transition from published")
}

type ClosedState struct{}

func (s *ClosedState) HandleSubmission(assessment *Assessment, submission *AssessmentSubmission) error {
	return fmt.Errorf("cannot submit to a closed assessment")
}

func (s *ClosedState) HandleStatusChange(assessment *Assessment, newStatus AssessmentStatus) error {
	return fmt.Errorf("cannot change status of a closed assessment")
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
			state:     &DraftState{},
			observers: make([]AssessmentObserver, 0),
		},
	}
}

type AssessmentFactory struct{}

func (f *AssessmentFactory) CreateAssessment(assessmentType AssessmentType) AssessmentBuilder {
	return NewAssessmentBuilder(assessmentType)
}

func (a *Assessment) AddObserver(observer AssessmentObserver) {
	a.observers = append(a.observers, observer)
}

func (a *Assessment) SetStatus(newStatus AssessmentStatus) {
	oldStatus := a.Status
	if err := a.state.HandleStatusChange(a, newStatus); err == nil {
		a.notifyObservers(oldStatus)
	}
}

func (a *Assessment) Submit(submission *AssessmentSubmission) error {
	return a.state.HandleSubmission(a, submission)
}

func (a *Assessment) notifyObservers(oldStatus AssessmentStatus) {
	for _, observer := range a.observers {
		observer.OnStatusChanged(a, oldStatus)
	}
}

type SubmissionStrategy interface {
	ValidateSubmission(submission *AssessmentSubmission) error
	ProcessSubmission(submission *AssessmentSubmission) error
}

type QuizSubmissionStrategy struct{}

func (s *QuizSubmissionStrategy) ValidateSubmission(submission *AssessmentSubmission) error {
	if submission.Answers == "" {
		return fmt.Errorf("quiz answers cannot be empty")
	}
	return nil
}

func (s *QuizSubmissionStrategy) ProcessSubmission(submission *AssessmentSubmission) error {
	return nil
}

type AssignmentSubmissionStrategy struct{}

func (s *AssignmentSubmissionStrategy) ValidateSubmission(submission *AssessmentSubmission) error {
	if submission.Answers == "" {
		return fmt.Errorf("assignment submission cannot be empty")
	}
	return nil
}

func (s *AssignmentSubmissionStrategy) ProcessSubmission(submission *AssessmentSubmission) error {
	return nil
}

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
