package controller

import "ModEd/eval/model"

type ExaminationFacade struct {
	Ctrl IExaminationController
}

func NewExaminationFacade(ctrl IExaminationController) *ExaminationFacade {
	return &ExaminationFacade{Ctrl: ctrl}
}

func (f *ExaminationFacade) CreateExamination(exam *model.Examination) error {
	return f.Ctrl.CreateExam(exam)
}

func (f *ExaminationFacade) GetAllExams() ([]model.Examination, error) {
	return f.Ctrl.GetAll()
}

func (f *ExaminationFacade) UpdateExamination(id uint, exam *model.Examination) error {
	return f.Ctrl.Update(id, exam)
}
