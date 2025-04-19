// MEP-1003 Student Recruitment
package controller

import (
	"ModEd/recruit/controller/SQL"
	"ModEd/recruit/model"
	"ModEd/recruit/util"

	"gorm.io/gorm"
)

type ApplicantController struct {
	sqlCtrl SQL.SQLController[model.Applicant]
}

func NewApplicantController(db *gorm.DB) *ApplicantController {
	return &ApplicantController{
		SQL.NewGormSQLController[model.Applicant](db),
	}
}

func (c *ApplicantController) RegisterApplicant(applicant *model.Applicant) error {
	return c.sqlCtrl.Create(applicant)
}

func (c *ApplicantController) GetAllApplicants() ([]model.Applicant, error) {
	return c.sqlCtrl.GetAll()
}

func (c *ApplicantController) GetApplicantByID(id uint) (model.Applicant, error) {
	return c.sqlCtrl.GetByID(id)
}

func (c *ApplicantController) UpdateApplicant(applicant *model.Applicant) error {
	return c.sqlCtrl.Update(applicant)
}

func (c *ApplicantController) DeleteApplicant(id uint) error {
	return c.sqlCtrl.Delete(id)
}

func (c *ApplicantController) ReadApplicantsFromFile(filePath string) ([]model.Applicant, error) {
	applicants, err := util.ReadOnlyFromCSVOrJSON[model.Applicant](filePath)
	if err != nil {
		return nil, err
	}

	return applicants, nil
}
