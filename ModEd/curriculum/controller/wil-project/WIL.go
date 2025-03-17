package controller

import (
	commonModel "ModEd/common/model"
	model "ModEd/curriculum/model/wil-project"

	"gorm.io/gorm"
)

type WIL struct {
	Connector *gorm.DB
}

func CreateWILProjectApplicationController(connector *gorm.DB) *WIL {
	wil := WIL{Connector: connector}
	connector.AutoMigrate(&commonModel.Student{}, &model.WILProjectApplication{}, &model.WILProjectMember{})

	return &wil
}

func (repo WIL) RegisterWILProjectApplications(applications []*model.WILProjectApplication) {
	for _, application := range applications {
		repo.Connector.Create(application)
	}
}

func (repo WIL) GetAllWILProjectApplications() ([]*model.WILProjectApplication, error) {
	applications := []*model.WILProjectApplication{}
	result := repo.Connector.Find(&applications)
	return applications, result.Error
}

func (repo WIL) GetWILProjectApplicationByID(id uint) (*model.WILProjectApplication, error) {
	application := &model.WILProjectApplication{}
	result := repo.Connector.Where("WILProjectApplicationId = ?", id).First(application)
	return application, result.Error
}

func (repo WIL) UpdateWILProjectApplication(application *model.WILProjectApplication) error {
	result := repo.Connector.Save(application)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
