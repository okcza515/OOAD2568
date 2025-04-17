package controller

import (
	model "ModEd/eval/model"
	// "errors"
	// "github.com/cockroachdb/errors"
	"gorm.io/gorm"
)

type IExaminationController interface {
	CreateExam(exam *model.Examination) error
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

func (c *ExaminationController) GetAllExam() error {
	var exam []model.Examination
	if err := c.db.Find(exam).Error; err != nil {
		return err
	}
	return nil
}