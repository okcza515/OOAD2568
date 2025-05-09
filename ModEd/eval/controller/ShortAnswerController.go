package controller

import (
	"ModEd/core"
	"ModEd/eval/model"

	"gorm.io/gorm"
)

type ShortAnswerController struct {
	db *gorm.DB
	core *core.BaseController[*model.ShortAnswer]
}

type IShortAnswer interface {
	CreateShortAnswer(shortAnswer *model.ShortAnswer) (shortAnswerId uint, err error)
	GetAllShortAnswers(preloads ...string) (shortAnswers []*model.ShortAnswer, err error)
	GetShortAnswer(shortAnswerId uint, preload ...string) (shortAnswer *model.ShortAnswer, err error)
	UpdateShortAnswer(updatedShortAnswer *model.ShortAnswer) (shortAnswer *model.ShortAnswer, err error)
	DeleteShortAnswer(shortAnswerId uint) (shortAnswer *model.ShortAnswer, err error)
}

func NewShortAnswerController(db *gorm.DB) *ShortAnswerController {
	return &ShortAnswerController{
		db: db,
		core: core.NewBaseController[*model.ShortAnswer](db),
	}
}

func (c *ShortAnswerController) CreateShortAnswer(shortAnswer *model.ShortAnswer) (shortAnswerId uint, err error) {
	if err := c.core.Insert(shortAnswer); err != nil {
		return 0, err
	}
	return shortAnswerId, nil
}

func (c *ShortAnswerController) GetAllShortAnswers(preloads ...string) (shortAnswers []*model.ShortAnswer, err error) {
	shortAnswers,err = c.core.List(nil,preloads...)
	if err != nil {
		return nil, err
	}
	return shortAnswers, nil
}

func (c *ShortAnswerController) GetShortAnswer(shortAnswerId uint, preload ...string) (shortAnswer *model.ShortAnswer, err error) {
	shortAnswer, err = c.core.RetrieveByCondition(map[string]interface{}{"id": shortAnswerId}, preload...)
	if err != nil {
		return nil, err
	}
	return shortAnswer, nil
}

func (c *ShortAnswerController) UpdateShortAnswer(updatedShortAnswer *model.ShortAnswer) (shortAnswer *model.ShortAnswer, err error){
	shortAnswer, err = c.core.RetrieveByCondition(map[string]interface{}{"id": updatedShortAnswer.ID})
	if err != nil {
		return nil, err
	}
	shortAnswer.QuestionID = updatedShortAnswer.QuestionID
	shortAnswer.ExpectedAnswer = updatedShortAnswer.ExpectedAnswer
	if err := c.core.UpdateByCondition(map[string]interface{}{"id": updatedShortAnswer.ID}, shortAnswer); err != nil{
		return nil, err
	}
	return shortAnswer, nil
} 

func (c *ShortAnswerController) DeleteShortAnswer(shortAnswerId uint) (shortAnswer *model.ShortAnswer, err error) {
	shortAnswer, err = c.core.RetrieveByCondition(map[string]interface{}{"id": shortAnswerId})
	if err != nil {
		return nil, err
	}

	if err := c.core.DeleteByCondition(map[string]interface{}{"id": shortAnswerId}); err != nil {
		return nil, err
	}
	return shortAnswer, nil
}