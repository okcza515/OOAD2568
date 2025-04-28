package controller

import (
	"ModEd/eval/model"

	"gorm.io/gorm"
)

type IExaminationFacade interface {
	CreateExamination(exam *model.Examination, question *model.Question) error
	GetAllExams() ([]model.Examination, error)
	UpdateExamination(id uint, exam *model.Examination) error
}

type ExaminationFacade struct {
	examCtrl     IExaminationController
	questionCtrl IQuestionController
    answerCtrl IAnswerController
}

func NewExaminationFacade(db *gorm.DB) *ExaminationFacade {
	examController := NewExaminationController(db)
	questionController := NewQuestionController(db)
	answerController := NewAnswerController(db)
	return &ExaminationFacade{examCtrl: examController, questionCtrl: questionController , answerCtrl: answerController}
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
