package controller

// MEP-1007

import "gorm.io/gorm"

type ExamModuleWrapper struct {
	ExamController                           *ExamController
	QuestionController                       *QuestionController
	ExamSectionController                    *ExamSectionController
	SubmissionController					 *SubmissionController
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
		SubmissionController: 					  NewSubmissionController(db),	
		MultipleChoiceAnswerController:           NewMultipleChoiceAnswerController(db),
		MultipleChoiceAnswerSubmissionController: NewMultipleChoiceAnswerSubmissionController(db),
		ShortAnswerController:               	  NewShortAnswerController(db),
		ShortAnswerSubmissionController:          NewShortAnswerSubmissionController(db),
		TrueFalseAnswerController:                NewTrueFalseAnswerController(db),
		TrueFalseAnswerSubmissionController:      NewTrueFalseAnswerSubmissionController(db),
	}
}