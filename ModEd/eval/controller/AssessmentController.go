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
	var assessments []model.Assessment
	if err := c.db.Preload("Submissions").Find(&assessments).Error; err != nil {
		return nil, err
	}
	return assessments, nil
}

func (c *assessmentController) GetByID(id uint) (model.Assessment, error) {
	var assessment model.Assessment
	if err := c.db.Preload("Submissions").First(&assessment, id).Error; err != nil {
		return nil, err
	}
	return assessment, nil
}

func (c *assessmentController) Create(assessment model.Assessment) error {
	if err := assessment.Validate(); err != nil {
		return err
	}
	return c.db.Create(assessment).Error
}

func (c *assessmentController) Update(assessment model.Assessment) error {
	if err := assessment.Validate(); err != nil {
		return err
	}
	return c.db.Save(assessment).Error
}

func (c *assessmentController) Delete(id uint) error {
	var assessment model.Assessment
	if err := c.db.First(&assessment, id).Error; err != nil {
		return errors.New("assessment not found")
	}
	return c.db.Delete(&assessment).Error
}

func (c *assessmentController) Submit(assessmentID uint, studentID string, content string) error {
	var assessment model.Assessment
	if err := c.db.First(&assessment, assessmentID).Error; err != nil {
		return errors.New("assessment not found")
	}

	// Check if assessment is still open
	if time.Now().After(assessment.GetEndTime()) {
		return errors.New("assessment is closed")
	}

	submission := model.Submission{
		Student:     model.Student{StudentCode: studentID},
		Assessment:  assessment.(*model.BaseAssessment),
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