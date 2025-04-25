package controller

import (
	"ModEd/eval/model"
)

type IExaminationFacade interface {
	CreateExamination(exam *model.Examination, question *model.Question) error
	GetAllExams() ([]model.Examination, error)
	UpdateExamination(id uint, exam *model.Examination) error
}

type ExaminationFacade struct {
	Ctrl IExaminationAndQuestionController
}

func NewExaminationFacade(ctrl IExaminationAndQuestionController) *ExaminationFacade {
	return &ExaminationFacade{Ctrl: ctrl}
}

func (f *ExaminationFacade) CreateExamination(exam *model.Examination, question *model.Question) error {
    if exam != nil {
        if err := f.Ctrl.CreateExam(exam); err != nil {
            return err
        }
    }

    if question != nil {
        if err := f.Ctrl.CreateQuestion(question); err != nil {
            return err
        }
    }

    return nil
}

func (f *ExaminationFacade) GetAllExams() ([]model.Examination, error) {
	return f.Ctrl.GetAll()
}

func (f *ExaminationFacade) UpdateExamination(id uint, exam *model.Examination) error {
	return f.Ctrl.Update(id, exam)
}
