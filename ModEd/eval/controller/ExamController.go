// MEP-1007
package controller

import (
	"ModEd/core"
	"ModEd/eval/model"
	curriculumModel "ModEd/curriculum/model"

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

func (c *ExamController) GetPerfectScoreByExamID(examID uint) (score float64, err error) {
	var sections []*model.ExamSection
	err = c.db.
		Where("exam_id = ?", examID).
		Find(&sections).Error

	if err != nil {
		return 0, err
	}

	var perfectScore = 0.0
	for _, section := range sections {
		perfectScore += section.Score
	}

	return perfectScore, nil
}

func (c *ExamController) ListActiveExamsByStudentID(studentID uint) (exams []*model.Exam, err error) {
	var classes []curriculumModel.Class
	err = c.db.
		Joins("JOIN class_students ON class_students.class_class_id = classes.class_id").
		Where("class_students.student_id = ?", studentID).
		Find(&classes).Error

	if err != nil {
		return nil, err
	}

	for _, class := range classes {
		var exam *model.Exam
		err = c.db.
			Where("class_id = ?", class.ClassId).
			Find(&exam).Error

		if err != nil {
			return nil, err
		}

		if exam.ExamStatus == "Publish" {
			exams = append(exams, exam)
		}
	} 

	return exams, nil
}