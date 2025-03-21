package controller

import "gorm.io/gorm"

type ICurriculumController interface {
	// Put methods here
	// eg. CreateCurriculum(curriculum *modelCurriculum.Curriculum) error
}

type CurriculumController struct {
	db *gorm.DB
}

func NewCurriculumController(db *gorm.DB) ICurriculumController {
	return &CurriculumController{db: db}
}
