package controller

import (
	"ModEd/core"
	model "ModEd/curriculum/model"
	"ModEd/curriculum/utils"
	"fmt"
	"time"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type InternshipApplicationController struct {
	*core.BaseController[model.InternshipApplication]
	Connector *gorm.DB
}

type SubmissionStrategy interface {
	Execute(application *model.InternshipApplication) error
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
		if err := repo.Connector.Create(application).Error; err != nil {
			return err
		}
	}
	return nil
}

func (repo InternshipApplicationController) GetInternshipApplicationByID(id string) (*model.InternshipApplication, error) {
	application := &model.InternshipApplication{}
	result := repo.Connector.Preload("InternshipReport").
		Preload("SupervisorReview").
		First(application, id)
	return application, result.Error
}

func (repo InternshipApplicationController) GetApplicationStatusByID(id string) (string, error) {
	var application model.InternshipApplication
	if err := repo.Connector.Select("approval_advisor_status", "approval_company_status").
		Where("student_code = ?", id).
		First(&application).Error; err != nil {
		return "", fmt.Errorf("failed to find application with Student Code %s: %w", id, err)
	}

	return fmt.Sprintf("Advisor Status: %s, Company Status: %s",
		application.ApprovalUniversityStatus, application.ApprovalCompanyStatus), nil
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

func (repo InternshipApplicationController) UpdateInternshipApplication(application *model.InternshipApplication) error {
	result := repo.Connector.Save(application)
	if result.Error != nil {
		return fmt.Errorf("failed to update internship application: %w", result.Error)
	}
	return nil
}

type StudentToUniversityStrategy struct{Connector *gorm.DB}

func (s *StudentToUniversityStrategy) Execute(app *model.InternshipApplication) error {
	studentCode := utils.GetUserInput("Enter Student Code: ")
	if len(studentCode) != 11 {
		fmt.Println("Student Code cannot be empty.")
	}

	companyName := utils.GetUserInput("Enter Company Name: ")
	if companyName == "" {
		fmt.Println("Company Name cannot be empty.")
		return fmt.Errorf("invalid input: company name is empty")
	}

	companyController := NewCompanyController(s.Connector)
	company, err := companyController.GetCompanyByName(companyName)
	if err != nil {
		fmt.Printf("Error finding company with name '%s': %v\n", companyName, err)
		return fmt.Errorf("failed to find company: %w", err)
	}

	app.TurninDate = time.Now()
	app.ApprovalUniversityStatus = model.INTERN_APP_IN_PROGRESS
	app.ApprovalCompanyStatus = model.INTERN_APP_NOT_START
	app.CompanyId = company.ID
	app.StudentCode = studentCode

	if err := s.Connector.Create(app).Error; err != nil {
		fmt.Println("Error creating internship application:", err)
		return fmt.Errorf("failed to register internship application: %w", err)
	}
	fmt.Println("Internship application created successfully!")
	return nil
}

type UniversityToCompanyStrategy struct{Connector *gorm.DB}

func (s *UniversityToCompanyStrategy) Execute(app *model.InternshipApplication) error {
	controller := InternshipApplicationController{Connector: s.Connector}
	application, err := controller.GetInternshipApplicationByID(fmt.Sprint(app.ID))
	if err != nil {
		return fmt.Errorf("failed to retrieve internship application: %w", err)
	}
	app = application
	if app == nil {
		return fmt.Errorf("internship application not found")
	}
	app.ApprovalCompanyStatus = model.INTERN_APP_IN_PROGRESS
	app.UpdatedAt = time.Now()
	if err := s.Connector.Save(app).Error; err != nil {
		return fmt.Errorf("failed to update internship application: %w", err)
	}
	fmt.Println("University responded to company")
	return nil
}

type CompanyToUniversityStrategy struct{Connector *gorm.DB}

func (s *CompanyToUniversityStrategy) Execute(app *model.InternshipApplication) error {
	controller := InternshipApplicationController{Connector: s.Connector}
	application, err := controller.GetInternshipApplicationByID(fmt.Sprint(app.ID))
	if err != nil {
		return fmt.Errorf("failed to retrieve internship application: %w", err)
	}
	app = application
	if app == nil {
		return fmt.Errorf("internship application not found")
	}
	app.ApprovalUniversityStatus = model.INTERN_APP_APPROVED
	app.ApprovalCompanyStatus = model.INTERN_APP_APPROVED
	app.UpdatedAt = time.Now()
	if err := s.Connector.Save(app).Error; err != nil {
		return fmt.Errorf("failed to update internship application: %w", err)
	}
	fmt.Println("Company responded to university")
	return nil
}

type SubmissionExecutor struct {
	Strategy SubmissionStrategy
}

func (e *SubmissionExecutor) SetStrategy(strategy SubmissionStrategy) {
	e.Strategy = strategy
}

func (e *SubmissionExecutor) Execute(app *model.InternshipApplication) error {
	if e.Strategy == nil {
		return errors.New("no strategy set")
	}
	return e.Strategy.Execute(app)
}

func SubmitApplication(app *model.InternshipApplication, role string) error {
	executor := &SubmissionExecutor{}
	switch role {
	case "student":
		executor.SetStrategy(&StudentToUniversityStrategy{})
	case "university":
		executor.SetStrategy(&UniversityToCompanyStrategy{})
	case "company":
		executor.SetStrategy(&CompanyToUniversityStrategy{})
	default:
		return errors.New("invalid role")
	}

	return executor.Execute(app)
}