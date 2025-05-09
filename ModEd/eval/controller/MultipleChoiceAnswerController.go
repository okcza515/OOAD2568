package controller

import (
	"ModEd/core"
	"ModEd/eval/model"

	"gorm.io/gorm"
)

type MultipleChoiceAnswerController struct {
	db *gorm.DB
	core *core.BaseController[*model.MultipleChoiceAnswer]
}

type IMultipleChoiceAnswer interface {
	CreateMultipleChoiceAnswer(mcAnswer *model.MultipleChoiceAnswer) (mcAnswerId uint, err error)
	GetAllMultipleChoiceAnswers(preloads ...string) (mcAnswers []*model.MultipleChoiceAnswer, err error)
	GetMultipleChoiceAnswer(mcAnswerId uint, preload ...string) (mcAnswer *model.MultipleChoiceAnswer, err error)
	UpdateMultipleChoiceAnswer(updatedMcAnswer *model.MultipleChoiceAnswer) (mcAnswer *model.MultipleChoiceAnswer, err error)
	DeleteMultipleChoiceAnswer(mcAnswerId uint) (mcAnswer *model.MultipleChoiceAnswer, err error)
}

func NewMultipleChoiceAnswerController(db *gorm.DB) *MultipleChoiceAnswerController {
	return &MultipleChoiceAnswerController{
		db: db,
		core: core.NewBaseController[*model.MultipleChoiceAnswer](db),
	}
}

func (c *MultipleChoiceAnswerController) CreateMultipleChoiceAnswer(mcAnswer *model.MultipleChoiceAnswer) (mcAnswerId uint, err error) {
	if err := c.core.Insert(mcAnswer); err != nil {
		return 0, err
	}
	return mcAnswerId, nil
}

func (c *MultipleChoiceAnswerController) GetAllMultipleChoiceAnswers(preloads ...string) (mcAnswers []*model.MultipleChoiceAnswer, err error) {
	mcAnswers,err = c.core.List(nil,preloads...)
	if err != nil {
		return nil, err
	}
	return mcAnswers, nil
}

func (c *MultipleChoiceAnswerController) GetMultipleChoiceAnswer(mcAnswerId uint, preload ...string) (mcAnswer *model.MultipleChoiceAnswer, err error) {
	mcAnswer, err = c.core.RetrieveByCondition(map[string]interface{}{"id": mcAnswerId}, preload...)
	if err != nil {
		return nil, err
	}
	return mcAnswer, nil
}

func (c *MultipleChoiceAnswerController) UpdateMultipleChoiceAnswer(updatedMcAnswer *model.MultipleChoiceAnswer) (mcAnswer *model.MultipleChoiceAnswer, err error){
	mcAnswer, err = c.core.RetrieveByCondition(map[string]interface{}{"id": updatedMcAnswer.ID})
	if err != nil {
		return nil, err
	}
	mcAnswer.QuestionID = updatedMcAnswer.QuestionID
	mcAnswer.AnswerLabel = updatedMcAnswer.AnswerLabel
	mcAnswer.IsExpected = updatedMcAnswer.IsExpected
	if err := c.core.UpdateByCondition(map[string]interface{}{"id": updatedMcAnswer.ID}, mcAnswer); err != nil{
		return nil, err
	}
	return mcAnswer, nil
} 

func (c *MultipleChoiceAnswerController) DeleteMultipleChoiceAnswer(mcAnswerId uint) (mcAnswer *model.MultipleChoiceAnswer, err error) {
	mcAnswer, err = c.core.RetrieveByCondition(map[string]interface{}{"id": mcAnswerId})
	if err != nil {
		return nil, err
	}

	if err := c.core.DeleteByCondition(map[string]interface{}{"id": mcAnswerId}); err != nil {
		return nil, err
	}
	return mcAnswer, nil
}