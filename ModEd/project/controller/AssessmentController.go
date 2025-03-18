package controller

import (
	"ModEd/project/model"

	"gorm.io/gorm"
)

type IAssessmentController interface {
	ListAllAssessments() ([]model.Assessment, error)
	RetrieveAssessment(id uint) (*model.Assessment, error)
	InsertAssessment(assessment *model.Assessment) error
	UpdateAssessment(assessment *model.Assessment) error
	DeleteAssessment(id uint) error
}

type AssessmentController struct {
	db *gorm.DB
}

func NewAssessmentController(db *gorm.DB) IAssessmentController {
	return &AssessmentController{db: db}
}

func (c *AssessmentController) ListAllAssessments() ([]model.Assessment, error) {
	var assessments []model.Assessment
	err := c.db.
		Find(&assessments).Error
	return assessments, err
}

func (c *AssessmentController) RetrieveAssessment(id uint) (*model.Assessment, error) {
	var assessment model.Assessment
	if err := c.db.
		Where("id = ?", id).
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

func (c *AssessmentController) DeleteAssessment(id uint) error {
	return c.db.Where("id = ?", id).Delete(&model.Assessment{}).Error
}
