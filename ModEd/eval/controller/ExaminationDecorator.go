package controller

import (
	"ModEd/eval/model"
	"fmt"
)

type LoggingDecorator struct {
	Wrapped IExaminationController
}

func NewLoggingDecorator(wrapped IExaminationController) *LoggingDecorator {
	return &LoggingDecorator{Wrapped: wrapped}
}

func (d *LoggingDecorator) CreateExam(exam *model.Examination) error {
	fmt.Println("[LOG] Creating exam:", exam.Exam_name)
	return d.Wrapped.CreateExam(exam)
}

func (d *LoggingDecorator) GetAll() ([]model.Examination, error) {
	fmt.Println("[LOG] Retrieving all exams")
	exams, err := d.Wrapped.GetAll()
	if err != nil {
		return nil, err
	}
	// fmt.Println("[LOG] Retrieved exams:", exams)
	return exams, nil
}

func (d *LoggingDecorator) Update(id uint, exam *model.Examination) error {
	fmt.Println("[LOG] Updating exam with ID:", id)
	return d.Wrapped.Update(id, exam)
}


