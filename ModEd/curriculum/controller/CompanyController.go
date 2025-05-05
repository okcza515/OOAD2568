// MEP-1009 Student Internship
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

func (cc *CompanyController) GetCompanyByName(name string) (*model.Company, error) {
    var company model.Company
    if err := cc.Connector.Where("name = ?", name).First(&company).Error; err != nil {
        return nil, fmt.Errorf("company with name '%s' not found: %w", name, err)
    }
    return &company, nil
}