package controller

import (
	model "ModEd/eval/model"
	// "errors"
	// "github.com/cockroachdb/errors"
	"gorm.io/gorm"
)

type IExaminationController interface {
	CreateExam(exam *model.Examination) error
	GetAll() ([]model.Examination, error)
	Update(id uint, exam *model.Examination) error
}

type ExaminationController struct {
	db *gorm.DB
}

func NewExaminationController(db *gorm.DB) *ExaminationController {
	return &ExaminationController{db: db}
}

func (c *ExaminationController) CreateExam(exam *model.Examination) error{
	if err := c.db.Create(exam).Error; err != nil {
		return err
	}
	return nil
}

func (c *ExaminationController) GetAll() ([]model.Examination, error) {
	var exam []model.Examination
	if err := c.db.Find(&exam).Error; err != nil {
		return nil, err
	}
	return exam, nil
}

func (c *ExaminationController) Update(id uint , exam *model.Examination)  error {

if err := c.db.Model(&exam).Where("id = ?", id).Updates(exam).Error; err != nil {
	return err
}
	return nil ;
}