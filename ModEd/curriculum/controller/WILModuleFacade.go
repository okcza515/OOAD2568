// MEP-1010 Work Integrated Learning (WIL)
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
	courseController CourseControllerInterface,
	classController ClassControllerInterface,
) *WILModuleFacade {
	return &WILModuleFacade{
		WILProjectController:            CreateWILProjectController(db),
		WILProjectApplicationController: CreateWILProjectApplicationController(db),
		WILProjectCurriculumController:  CreateWILProjectCurriculumController(db, courseController, classController),
		IndependentStudyController:      CreateIndependentStudyController(db),
	}
}
