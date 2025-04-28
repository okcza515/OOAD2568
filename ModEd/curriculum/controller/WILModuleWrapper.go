// MEP-1010 Work Integrated Learning (WIL)
package controller

import (
	"gorm.io/gorm"
)

type WILModuleWrapper struct {
	WILProjectController            *WILProjectController
	WILProjectApplicationController *WILProjectApplicationController
	WILProjectCurriculumController  *WILProjectCurriculumFacadeController
	IndependentStudyController      *IndependentStudyController
}

func NewWILModuleWrapper(
	db *gorm.DB,
	courseController CourseControllerInterface,
	classController ClassControllerInterface,
) *WILModuleWrapper {
	return &WILModuleWrapper{
		WILProjectController:            CreateWILProjectController(db),
		WILProjectApplicationController: CreateWILProjectApplicationController(db),
		WILProjectCurriculumController:  CreateWILProjectCurriculumFacadeController(db, courseController, classController),
		IndependentStudyController:      CreateIndependentStudyController(db),
	}
}
