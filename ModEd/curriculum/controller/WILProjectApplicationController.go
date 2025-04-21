package controller

import (
	commonModel "ModEd/common/model"
	model "ModEd/curriculum/model"

	"gorm.io/gorm"
)

type WILProjectApplicationController struct {
	Connector *gorm.DB
}

func CreateWILProjectApplicationController(connector *gorm.DB) *WILProjectApplicationController {
	wil := WILProjectApplicationController{Connector: connector}
	connector.AutoMigrate(&commonModel.Student{}, &model.WILProjectApplication{}, &model.WILProjectMember{})

	return &wil
}

func (repo WILProjectApplicationController) RegisterWILProjectApplications(applications []*model.WILProjectApplication) {
	for _, application := range applications {
		repo.Connector.Create(application)
	}
}

func (repo WILProjectApplicationController) GetAllWILProjectApplications() ([]*model.WILProjectApplication, error) {
	applications := []*model.WILProjectApplication{}
	result := repo.Connector.Find(&applications)
	return applications, result.Error
}

func (repo WILProjectApplicationController) GetWILProjectApplicationByID(id uint) (*model.WILProjectApplication, error) {
	application := &model.WILProjectApplication{}
	result := repo.Connector.Where("WILProjectApplicationId = ?", id).First(application)
	return application, result.Error
}

func (repo WILProjectApplicationController) UpdateWILProjectApplication(application *model.WILProjectApplication) error {
	result := repo.Connector.Save(application)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
