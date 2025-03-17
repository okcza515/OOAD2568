package controller

import (
	"ModEd/curriculum/model"
	"time"

	"gorm.io/gorm"
)

type WILProjectController struct {
	Connector *gorm.DB
}

func CreateWILProjectController(connector *gorm.DB) *WILProjectController {
	connector.AutoMigrate(&model.WILProject{})
	return &WILProjectController{Connector: connector}
}

func (repo WILProjectController) RegisterWILProjects(projects []*model.WILProject) {
	for _, project := range projects {
		repo.Connector.Create(project)
	}
}

func (repo WILProjectController) CreateWILProject(project *model.WILProject) error {
	result := repo.Connector.Create(project)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo WILProjectController) GetAllWILProjects() ([]*model.WILProject, error) {
	projects := []*model.WILProject{}
	result := repo.Connector.Find(&projects, "DeletedAt IS NULL")
	if result.Error != nil {
		return nil, result.Error
	}
	return projects, nil
}

func (repo WILProjectController) GetWILProjectById(id string) (*model.WILProject, error) {
	project := &model.WILProject{}
	result := repo.Connector.Where("WilProjectId = ?", id).First(project)
	if result.Error != nil {
		return nil, result.Error
	}
	return project, nil
}

func (repo WILProjectController) UpdateWILProject(project *model.WILProject) error {
	result := repo.Connector.Save(project)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo WILProjectController) DeleteWILProject(id string) error {
	result := repo.Connector.Model(&model.WILProject{}).Where("WilProjectId = ?", id).Update("DeletedAt", gorm.DeletedAt{Time: time.Now(), Valid: true})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
