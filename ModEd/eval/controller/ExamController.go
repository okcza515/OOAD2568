// MEP-1007
package controller

import (
	"ModEd/core"
	"ModEd/eval/model"

	"github.com/pkg/errors"
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

func (c *ExamController) UpdateExam(updatedExam *model.Exam) (exam *model.Exam, err error) {
	exam, err = c.BaseController.RetrieveByCondition(map[string]interface{}{"id": updatedExam.ID})
	if err != nil {
		return nil, err
	}
	if exam.ExamStatus == "Publish" {
		return exam, errors.Wrap(err, "publish exam can not update")
	}
	if exam.ExamStatus == "Hidden" {
		return exam, errors.Wrap(err, "hidden exam can not update")
	}
	exam.ExamName = updatedExam.ExamName
	exam.InstructorID = updatedExam.InstructorID
	exam.CourseID = updatedExam.CourseID
	exam.ExamStatus = updatedExam.ExamStatus
	exam.Description = updatedExam.Description
	exam.Attempt = updatedExam.Attempt
	exam.StartDate = updatedExam.StartDate
	exam.EndDate = updatedExam.EndDate
	if err := c.BaseController.UpdateByCondition(map[string]interface{}{"id": updatedExam.ID}, exam); err != nil {
		return nil, err
	}
	return exam, nil
}
