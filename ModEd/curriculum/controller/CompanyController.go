package controller

import (
	"ModEd/curriculum/model"
	"fmt"

	"gorm.io/gorm"
)

type CompanyController struct {
	Connector *gorm.DB
}

func NewCompanyController(connector *gorm.DB) *CompanyController {
	return &CompanyController{
		Connector: connector,
	}
}

func (cc *CompanyController) Create(company *model.Company) error {
	if err := cc.Connector.Create(company).Error; err != nil {
		return fmt.Errorf("failed to create company: %w", err)
	}
	return nil
}

func (cc *CompanyController) RetrieveByID(id uint) (*model.Company, error) {
	var company model.Company
	if err := cc.Connector.First(&company, id).Error; err != nil {
		return nil, fmt.Errorf("failed to retrieve company with ID %d: %w", id, err)
	}
	return &company, nil
}

func (cc *CompanyController) Update(company *model.Company) error {
	if err := cc.Connector.Save(company).Error; err != nil {
		return fmt.Errorf("failed to update company with ID %d: %w", company.ID, err)
	}
	return nil
}

func (cc *CompanyController) DeleteByID(id uint) error {
	if err := cc.Connector.Delete(&model.Company{}, id).Error; err != nil {
		return fmt.Errorf("failed to delete company with ID %d: %w", id, err)
	}
	return nil
}

func (cc *CompanyController) ListAll() ([]model.Company, error) {
	var companies []model.Company
	if err := cc.Connector.Find(&companies).Error; err != nil {
		return nil, fmt.Errorf("failed to list companies: %w", err)
	}
	return companies, nil
}

func (cc *CompanyController) GetCompanyByName(name string) (*model.Company, error) {
	var company model.Company
	if err := cc.Connector.Where("company_name = ?", name).First(&company).Error; err != nil {
		return nil, fmt.Errorf("company with name '%s' not found: %w", name, err)
	}
	return &company, nil
}
