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
	core *core.BaseController
}

type CurriculumControllerInterface interface {
	CreateSeedCurriculum(path string) (curriculums []*model.Curriculum, err error)
	CreateCurriculum(curriculum *model.Curriculum) (curriculumId uint, err error)
	GetCurriculum(curriculumId uint) (curriculum *model.Curriculum, err error)
	GetCurriculums() (curriculums []*model.Curriculum, err error)
	UpdateCurriculum(updated *model.Curriculum) (curriculum *model.Curriculum, err error)
	DeleteCurriculum(curriculumId uint) (curriculum *model.Curriculum, err error)
}

func NewCurriculumController(db *gorm.DB) *CurriculumController {
	return &CurriculumController{
		db:   db,
		core: core.NewBaseController("Curriculum", db),
	}
}

// Create
func (c *CurriculumController) CreateCurriculum(curriculum *model.Curriculum) (curriculumId uint, err error) {
	if err := c.db.Create(&curriculum).Error; err != nil {
		return 0, err
	}
	return curriculum.CurriculumId, nil
}

// Read one
func (c *CurriculumController) GetCurriculum(curriculumId uint) (curriculum *model.Curriculum, err error) {
	curriculum = &model.Curriculum{}
	if err := c.db.First(curriculum, curriculumId).Error; err != nil {
		return nil, err
	}
	return curriculum, nil
}

// Read all
func (c *CurriculumController) GetCurriculums() (curriculums []*model.Curriculum, err error) {
	if err := c.db.Find(&curriculums).Error; err != nil {
		return nil, err
	}
	return curriculums, nil
}

// Update
func (c *CurriculumController) UpdateCurriculum(updated *model.Curriculum) (curriculum *model.Curriculum, err error) {
	curriculum = &model.Curriculum{}
	if err := c.db.First(curriculum, updated.CurriculumId).Error; err != nil {
		return nil, err
	}

	// update fields
	curriculum.CurriculumId = updated.CurriculumId
	curriculum.Name = updated.Name
	curriculum.StartYear = updated.StartYear
	curriculum.EndYear = updated.EndYear
	curriculum.DepartmentName = updated.DepartmentName
	curriculum.ProgramType = updated.ProgramType

	if err := c.db.Updates(curriculum).Error; err != nil {
		return nil, err
	}
	return curriculum, nil
}

// Delete
func (c *CurriculumController) DeleteCurriculum(curriculumId uint) (curriculum *model.Curriculum, err error) {
	curriculum = &model.Curriculum{}
	if err := c.db.First(curriculum, curriculumId).Error; err != nil {
		return nil, err
	}
	if err := c.db.Delete(curriculum).Error; err != nil {
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
