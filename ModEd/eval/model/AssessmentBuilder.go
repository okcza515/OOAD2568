package model

import (
	"time"

	commonModel "ModEd/common/model"
	curriculumModel "ModEd/curriculum/model"
)

// AssessmentBuilder is a builder for creating assessments
type AssessmentBuilder struct {
	assessmentType AssessmentType
	title          string
	description    string
	courseID       curriculumModel.Course
	instructor     commonModel.Instructor
	startTime      time.Time
	endTime        time.Time
	timeLimit      time.Duration
	maxAttempts    int
	showAnswers    bool
	randomize      bool
	maxFileSize    int64
	allowedTypes   []string
	groupSize      int
	isGroup        bool
}

// NewAssessmentBuilder creates a new assessment builder
func NewAssessmentBuilder() *AssessmentBuilder {
	return &AssessmentBuilder{
		timeLimit:   30 * time.Minute, // Default 30 minutes
		maxAttempts: 1,
		maxFileSize: 10 * 1024 * 1024, // Default 10MB
		groupSize:   1,
	}
}

// SetType sets the assessment type
func (b *AssessmentBuilder) SetType(assessmentType AssessmentType) *AssessmentBuilder {
	b.assessmentType = assessmentType
	return b
}

// SetTitle sets the assessment title
func (b *AssessmentBuilder) SetTitle(title string) *AssessmentBuilder {
	b.title = title
	return b
}

// SetDescription sets the assessment description
func (b *AssessmentBuilder) SetDescription(description string) *AssessmentBuilder {
	b.description = description
	return b
}

// SetCourseID sets the course ID
func (b *AssessmentBuilder) SetCourseID(courseID curriculumModel.Course) *AssessmentBuilder {
	b.courseID = courseID
	return b
}

// SetInstructor sets the instructor
func (b *AssessmentBuilder) SetInstructor(instructor commonModel.Instructor) *AssessmentBuilder {
	b.instructor = instructor
	return b
}

// SetTimeRange sets the start and end time
func (b *AssessmentBuilder) SetTimeRange(startTime, endTime time.Time) *AssessmentBuilder {
	b.startTime = startTime
	b.endTime = endTime
	return b
}

// SetQuizOptions sets quiz-specific options
func (b *AssessmentBuilder) SetQuizOptions(timeLimit time.Duration, maxAttempts int, showAnswers, randomize bool) *AssessmentBuilder {
	b.timeLimit = timeLimit
	b.maxAttempts = maxAttempts
	b.showAnswers = showAnswers
	b.randomize = randomize
	return b
}

// SetAssignmentOptions sets assignment-specific options
func (b *AssessmentBuilder) SetAssignmentOptions(maxFileSize int64, allowedTypes []string, groupSize int, isGroup bool) *AssessmentBuilder {
	b.maxFileSize = maxFileSize
	b.allowedTypes = allowedTypes
	b.groupSize = groupSize
	b.isGroup = isGroup
	return b
}

// Build creates the assessment based on the builder configuration
func (b *AssessmentBuilder) Build() (Assessment, error) {
	factory := NewAssessmentFactory()

	switch b.assessmentType {
	case AssessmentTypeQuiz:
		return factory.CreateQuiz(
			b.title,
			b.description,
			b.courseID,
			b.instructor,
			b.startTime,
			b.endTime,
			b.timeLimit,
			b.maxAttempts,
			b.showAnswers,
			b.randomize,
		)

	case AssessmentTypeAssignment:
		return factory.CreateAssignment(
			b.title,
			b.description,
			b.courseID,
			b.instructor,
			b.startTime,
			b.endTime,
			b.maxFileSize,
			b.allowedTypes,
			b.groupSize,
			b.isGroup,
		)

	default:
		return nil, fmt.Errorf("unsupported assessment type: %s", b.assessmentType)
	}
} 