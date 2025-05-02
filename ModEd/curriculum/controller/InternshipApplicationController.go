package controller

import (
    "ModEd/core"
    model "ModEd/curriculum/model"
    "fmt"

    "gorm.io/gorm"
)

type InternshipApplicationController struct {
    *core.BaseController[model.InternshipApplication]
    Connector *gorm.DB
}

type InternshipApplicationControllerInterface interface {
	RegisterInternshipApplications(applications []*model.InternshipApplication) error
	GetAllInternshipApplications() ([]*model.InternshipApplication, error)
	GetInternshipApplicationByID(id uint) (*model.InternshipApplication, error)
	GetApplicationStatusByID(id uint) (string, error)
	DeleteApplicationByID(id uint) error
}

func NewInternshipApplicationController(connector *gorm.DB) *InternshipApplicationController {
	return &InternshipApplicationController{
			Connector:      connector,
			BaseController: core.NewBaseController[model.InternshipApplication](connector),
	}
}

func (repo InternshipApplicationController) RegisterInternshipApplications(applications []*model.InternshipApplication) error {
	for _, application := range applications {
			application.InternshipReport = model.InternshipReport{}
			application.SupervisorReview = model.SupervisorReview{}

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

func (repo InternshipApplicationController) GetInternshipApplicationByID(id string) (*model.InternshipApplication, error) {
	application := &model.InternshipApplication{}
	result := repo.Connector.Preload("InternshipReport").
			Preload("SupervisorReview").
			Preload("InternshipSchedule").
			First(application, id)
	return application, result.Error
}

func (repo InternshipApplicationController) GetApplicationStatusByID(id string) (string, error) {
	var application model.InternshipApplication
	if err := repo.Connector.Select("approval_advisor_status", "approval_company_status").
			Where("id = ?", id).
			First(&application).Error; err != nil {
			return "", fmt.Errorf("failed to find application with ID %s: %w", id, err)
	}

	return fmt.Sprintf("Advisor Status: %s, Company Status: %s",
			application.ApprovalAdvisorStatus, application.ApprovalCompanyStatus), nil
}

func (repo InternshipApplicationController) DeleteApplicationByID(id string) error {
	result := repo.Connector.Delete(&model.InternshipApplication{}, id)
	if result.Error != nil {
			return fmt.Errorf("failed to delete application with ID %s: %w", id, result.Error)
	}
	if result.RowsAffected == 0 {
			return fmt.Errorf("no application found with ID %s", id)
	}
	return nil
}