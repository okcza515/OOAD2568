package controller

import "gorm.io/gorm"

type IClassController interface {
	// Put methods here
	// eg. CreateClass(class *modelCurriculum.Class) error
}

type ClassController struct {
	db *gorm.DB
}

func NewClassController(db *gorm.DB) IClassController {
	return &ClassController{db: db}
}
