package controller

import (
	"ModEd/eval/model"
	"ModEd/eval/service"
)

type IExaminationController interface {
	CreateExam(exam *model.Examination) error
	PublishExam(id uint) error
	CloseExam(id uint) error
	GetAll() ([]model.Examination, error)
	Update(id uint, exam *model.Examination) error
	Delete(id uint) error
}

type ExaminationController struct {
	service *service.ExaminationService
}

func NewExaminationController(svc *service.ExaminationService) *ExaminationController {
	return &ExaminationController{
		service: svc,
	}
}

func (c *ExaminationController) CreateExam(exam *model.Examination) error {
	return c.service.CreateExam(exam)
}

func (c *ExaminationController) PublishExam(id uint) error {
	return c.service.PublishExam(id)
}

func (c *ExaminationController) CloseExam(id uint) error {
	return c.service.CloseExam(id)
}

func (c *ExaminationController) GetAll() ([]model.Examination, error) {
	return c.service.GetAllExams()
}

func (c *ExaminationController) Update(id uint, exam *model.Examination) error {
	return c.service.UpdateExam(id, exam)
}

func (c *ExaminationController) Delete(id uint) error {
	return c.service.DeleteExam(id)
}
