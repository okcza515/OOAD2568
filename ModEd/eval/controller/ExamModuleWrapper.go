package controller

import "gorm.io/gorm"

type ExamModuleWrapper struct {
	ExamController                           *ExamController
	QuestionController                       *QuestionController
	ExamSectionController                    *ExamSectionController
	MultipleChoiceAnswerController           *MultipleChoiceAnswerController
	MultipleChoiceAnswerSubmissionController *MultipleChoiceAnswerSubmissionController
	ShortAnswerController                    *ShortAnswerController
	ShortAnswerSubmissionController          *ShortAnswerSubmissionController
	TrueFalseAnswerController                *TrueFalseAnswerController
	TrueFalseAnswerSubmissionController      *TrueFalseAnswerSubmissionController
}

func NewExamModuleWrapper(db *gorm.DB) *ExamModuleWrapper {
	return &ExamModuleWrapper{
		ExamController:                           NewExamController(db),
		QuestionController:                       NewQuestionController(db),
		ExamSectionController:                    NewExamSectionController(db),
		MultipleChoiceAnswerController:           NewMultipleChoiceAnswerController(db),
		MultipleChoiceAnswerSubmissionController: NewMultipleChoiceAnswerSubmissionController(db),
		ShortAnswerController:               	  NewShortAnswerController(db),
		ShortAnswerSubmissionController:          NewShortAnswerSubmissionController(db),
		TrueFalseAnswerController:                NewTrueFalseAnswerController(db),
		TrueFalseAnswerSubmissionController:      NewTrueFalseAnswerSubmissionController(db),
	}
}