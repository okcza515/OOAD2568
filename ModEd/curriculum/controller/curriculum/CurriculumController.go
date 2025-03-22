package controller

import (
<<<<<<< Updated upstream
	modelCommon "ModEd/common/model"
	model "ModEd/curriculum/model/Internship"
	modelWorkload "ModEd/curriculum/model/instructor-workload"
	"errors"

	"gorm.io/gorm"
=======
	"gorm.io/gorm"

	modelCurriculum "ModEd/curriculum/model"
>>>>>>> Stashed changes
)

type ICurriculumController interface {
	CreateCurriculum(curriculum modelCurriculum.Curriculum) (curriculumId uint, err error)
	GetCurriculum(curriculumId uint) (curriculum *modelCurriculum.Curriculum, err error)
	GetCurriculums() (curriculums []*modelCurriculum.Curriculum, err error)
	UpdateCurriculum(updatedCurriculum modelCurriculum.Curriculum) (curriculum *modelCurriculum.Curriculum, err error)
	DeleteCurriculum(curriculumId uint) (curriculum *modelCurriculum.Curriculum, err error)
}

type CurriculumController struct {
	Db *gorm.DB
}

func NewCurriculumController(db *gorm.DB) ICurriculumController {
	return &CurriculumController{Db: db}
}

func (c *CurriculumController) MigrateToDB() error {
	err := c.Db.AutoMigrate(
		&model.InternStudent{},
		&model.Company{},
		&model.InternshipSchedule{},
		&model.SupervisorReview{},
		&modelCommon.Student{},
		&model.InternshipReport{},
		&model.InternshipApplication{},
		&modelWorkload.StudentAdvisor{},
	)
	if err != nil {
		return errors.New("err: migration failed")
	}

	return nil
}

// Create
func (c *CurriculumController) CreateCurriculum(curriculum modelCurriculum.Curriculum) (curriculumId uint, err error) {
	if err := c.db.Create(&curriculum).Error; err != nil {
		return 0, err
	}
	return curriculum.ID, nil
}

// Read one
func (c *CurriculumController) GetCurriculum(curriculumId uint) (curriculum *modelCurriculum.Curriculum, err error) {
	curriculum = &modelCurriculum.Curriculum{}
	if err := c.db.First(curriculum, curriculumId).Error; err != nil {
		return nil, err
	}
	return curriculum, nil
}

// Read all
func (c *CurriculumController) GetCurriculums() (curriculums []*modelCurriculum.Curriculum, err error) {
	if err := c.db.Find(&curriculums).Error; err != nil {
		return nil, err
	}
	return curriculums, nil
}

// Update
func (c *CurriculumController) UpdateCurriculum(updated modelCurriculum.Curriculum) (curriculum *modelCurriculum.Curriculum, err error) {
	curriculum = &modelCurriculum.Curriculum{}
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
func (c *CurriculumController) DeleteCurriculum(curriculumId uint) (curriculum *modelCurriculum.Curriculum, err error) {
	curriculum = &modelCurriculum.Curriculum{}
	if err := c.db.First(curriculum, curriculumId).Error; err != nil {
		return nil, err
	}
	if err := c.db.Delete(curriculum).Error; err != nil {
		return nil, err
	}
	return curriculum, nil
}
