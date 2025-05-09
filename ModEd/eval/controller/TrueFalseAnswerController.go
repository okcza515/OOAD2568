package controller

import (
	"ModEd/core"
	"ModEd/eval/model"

	"gorm.io/gorm"
)

type TrueFalseAnswerController struct {
	db   *gorm.DB
	core *core.BaseController[*model.TrueFalseAnswer]
}

type ITrueFalseAnswer interface {
	CreateTrueFalseAnswer(tfAnswer *model.TrueFalseAnswer) (tfAnswerId uint, err error)
	GetAllTrueFalseAnswers(preloads ...string) (tfAnswers []*model.TrueFalseAnswer, err error)
	GetTrueFalseAnswer(tfAnswerId uint, preload ...string) (tfAnswer *model.TrueFalseAnswer, err error)
	UpdateTrueFalseAnswer(updatedTfAnswer *model.TrueFalseAnswer) (tfAnswer *model.TrueFalseAnswer, err error)
	DeleteTrueFalseAnswer(tfAnswerId uint) (tfAnswer *model.TrueFalseAnswer, err error)
}

func NewTrueFalseAnswerController(db *gorm.DB) *TrueFalseAnswerController {
	return &TrueFalseAnswerController{
		db:   db,
		core: core.NewBaseController[*model.TrueFalseAnswer](db),
	}
}

func (c *TrueFalseAnswerController) CreateTrueFalseAnswer(tfAnswer *model.TrueFalseAnswer) (tfAnswerId uint, err error) {
	if err := c.core.Insert(tfAnswer); err != nil {
		return 0, err
	}
	return tfAnswer.ID, nil
}

func (c *TrueFalseAnswerController) GetAllTrueFalseAnswers(preloads ...string) (tfAnswers []*model.TrueFalseAnswer, err error) {
	tfAnswers, err = c.core.List(nil, preloads...)
	if err != nil {
		return nil, err
	}
	return tfAnswers, nil
}

func (c *TrueFalseAnswerController) GetTrueFalseAnswer(tfAnswerId uint, preload ...string) (tfAnswer *model.TrueFalseAnswer, err error) {
	tfAnswer, err = c.core.RetrieveByCondition(map[string]interface{}{"id": tfAnswerId}, preload...)
	if err != nil {
		return nil, err
	}
	return tfAnswer, nil
}

func (c *TrueFalseAnswerController) UpdateTrueFalseAnswer(updatedTfAnswer *model.TrueFalseAnswer) (tfAnswer *model.TrueFalseAnswer, err error) {
	tfAnswer, err = c.core.RetrieveByCondition(map[string]interface{}{"id": updatedTfAnswer.ID})
	if err != nil {
		return nil, err
	}
	tfAnswer.QuestionID = updatedTfAnswer.QuestionID
	tfAnswer.IsExpected = updatedTfAnswer.IsExpected
	if err := c.core.UpdateByCondition(map[string]interface{}{"id": updatedTfAnswer.ID}, tfAnswer); err != nil {
		return nil, err
	}
	return tfAnswer, nil
}

func (c *TrueFalseAnswerController) DeleteTrueFalseAnswer(tfAnswerId uint) (tfAnswer *model.TrueFalseAnswer, err error) {
	tfAnswer, err = c.core.RetrieveByCondition(map[string]interface{}{"id": tfAnswerId})
	if err != nil {
		return nil, err
	}

	if err := c.core.DeleteByCondition(map[string]interface{}{"id": tfAnswerId}); err != nil {
		return nil, err
	}
	return tfAnswer, nil
}
