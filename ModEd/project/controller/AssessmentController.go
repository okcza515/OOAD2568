package controller

import (
	"ModEd/project/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IAssessmentController interface {
	ListAllAssessments() ([]model.Assessment, error)
	RetrieveAssessment(id uuid.UUID) (*model.Assessment, error)
	InsertAssessment(assessment *model.Assessment) error
	UpdateAssessment(assessment *model.Assessment) error
	DeleteAssessment(id uuid.UUID) error
}

type AssessmentController struct {
	db *gorm.DB
}

func NewAssessmentController(db *gorm.DB) IAssessmentController {
	return &AssessmentController{db: db}
}

func (c *AssessmentController) ListAllAssessments() ([]model.Assessment, error) {
	var assessments []model.Assessment
	err := c.db.Preload("AssessmentCriteria").
		Preload("ScoreAssessmentAdvisor").
		Preload("ScoreAssessmentComittee").
		Preload("SeniorProject").
		Find(&assessments).Error
	return assessments, err
}

func (c *AssessmentController) RetrieveAssessment(id uuid.UUID) (*model.Assessment, error) {
	var assessment model.Assessment
	if err := c.db.Preload("AssessmentCriteria").
		Preload("ScoreAssessmentAdvisor").
		Preload("ScoreAssessmentComittee").
		Preload("SeniorProject").
		Where("assessment_id = ?", id).
		First(&assessment).Error; err != nil {
		return nil, err
	}
	return &assessment, nil
}

func (c *AssessmentController) InsertAssessment(assessment *model.Assessment) error {
	return c.db.Create(assessment).Error
}

func (c *AssessmentController) UpdateAssessment(assessment *model.Assessment) error {
	return c.db.Save(assessment).Error
}

func (c *AssessmentController) DeleteAssessment(id uuid.UUID) error {
	return c.db.Where("assessment_id = ?", id).Delete(&model.Assessment{}).Error
}