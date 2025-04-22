package controller

import (
	"gorm.io/gorm"
)

type WILModuleFacade struct {
	WILProjectController            *WILProjectController
	WILProjectApplicationController *WILProjectApplicationController
	WILProjectCurriculumController  *WILProjectCurriculumController
	IndependentStudyController      *IndependentStudyController
}

func NewWILModuleFacade(
	db *gorm.DB,
	courseController *CourseController,
	classController *ClassController,
) *WILModuleFacade {
	return &WILModuleFacade{
		WILProjectController:            CreateWILProjectController(db),
		WILProjectApplicationController: CreateWILProjectApplicationController(db),
		WILProjectCurriculumController:  CreateWILProjectCurriculumController(db, courseController, classController),
		IndependentStudyController:      CreateIndependentStudyController(db),
	}
}
