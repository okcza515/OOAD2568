// MEP-1007
package controller

import (
	"ModEd/core"
	"ModEd/eval/model"

	// "ModEd/utils/deserializer"
	// "fmt"

	"gorm.io/gorm"
)

type ExaminationController struct {
	db *gorm.DB
	*core.BaseController[*model.Exam]
}

func NewExaminationController(db *gorm.DB) *ExaminationController {
	return &ExaminationController{
		db:             db,
		BaseController: core.NewBaseController[*model.Exam](db),
	}
}
