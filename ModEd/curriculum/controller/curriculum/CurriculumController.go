package controller

import (
	modelCommon "ModEd/common/model"
	model "ModEd/curriculum/model/Internship"
	modelWorkload "ModEd/curriculum/model/instructor-workload"
	"errors"

	"gorm.io/gorm"
)

type ICurriculumController interface {
	// Put methods here
	// eg. CreateCurriculum(curriculum *modelCurriculum.Curriculum) error
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
