// MEP-1007
package controller

import (
	"ModEd/core"
	"ModEd/eval/model"

	"gorm.io/gorm"
)

type ExamSectionController struct {
	db *gorm.DB
	*core.BaseController[*model.ExamSection]
}

func NewExamSectionController(db *gorm.DB) *ExamSectionController {
	return &ExamSectionController{
		db:             db,
		BaseController: core.NewBaseController[*model.ExamSection](db),
	}
}