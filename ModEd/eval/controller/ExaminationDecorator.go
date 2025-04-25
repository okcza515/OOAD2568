package controller

import (
	"ModEd/eval/model"
	"fmt"
)

type LoggingDecorator struct {
	Wrapped IExaminationFacade
}

func NewLoggingDecorator(wrapped IExaminationFacade) *LoggingDecorator {
	return &LoggingDecorator{Wrapped: wrapped}
}

// func (d *LoggingDecorator) CreateExam(exam *model.Examination) error {
// 	fmt.Println("[LOG] Creating exam:", exam.Exam_name)
// 	return d.Wrapped.Ctrl.CreateExam(exam)
// }

func (d *LoggingDecorator) CreateExamAndQuestion(exam *model.Examination, question *model.Question) error {
	fmt.Println("[LOG] Starting CreateExam process...")

	if exam.Exam_name == "" {
		fmt.Println("[LOG] Warning: Exam name is empty!")
	}

	if err := d.Wrapped.CreateExamination(exam, nil); err != nil {
		fmt.Println("[LOG] Error during CreateExamination (Exam):", err)
		return err
	}
	fmt.Println("[LOG] Successfully created the exam.")

	fmt.Println("[LOG] Now processing the question...")

	if question != nil {
		fmt.Println("[LOG] Warning: Question detail is empty!")
	}

	if err := d.Wrapped.CreateExamination(nil, question); err != nil {
		fmt.Println("[LOG] Error during CreateExamination (Question):", err)
		return err
	}
	fmt.Println("[LOG] Successfully created the question.")

	return nil
}

// func (d *LoggingDecorator) GetAll() ([]model.Examination, error) {
// 	fmt.Println("[LOG] Retrieving all exams")
// 	exams, err := d.Wrapped()
// 	if err != nil {
// 		return nil, err
// 	}
// 	fmt.Println("[LOG] Retrieved exams:", exams)
// 	return exams, nil
// }

// func (d *LoggingDecorator) Update(id uint, exam *model.Examination) error {
// 	fmt.Println("[LOG] Updating exam with ID:", id)
// 	return d.Wrapped.Update(id, exam)
// }
