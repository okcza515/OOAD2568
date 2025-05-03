package model

import (
	"time"

	commonModel "ModEd/common/model"
	curriculumModel "ModEd/curriculum/model"
)

// AssessmentFactory is responsible for creating different types of assessments
type AssessmentFactory struct{}

// NewAssessmentFactory creates a new assessment factory
func NewAssessmentFactory() *AssessmentFactory {
	return &AssessmentFactory{}
}

// CreateAssessment creates a new assessment based on the type
func (f *AssessmentFactory) CreateAssessment(
	assessmentType AssessmentType,
	title string,
	description string,
	courseID curriculumModel.Course,
	instructor commonModel.Instructor,
	startTime time.Time,
	endTime time.Time,
) (Assessment, error) {
	switch assessmentType {
	case AssessmentTypeQuiz:
		quiz := NewQuiz()
		quiz.Title = title
		quiz.Description = description
		quiz.CourseID = courseID
		quiz.Instructor = instructor
		quiz.StartTime = startTime
		quiz.EndTime = endTime
		return quiz, nil

	case AssessmentTypeAssignment:
		assignment := NewAssignment()
		assignment.Title = title
		assignment.Description = description
		assignment.CourseID = courseID
		assignment.Instructor = instructor
		assignment.StartTime = startTime
		assignment.EndTime = endTime
		return assignment, nil

	default:
		return nil, fmt.Errorf("unsupported assessment type: %s", assessmentType)
	}
}

// CreateQuiz creates a new quiz with the given parameters
func (f *AssessmentFactory) CreateQuiz(
	title string,
	description string,
	courseID curriculumModel.Course,
	instructor commonModel.Instructor,
	startTime time.Time,
	endTime time.Time,
	timeLimit time.Duration,
	maxAttempts int,
	showAnswers bool,
	randomize bool,
) (*Quiz, error) {
	quiz := NewQuiz()
	quiz.Title = title
	quiz.Description = description
	quiz.CourseID = courseID
	quiz.Instructor = instructor
	quiz.StartTime = startTime
	quiz.EndTime = endTime
	quiz.TimeLimit = timeLimit
	quiz.MaxAttempts = maxAttempts
	quiz.ShowAnswers = showAnswers
	quiz.Randomize = randomize

	if err := quiz.Validate(); err != nil {
		return nil, err
	}

	return quiz, nil
}

// CreateAssignment creates a new assignment with the given parameters
func (f *AssessmentFactory) CreateAssignment(
	title string,
	description string,
	courseID curriculumModel.Course,
	instructor commonModel.Instructor,
	startTime time.Time,
	endTime time.Time,
	maxFileSize int64,
	allowedTypes []string,
	groupSize int,
	isGroup bool,
) (*Assignment, error) {
	assignment := NewAssignment()
	assignment.Title = title
	assignment.Description = description
	assignment.CourseID = courseID
	assignment.Instructor = instructor
	assignment.StartTime = startTime
	assignment.EndTime = endTime
	assignment.MaxFileSize = maxFileSize
	assignment.AllowedTypes = allowedTypes
	assignment.GroupSize = groupSize
	assignment.IsGroup = isGroup

	if err := assignment.Validate(); err != nil {
		return nil, err
	}

	return assignment, nil
} 