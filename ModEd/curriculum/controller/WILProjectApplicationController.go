package controller

import (
	"ModEd/core"
	model "ModEd/curriculum/model"

	"gorm.io/gorm"
)

type WILProjectApplicationController struct {
	connector *gorm.DB
	*core.BaseController
}

type WILProjectApplicationControllerInterface interface {
	RegisterWILProjectsApplication(projects []core.RecordInterface)
	Insert(data core.RecordInterface) error
	UpdateByID(data core.RecordInterface) error
	RetrieveByID(id uint, preloads ...string) (*core.RecordInterface, error)
	DeleteByID(id uint) error
	ListPagination(condition map[string]interface{}, page int, pageSize int)
}

func CreateWILProjectApplicationController(connector *gorm.DB) *WILProjectApplicationController {
	wil := WILProjectApplicationController{connector: connector, BaseController: core.NewBaseController("WILProjectApplication", connector)}
	return &wil
}

func (repo WILProjectApplicationController) RegisterWILProjectApplication(application *model.WILProjectApplication) {
	repo.connector.Create(application)
}

func (repo WILProjectApplicationController) RegisterWILProjectApplications(applications []*model.WILProjectApplication) {
	for _, application := range applications {
		repo.connector.Create(application)
	}
}

func (repo WILProjectApplicationController) GetAllWILProjectApplications() ([]*model.WILProjectApplication, error) {
	applications := []*model.WILProjectApplication{}
	result := repo.connector.Find(&applications)
	return applications, result.Error
}

func (repo WILProjectApplicationController) GetWILProjectApplicationByID(id uint) (*model.WILProjectApplication, error) {
	application := &model.WILProjectApplication{}
	result := repo.connector.Where("WILProjectApplicationId = ?", id).First(application)
	return application, result.Error
}

func (repo WILProjectApplicationController) UpdateWILProjectApplication(application *model.WILProjectApplication) error {
	result := repo.connector.Save(application)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
