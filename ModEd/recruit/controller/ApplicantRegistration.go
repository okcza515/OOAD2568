package controller

import (
    "ModEd/recruit/model"
    "gorm.io/gorm"
)

type ApplicantController struct {
    DB *gorm.DB
}

func CreateApplicantController(db *gorm.DB) *ApplicantController {
    return &ApplicantController{DB: db}
}

func (c *ApplicantController) RegisterApplicant(applicant *model.Applicant) error {
    return c.DB.Create(applicant).Error
}

func (c *ApplicantController) GetAllApplicants() ([]model.Applicant, error) {
    var applicants []model.Applicant
    err := c.DB.Find(&applicants).Error
    return applicants, err
}
