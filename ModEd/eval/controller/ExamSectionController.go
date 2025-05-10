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

func (c *ExamSectionController) RetrieveByExamID(examID uint) ([]*model.ExamSection, error) {
	var examSections []*model.ExamSection
	if err := c.db.Where("exam_id = ?", examID).Find(&examSections).Error; err != nil {
		return nil, err
	}
	return examSections, nil
}