package controller

import (
	"ModEd/project/model"

	"gorm.io/gorm"
)

type ISeniorProjectController interface {
	ListAllSeniorProjects() ([]model.SeniorProject, error)
	RetrieveSeniorProject(id uint) (*model.SeniorProject, error)
	InsertSeniorProject(SeniorProject *model.SeniorProject) error
	UpdateSeniorProject(SeniorProject *model.SeniorProject) error
	DeleteSeniorProject(id uint) error
}

type SeniorProjectController struct {
	db *gorm.DB
}

func NewSeniorProjectController(db *gorm.DB) ISeniorProjectController {
	return &SeniorProjectController{db: db}
}

func (c *SeniorProjectController) ListAllSeniorProjects() ([]model.SeniorProject, error) {
	var SeniorProjects []model.SeniorProject
	err := c.db.Find(&SeniorProjects).Error
	return SeniorProjects, err
}

func (c *SeniorProjectController) RetrieveSeniorProject(id uint) (*model.SeniorProject, error) {
	var SeniorProject model.SeniorProject
	if err := c.db.Where("id = ?", id).First(&SeniorProject).Error; err != nil {
		return nil, err
	}
	return &SeniorProject, nil
}

func (c *SeniorProjectController) InsertSeniorProject(SeniorProject *model.SeniorProject) error {
	return c.db.Create(SeniorProject).Error
}

func (c *SeniorProjectController) UpdateSeniorProject(SeniorProject *model.SeniorProject) error {
	return c.db.Save(SeniorProject).Error
}

func (c *SeniorProjectController) DeleteSeniorProject(id uint) error {
	return c.db.Where("id = ?", id).Delete(&model.SeniorProject{}).Error
}
