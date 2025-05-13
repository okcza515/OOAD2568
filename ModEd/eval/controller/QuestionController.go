// MEP-1007
package controller

import (
	"ModEd/core"
	"ModEd/eval/model"
	"fmt"

	"gorm.io/gorm"
)

type QuestionController struct {
	db *gorm.DB
	*core.BaseController[*model.Question]
	MCAnswerController *MultipleChoiceAnswerController
	SAAnswerController *ShortAnswerController
	TFAnswerController *TrueFalseAnswerController
}

func NewQuestionController(db *gorm.DB) *QuestionController {
	return &QuestionController{
		db:                 db,
		BaseController:     core.NewBaseController[*model.Question](db),
		MCAnswerController: NewMultipleChoiceAnswerController(db),
		SAAnswerController: NewShortAnswerController(db),
		TFAnswerController: NewTrueFalseAnswerController(db),
	}
}

func (qc *QuestionController) getDeleteFunc(questionType model.QuestionType) (func(uint) error, error) {
	deleteFuncs := map[model.QuestionType]func(uint) error{
		model.MultipleChoiceQuestion: func(id uint) error {
			return qc.MCAnswerController.DeleteByCondition(map[string]interface{}{"question_id": id})
		},
		model.ShortAnswerQuestion: func(id uint) error {
			return qc.SAAnswerController.DeleteByCondition(map[string]interface{}{"question_id": id})
		},
		model.TrueFalseQuestion: func(id uint) error {
			return qc.TFAnswerController.DeleteByCondition(map[string]interface{}{"question_id": id})
		},
	}

	deleteFunc, ok := deleteFuncs[questionType]
	if !ok {
		return nil, fmt.Errorf("unsupported question type: %s", questionType)
	}
	return deleteFunc, nil
}

func (qc *QuestionController) DeleteAnswerByQuestionType(questionType model.QuestionType, questionID uint) error {
	deleteFunc, err := qc.getDeleteFunc(questionType)
	if err != nil {
		return err
	}
	if err := deleteFunc(questionID); err != nil {
		return fmt.Errorf("failed to delete answers of type %s: %w", questionType, err)
	}
	return nil
}

func (qc *QuestionController) DeleteByQuestionID(questionID uint) error {

	question, err := qc.RetrieveByID(questionID)
	if err != nil {
		return fmt.Errorf("failed to retrieve question ID %d: %w", questionID, err)
	}

	if err := qc.DeleteAnswerByQuestionType(question.QuestionType, questionID); err != nil {
		return fmt.Errorf("failed to delete answers for question ID %d: %w", questionID, err)
	}

	if err := qc.DeleteByID(questionID); err != nil {
		return fmt.Errorf("failed to delete question ID %d: %w", questionID, err)
	}

	return nil
}

func (qc *QuestionController) UpdateQuestionType(newQuestionType model.QuestionType, question *model.Question) error {
	deleteFunc, err := qc.getDeleteFunc(question.QuestionType)
	if err != nil {
		return err
	}
	if err := deleteFunc(question.ID); err != nil {
		return fmt.Errorf("failed to delete existing answers for question type %s: %w", question.QuestionType, err)
	}

	question.QuestionType = newQuestionType

	if err := qc.UpdateByID(question); err != nil {
		return fmt.Errorf("failed to update question: %w", err)
	}

	return nil
}
