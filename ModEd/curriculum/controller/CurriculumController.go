// MEP-1002
package controller

import (
	"fmt"

	"github.com/pkg/errors"
	"gorm.io/gorm"

	"ModEd/core"
	"ModEd/curriculum/model"
	"ModEd/utils/deserializer"
)

type CurriculumController struct {
	db   *gorm.DB
	core *core.BaseController[*model.Curriculum]
}

type CurriculumControllerInterface interface {
	CreateSeedCurriculum(path string) (curriculums []*model.Curriculum, err error)
	CreateCurriculum(curriculum *model.Curriculum) (curriculumId uint, err error)
	GetCurriculum(curriculumId uint, preload ...string) (curriculum *model.Curriculum, err error)
	GetCurriculums(preload ...string) (curriculums []*model.Curriculum, err error)
	UpdateCurriculum(updated *model.Curriculum) (curriculum *model.Curriculum, err error)
	DeleteCurriculum(curriculumId uint) (curriculum *model.Curriculum, err error)
}

func NewCurriculumController(db *gorm.DB) *CurriculumController {
	return &CurriculumController{
		db:   db,
		core: core.NewBaseController[*model.Curriculum](db),
	}
}

// Create
func (c *CurriculumController) CreateCurriculum(curriculum *model.Curriculum) (curriculumId uint, err error) {
	if err := c.core.Insert(curriculum); err != nil {
		return 0, err
	}
	return curriculum.CurriculumId, nil
}

// Read one
func (c *CurriculumController) GetCurriculum(curriculumId uint, preload ...string) (curriculum *model.Curriculum, err error) {
	curriculum, err = c.core.RetrieveByCondition(map[string]interface{}{"curriculum_id": curriculumId}, preload...)
	if err != nil {
		return nil, err
	}
	return curriculum, nil
}

// Read all
func (c *CurriculumController) GetCurriculums(preload ...string) (curriculums []*model.Curriculum, err error) {
	curriculums, err = c.core.List(nil, preload...)
	if err != nil {
		return nil, err
	}
	return curriculums, nil
}

// Update
func (c *CurriculumController) UpdateCurriculum(updatedCurriculum *model.Curriculum) (curriculum *model.Curriculum, err error) {
	curriculum, err = c.core.RetrieveByCondition(map[string]interface{}{"curriculum_id": updatedCurriculum.CurriculumId})
	if err != nil {
		return nil, err
	}

	// update fields
	curriculum.CurriculumId = updatedCurriculum.CurriculumId
	curriculum.Name = updatedCurriculum.Name
	curriculum.StartYear = updatedCurriculum.StartYear
	curriculum.EndYear = updatedCurriculum.EndYear
	curriculum.DepartmentId = updatedCurriculum.DepartmentId
	curriculum.ProgramType = updatedCurriculum.ProgramType

	if err := c.core.UpdateByCondition(map[string]interface{}{"curriculum_id": updatedCurriculum.CurriculumId}, curriculum); err != nil {
		return nil, err
	}
	return curriculum, nil
}

// Delete
func (c *CurriculumController) DeleteCurriculum(curriculumId uint) (curriculum *model.Curriculum, err error) {
	curriculum, err = c.core.RetrieveByCondition(map[string]interface{}{"curriculum_id": curriculumId})
	if err != nil {
		return nil, err
	}

	if err := c.core.DeleteByCondition(map[string]interface{}{"curriculum_id": curriculumId}); err != nil {
		return nil, err
	}
	return curriculum, nil
}

func (c *CurriculumController) CreateSeedCurriculum(path string) (curriculums []*model.Curriculum, err error) {
	deserializer, err := deserializer.NewFileDeserializer(path)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create file deserializer")
	}

	if err := deserializer.Deserialize(&curriculums); err != nil {
		return nil, errors.Wrap(err, "failed to deserialize curriculums")
	}

	for _, curriculum := range curriculums {
		_, err := c.CreateCurriculum(curriculum)
		if err != nil {
			return nil, errors.Wrap(err, "failed to create curriculum")
		}
	}
	fmt.Println("Create Curriculum Seed Successfully")
	return curriculums, nil
}
