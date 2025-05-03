package controller

import (
	"errors"
	"time"

	"ModEd/eval/model"

	"gorm.io/gorm"
)

// AssessmentController handles operations for all types of assessments
type AssessmentController interface {
	GetAll() ([]model.Assessment, error)
	GetByID(id uint) (model.Assessment, error)
	Create(assessment model.Assessment) error
	Update(assessment model.Assessment) error
	Delete(id uint) error
	Submit(assessmentID uint, studentID string, content string) error
	Grade(assessmentID uint, studentID string, score float64, feedback string) error
}

type assessmentController struct {
	db *gorm.DB
}

// NewAssessmentController creates a new assessment controller
func NewAssessmentController(db *gorm.DB) AssessmentController {
	return &assessmentController{db: db}
}

func (c *assessmentController) GetAll() ([]model.Assessment, error) {
	var quizzes []*model.Quiz
	var assignments []*model.Assignment
	var assessments []model.Assessment

	// Get all quizzes
	if err := c.db.Preload("Submissions").Find(&quizzes).Error; err != nil {
		return nil, err
	}
	for _, q := range quizzes {
		assessments = append(assessments, q)
	}

	// Get all assignments
	if err := c.db.Preload("Submissions").Find(&assignments).Error; err != nil {
		return nil, err
	}
	for _, a := range assignments {
		assessments = append(assessments, a)
	}

	return assessments, nil
}

func (c *assessmentController) GetByID(id uint) (model.Assessment, error) {
	// Try to get as quiz first
	var quiz model.Quiz
	if err := c.db.Preload("Submissions").First(&quiz, id).Error; err == nil {
		return &quiz, nil
	}

	// If not a quiz, try as assignment
	var assignment model.Assignment
	if err := c.db.Preload("Submissions").First(&assignment, id).Error; err == nil {
		return &assignment, nil
	}

	return nil, errors.New("assessment not found")
}

func (c *assessmentController) Create(assessment model.Assessment) error {
	if err := assessment.Validate(); err != nil {
		return err
	}

	switch a := assessment.(type) {
	case *model.Quiz:
		return c.db.Create(a).Error
	case *model.Assignment:
		return c.db.Create(a).Error
	default:
		return errors.New("unsupported assessment type")
	}
}

func (c *assessmentController) Update(assessment model.Assessment) error {
	if err := assessment.Validate(); err != nil {
		return err
	}

	switch a := assessment.(type) {
	case *model.Quiz:
		return c.db.Save(a).Error
	case *model.Assignment:
		return c.db.Save(a).Error
	default:
		return errors.New("unsupported assessment type")
	}
}

func (c *assessmentController) Delete(id uint) error {
	// Try to delete as quiz first
	if err := c.db.Delete(&model.Quiz{}, id).Error; err == nil {
		return nil
	}

	// If not a quiz, try as assignment
	if err := c.db.Delete(&model.Assignment{}, id).Error; err == nil {
		return nil
	}

	return errors.New("assessment not found")
}

func (c *assessmentController) Submit(assessmentID uint, studentID string, content string) error {
	assessment, err := c.GetByID(assessmentID)
	if err != nil {
		return errors.New("assessment not found")
	}

	// Check if assessment is still open
	if time.Now().After(assessment.GetEndTime()) {
		return errors.New("assessment is closed")
	}

	submission := model.Submission{
		Student:     model.Student{StudentCode: studentID},
		AssessmentID: assessmentID,
		Content:     content,
		Status:      model.SubmissionStatusSubmitted,
		SubmittedAt: time.Now(),
	}

	return c.db.Create(&submission).Error
}

func (c *assessmentController) Grade(assessmentID uint, studentID string, score float64, feedback string) error {
	var submission model.Submission
	if err := c.db.Where("assessment_id = ? AND student_id = ?", assessmentID, studentID).First(&submission).Error; err != nil {
		return errors.New("submission not found")
	}

	submission.Score = &score
	submission.Feedback = feedback
	submission.Status = model.SubmissionStatusGraded
	now := time.Now()
	submission.GradedAt = &now

	return c.db.Save(&submission).Error
} 