package controller

import (
	"ModEd/eval/model"

	"gorm.io/gorm"
)

type IExaminationFacade interface {
	CreateExamination(exam *model.Examination, question *model.Question) error
	GetAllExams() ([]model.Examination, error)
	UpdateExamination(id uint, exam *model.Examination) error
	UpdateQuestionForExam(id uint, updatedQuestion *model.Question) error
	GetQuestionsForExam(examID uint) ([]model.Question, error) 
	DeleteExamination(id uint) error
	DeleteQuestion(id uint) error
}

type ExaminationFacade struct {
	examCtrl     IExaminationController
	questionCtrl IQuestionController
}

func NewExaminationFacade(db *gorm.DB) *ExaminationFacade {
	examController := NewExaminationController(db)
	questionController := NewQuestionController(db)
	return &ExaminationFacade{examCtrl: examController, questionCtrl: questionController}
}

func (f *ExaminationFacade) CreateExamination(exam *model.Examination, question *model.Question) error {
	if exam != nil {
		if err := f.examCtrl.CreateExam(exam); err != nil {
			return err
		}
	}

	if question != nil {
		if err := f.questionCtrl.CreateQuestion(question); err != nil {
			return err
		}
	}

	return nil
}

func (f *ExaminationFacade) GetAllExams() ([]model.Examination, error) {
	return f.examCtrl.GetAll()
}

func (f *ExaminationFacade) UpdateExamination(id uint, exam *model.Examination) error {
	return f.examCtrl.Update(id, exam)
}

func (f *ExaminationFacade) UpdateQuestionForExam(id uint, updatedQuestion *model.Question) error {
    return f.questionCtrl.UpdateQuestion(id, updatedQuestion)
}

func (f *ExaminationFacade) GetQuestionsForExam(examID uint) ([]model.Question, error) {
    return f.questionCtrl.GetQuestionsByExamID(examID)
}

func (f *ExaminationFacade) DeleteExamination(id uint) error {
    return f.examCtrl.Delete(id)
}

func (f *ExaminationFacade) DeleteQuestion(id uint) error {
    return f.questionCtrl.DeleteQuestion(id)
}
