package controller

import (
	"ModEd/core"
	model "ModEd/curriculum/model/Internship"

	"gorm.io/gorm"
)

type InternshipApplicationController struct {
	*core.BaseController
	Connector *gorm.DB
}

func CreateInternshipApplicationController(connector *gorm.DB) *InternshipApplicationController {
	return &InternshipApplicationController{
		Connector:      connector,
		BaseController: core.NewBaseController("InternshipApplication", connector),
	}
}

func (repo InternshipApplicationController) RegisterInternshipApplications(applications []*model.InternshipApplication) error {
	for _, application := range applications {
		application.InternshipReport = model.InternshipReport{}
		application.SupervisorReview = model.SupervisorReview{}
		application.InternshipSchedule = model.InternshipSchedule{}

		if err := repo.Connector.Create(application).Error; err != nil {
			return err
		}
	}
	return nil
}

func (repo InternshipApplicationController) GetAllInternshipApplications() ([]*model.InternshipApplication, error) {
	applications := []*model.InternshipApplication{}
	result := repo.Connector.Preload("InternshipReport").
		Preload("SupervisorReview").
		Preload("InternshipSchedule").
		Find(&applications)
	return applications, result.Error
}

func (repo InternshipApplicationController) GetInternshipApplicationByID(id uint) (*model.InternshipApplication, error) {
	application := &model.InternshipApplication{}
	result := repo.Connector.Preload("InternshipReport").
		Preload("SupervisorReview").
		Preload("InternshipSchedule").
		First(application, id)
	return application, result.Error
}
