package controller

import (
	commonModel "ModEd/common/model"
	"ModEd/curriculum/model"

	"gorm.io/gorm"
)

type WIL struct {
	Connector *gorm.DB
}

func CreateWIL(connector *gorm.DB) *WIL {
	wil := WIL{Connector: connector}
	connector.AutoMigrate(&commonModel.Student{}, &model.WILProjectApplication{}, &model.WILProjectMembers{})

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
