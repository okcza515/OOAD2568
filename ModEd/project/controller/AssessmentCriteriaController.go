package controller

import (
	"ModEd/project/model"

	"gorm.io/gorm"
)

type IAssessmentCriteriaController interface {
	ListAllAssessmentCriterias() ([]model.AssessmentCriteria, error)
	RetrieveAssessmentCriteria(id uint) (*model.AssessmentCriteria, error)
	InsertAssessmentCriteria(AssessmentCriteria *model.AssessmentCriteria) error
	UpdateAssessmentCriteria(AssessmentCriteria *model.AssessmentCriteria) error
	DeleteAssessmentCriteria(id uint) error
}

type AssessmentCriteriaController struct {
	db *gorm.DB
}

func NewAssessmentCriteriaController(db *gorm.DB) IAssessmentCriteriaController {
	return &AssessmentCriteriaController{db: db}
}

func (c *AssessmentCriteriaController) ListAllAssessmentCriterias() ([]model.AssessmentCriteria, error) {
	var AssessmentCriterias []model.AssessmentCriteria
	err := c.db.Find(&AssessmentCriterias).Error
	return AssessmentCriterias, err
}

func (c *AssessmentCriteriaController) RetrieveAssessmentCriteria(id uint) (*model.AssessmentCriteria, error) {
	var AssessmentCriteria model.AssessmentCriteria
	if err := c.db.Where("id = ?", id).First(&AssessmentCriteria).Error; err != nil {
		return nil, err
	}
	return &AssessmentCriteria, nil
}

func (c *AssessmentCriteriaController) InsertAssessmentCriteria(AssessmentCriteria *model.AssessmentCriteria) error {
	return c.db.Create(AssessmentCriteria).Error
}

func (c *AssessmentCriteriaController) UpdateAssessmentCriteria(AssessmentCriteria *model.AssessmentCriteria) error {
	return c.db.Save(AssessmentCriteria).Error
}

func (c *AssessmentCriteriaController) DeleteAssessmentCriteria(id uint) error {
	return c.db.Where("id = ?", id).Delete(&model.AssessmentCriteria{}).Error
}
