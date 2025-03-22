package controller

import (
	"gorm.io/gorm"

	"ModEd/curriculum/model"
)

type ICurriculumController interface {
	CreateCurriculum(curriculum model.Curriculum) (curriculumId uint, err error)
	GetCurriculum(curriculumId uint) (curriculum *model.Curriculum, err error)
	GetCurriculums() (curriculums []*model.Curriculum, err error)
	UpdateCurriculum(updatedCurriculum model.Curriculum) (curriculum *model.Curriculum, err error)
	DeleteCurriculum(curriculumId uint) (curriculum *model.Curriculum, err error)
}

type CurriculumController struct {
	db *gorm.DB
}

func NewCurriculumController(db *gorm.DB) ICurriculumController {
	return &CurriculumController{db: db}
}

// Create
func (c *CurriculumController) CreateCurriculum(curriculum model.Curriculum) (curriculumId uint, err error) {
	if err := c.db.Create(&curriculum).Error; err != nil {
		return 0, err
	}
	return curriculum.ID, nil
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
func (c *CurriculumController) UpdateCurriculum(updated model.Curriculum) (curriculum *model.Curriculum, err error) {
	curriculum = &model.Curriculum{}
	if err := c.db.First(curriculum, updated.ID).Error; err != nil {
		return nil, err
	}

	// update fields
	curriculum.CurriculumId = updated.CurriculumId
	curriculum.Name = updated.Name
	curriculum.StartYear = updated.StartYear
	curriculum.EndYear = updated.EndYear
	curriculum.DepartmentName = updated.DepartmentName
	curriculum.ProgramType = updated.ProgramType

	if err := c.db.Save(curriculum).Error; err != nil {
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
