// MEP-1007
package controller

import (

	"ModEd/core"
	"ModEd/eval/model"
	"ModEd/utils/deserializer"
	"fmt"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)
type ExaminationController struct {
	db *gorm.DB
	core *core.BaseController[*model.Examination]
}
type IExaminationController interface {
	CreateSeedExamination(path string) (exams []*model.Examination, err error)
	CreateExam(exam *model.Examination) (examId uint, err error)
	UpdateExamStatus(examId uint) (exam *model.Examination, err error)
	GetAllExams(preloads ...string) (exams []*model.Examination, err error)
	UpdateExam(updatedExam *model.Examination) (exam *model.Examination, err error)
	DeleteExam(examId uint) (exam *model.Examination, err error)
}


func NewExaminationController(db *gorm.DB) *ExaminationController {
	return &ExaminationController{
		db: db,
		core: core.NewBaseController[*model.Examination](db),
	}
}

func (c *ExaminationController) CreateExam(exam *model.Examination) (examId uint, err error) {
	if err := c.core.Insert(exam); err != nil {
		return 0, err
	}
	return examId, nil
}

// Read all exams
func (c *ExaminationController) GetAllExams(preloads ...string) (exams []*model.Examination, err error) {
	exams,err = c.core.List(nil,preloads...)
	if err != nil {
		return nil, err
	}
	return exams, nil
}

// Read one exam
func (c *ExaminationController) GetExam(examId uint, preload ...string) (exam *model.Examination, err error) {
	exam, err = c.core.RetrieveByCondition(map[string]interface{}{"id": examId}, preload...)
	if err != nil {
		return nil, err
	}
	return exam, nil
}

func (c *ExaminationController) UpdateExam(updatedExam *model.Examination) (exam *model.Examination, err error) {
	exam, err = c.core.RetrieveByCondition(map[string]interface{}{"id": updatedExam.ID})
	if err != nil {
		return nil, err
	}
	if exam.ExamStatus == "Publish" {
		return exam, errors.Wrap(err,"publish exam can not update")
	}
	if exam.ExamStatus == "Hidden" {
		return exam, errors.Wrap(err,"hidden exam can not update")
	}
	exam.ExamName = updatedExam.ExamName
	exam.InstructorID = updatedExam.InstructorID
	exam.CourseID = updatedExam.CourseID
	exam.ExamStatus = updatedExam.ExamStatus
	exam.Description = updatedExam.Description
	exam.Attempt = updatedExam.Attempt
	exam.StartDate = updatedExam.StartDate
	exam.EndDate = updatedExam.EndDate
	if err := c.core.UpdateByCondition(map[string]interface{}{"id": updatedExam.ID}, exam); err != nil{
		return nil, err
	}
	return exam, nil
}

// Not finish delete exam
// func (c *ExaminationController) DeleteExam(examId uint) (exam *model.Examination, err error) {
// 	c.db.Where("ExamID = ?",examId).Delete(&model.ExamSection{})
// 	c.db.Where("ExamID = ?",examId).Delete(&model.Submission{})
// 	exam,err = c.core.RetrieveByCondition(map[string]interface{}{"id":examId})
// 	if err != nil {
// 		return nil, err
// 	}

// 	if err := c.core.DeleteByCondition(map[string]interface{}{"id": examId}); err != nil {
// 		return nil, err
// 	}
// 	return exam, nil
// }

func (c *ExaminationController) CreateSeedExamination(path string) (exams []*model.Examination, err error) {
	deserializer, err := deserializer.NewFileDeserializer(path)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create file deserializer")
	}

	if err := deserializer.Deserialize(&exams); err != nil {
		return nil, errors.Wrap(err, "failed to deserialize exam")
	}

	for _, exam := range exams {
		_, err := c.CreateExam(exam)
		if err != nil {
			return nil, errors.Wrap(err, "failed to create exam")
		}
	}
	fmt.Println("Create Exam Seed Successfully")
	return exams, nil
}
