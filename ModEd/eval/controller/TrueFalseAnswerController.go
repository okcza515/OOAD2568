// MEP-1007
package controller

import (
	"ModEd/core"
	"ModEd/eval/model"

	"gorm.io/gorm"
)

type TrueFalseAnswerController struct {
	db *gorm.DB
	*core.BaseController[*model.TrueFalseAnswer]
}

func NewTrueFalseAnswerController(db *gorm.DB) *TrueFalseAnswerController {
	return &TrueFalseAnswerController{
		db:             db,
		BaseController: core.NewBaseController[*model.TrueFalseAnswer](db),
	}
}

func (c *TrueFalseAnswerController) GetTrueFalseAnswerByQuestionID(questionID uint) (tfAnswer *model.TrueFalseAnswer, err error) {
	err = c.db.
		Where("question_id = ?", questionID).
		Find(&tfAnswer).Error

	if err != nil {
		return nil, err
	}

	return tfAnswer, err
}
