// MEP-1003 Student Recruitment
package controller

import (
	"ModEd/core"
	"ModEd/recruit/model"
	"ModEd/recruit/util"

	"gorm.io/gorm"
)

type ApplicantController struct {
	Base *core.BaseController[*model.Applicant]
	DB   *gorm.DB
}

func NewApplicantController(db *gorm.DB) *ApplicantController {
	return &ApplicantController{
		Base: core.NewBaseController[*model.Applicant](db),
		DB:   db,
	}
}

func (c *ApplicantController) RegisterApplicant(applicant *model.Applicant) error {
	return c.Base.Insert(applicant)
}

func (c *ApplicantController) GetAllApplicants() ([]*model.Applicant, error) {
	return c.Base.List(nil)
}

func (c *ApplicantController) GetApplicantByID(id uint) (*model.Applicant, error) {
	return c.Base.RetrieveByID(id)
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
