package model

import (
	"fmt"
	"time"

	commonModel "ModEd/common/model"

	curriculumModel "ModEd/curriculum/model"

	"gorm.io/gorm"
)

// Quiz represents a quiz assessment
type Quiz struct {
	BaseAssessment
	TimeLimit    time.Duration
	Questions    []Question
	MaxAttempts  int
	ShowAnswers  bool
	Randomize    bool
}

// NewQuiz creates a new quiz with default values
func NewQuiz() *Quiz {
	return &Quiz{
		BaseAssessment: BaseAssessment{
			Type:   AssessmentTypeQuiz,
			Status: AssessmentStatusDraft,
		},
		MaxAttempts: 1,
		ShowAnswers: false,
		Randomize:   false,
	}
}

// Validate implements additional quiz-specific validation
func (q *Quiz) Validate() error {
	if err := q.BaseAssessment.Validate(); err != nil {
		return err
	}
	if q.TimeLimit < 0 {
		return fmt.Errorf("time limit cannot be negative")
	}
	if q.MaxAttempts < 1 {
		return fmt.Errorf("max attempts must be at least 1")
	}
	return nil
}

// AddQuestion adds a question to the quiz
func (q *Quiz) AddQuestion(question Question) {
	q.Questions = append(q.Questions, question)
}

// RemoveQuestion removes a question from the quiz
func (q *Quiz) RemoveQuestion(questionID uint) {
	for i, question := range q.Questions {
		if question.ID == questionID {
			q.Questions = append(q.Questions[:i], q.Questions[i+1:]...)
			break
		}
	}
}

// GetQuestions returns all questions in the quiz
func (q *Quiz) GetQuestions() []Question {
	return q.Questions
}

// GetTimeLimit returns the time limit for the quiz
func (q *Quiz) GetTimeLimit() time.Duration {
	return q.TimeLimit
}

// GetMaxAttempts returns the maximum number of attempts allowed
func (q *Quiz) GetMaxAttempts() int {
	return q.MaxAttempts
}

// IsShowAnswers returns whether answers should be shown after submission
func (q *Quiz) IsShowAnswers() bool {
	return q.ShowAnswers
}

// IsRandomize returns whether questions should be randomized
func (q *Quiz) IsRandomize() bool {
	return q.Randomize
}

type QuizSubmission struct {
	gorm.Model
	StudentCode commonModel.Student
	FirstName   commonModel.Student
	LastName    commonModel.Student
	Email       commonModel.Student
	Answers     string
	Submitted   bool
	SubmittedAt time.Time
}
