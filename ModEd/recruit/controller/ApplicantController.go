package controller

import (
	"ModEd/core"
	"ModEd/recruit/model"
	"ModEd/recruit/util"

	"gorm.io/gorm"
)

type ApplicantController struct {
	Base *core.BaseController
	DB   *gorm.DB
}

func NewApplicantController(db *gorm.DB) *ApplicantController {
	return &ApplicantController{
		Base: core.NewBaseController("Applicant", db),
		DB:   db,
	}
}

func (c *ApplicantController) RegisterApplicant(applicant *model.Applicant) error {
	return c.Base.Insert(applicant)
}

func (c *ApplicantController) GetAllApplicants() ([]*model.Applicant, error) {
	records, err := c.Base.List(nil)
	if err != nil {
		return nil, err
	}

	var applicants []*model.Applicant
	for _, record := range records {
		if applicant, ok := record.(*model.Applicant); ok {
			applicants = append(applicants, applicant)
		}
	}
	return applicants, nil
}

func (c *ApplicantController) GetApplicantByID(id uint) (*model.Applicant, error) {
	var applicant model.Applicant
	if err := c.DB.Where("applicant_id = ?", id).First(&applicant).Error; err != nil {
		return nil, err
	}
	return &applicant, nil
}

func (c *ApplicantController) UpdateApplicant(applicant *model.Applicant) error {
	return c.Base.UpdateByID(applicant)
}

func (c *ApplicantController) DeleteApplicant(id uint) error {
	return c.Base.DeleteByID(id)
}

func (c *ApplicantController) ReadApplicantsFromFile(filePath string) ([]model.Applicant, error) {
	applicants, err := util.ReadOnlyFromCSVOrJSON[model.Applicant](filePath)
	if err != nil {
		return nil, err
	}

	return applicants, nil
}
