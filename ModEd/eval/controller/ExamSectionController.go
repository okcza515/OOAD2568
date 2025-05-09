package controller

// MEP-1007

import (
	"ModEd/core"
	"ModEd/eval/model"
	"fmt"

	"ModEd/utils/deserializer"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type ExamSectionController struct {
	db *gorm.DB
	core *core.BaseController[*model.ExamSection]
}

type IExamSection interface {
	CreateSeedExamSection(path string) (examSecs []*model.ExamSection, err error)
	CreateExamSection(examSec *model.ExamSection) (examSecId uint, err error)
	GetAllExamSections(preloads ...string) (examSecs []*model.ExamSection, err error)
	GetExamSection(examSecId uint, preload ...string) (examSec *model.ExamSection, err error)
	UpdateExamSection(updatedExamSec *model.ExamSection) (examSec *model.ExamSection, err error)
	DeleteExamSection(examSecId uint) (examSec *model.ExamSection, err error)
}

func NewExamSectionController(db *gorm.DB) *ExamSectionController {
	return &ExamSectionController{
		db: db,
		core: core.NewBaseController[*model.ExamSection](db),
	}
}

func (c *ExamSectionController) CreateExamSection(examSec *model.ExamSection) (examSecId uint, err error) {
	if err := c.core.Insert(examSec); err != nil {
		return 0, err
	}
	return examSecId, nil
}

func (c *ExamSectionController) GetAllExamSections(preloads ...string) (examSecs []*model.ExamSection, err error) {
	examSecs,err = c.core.List(nil,preloads...)
	if err != nil {
		return nil, err
	}
	return examSecs, nil
}

func (c *ExamSectionController) GetExamSection(examSecId uint, preload ...string) (examSec *model.ExamSection, err error) {
	examSec, err = c.core.RetrieveByCondition(map[string]interface{}{"id": examSecId}, preload...)
	if err != nil {
		return nil, err
	}
	return examSec, nil
}

func (c *ExamSectionController) UpdateExamSection(updatedExamSec *model.ExamSection) (examSec *model.ExamSection, err error){
	examSec, err = c.core.RetrieveByCondition(map[string]interface{}{"id": updatedExamSec.ID})
	if err != nil {
		return nil, err
	}
	examSec.ExamID = updatedExamSec.ExamID
	examSec.SectionNo = updatedExamSec.SectionNo
	examSec.Description = updatedExamSec.Description
	examSec.NumQuestions = updatedExamSec.NumQuestions
	examSec.Score = updatedExamSec.Score
	if err := c.core.UpdateByCondition(map[string]interface{}{"id": updatedExamSec.ID}, examSec); err != nil{
		return nil, err
	}
	return examSec, nil
} 

// not finish delete exam section
// func (c *ExamSectionController) DeleteExamSection(examSecId uint) (examSec *model.ExamSection, err error) {
// 	examSec, err = c.core.RetrieveByCondition(map[string]interface{}{"id": examSecId})
// 	if err != nil {
// 		return nil, err
// 	}

// 	if err := c.core.DeleteByCondition(map[string]interface{}{"id": examSecId}); err != nil {
// 		return nil, err
// 	}
// 	return examSec, nil
// }

func (c *ExamSectionController) CreateSeedExamSection(path string) (examSecs []*model.ExamSection, err error) {
	deserializer, err := deserializer.NewFileDeserializer(path)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create file deserializer")
	}

	if err := deserializer.Deserialize(&examSecs); err != nil {
		return nil, errors.Wrap(err, "failed to deserialize exam section")
	}

	for _, examSec := range examSecs {
		_, err := c.CreateExamSection(examSec)
		if err != nil {
			return nil, errors.Wrap(err, "failed to create exam section")
		}
	}
	fmt.Println("Create Exam Section Seed Successfully")
	return examSecs, nil
}