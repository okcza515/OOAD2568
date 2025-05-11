// MEP-1007
package controller

import (
	"ModEd/core"
	"ModEd/eval/model"

	"gorm.io/gorm"
)

type ExamController struct {
	db *gorm.DB
	*core.BaseController[*model.Exam]
}

func NewExamController(db *gorm.DB) *ExamController {
	return &ExamController{
		db:             db,
		BaseController: core.NewBaseController[*model.Exam](db),
	}
}