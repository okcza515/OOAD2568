// MEP-1010 Work Integrated Learning (WIL)
package controller

import (
	"gorm.io/gorm"
)

type WILModuleProxy struct {
	WILProjectController            *WILProjectController
	WILProjectApplicationController *WILProjectApplicationController
	WILProjectCurriculumController  *WILProjectCurriculumFacadeController
	IndependentStudyController      *IndependentStudyController
}

func NewWILModuleProxy(
	db *gorm.DB,
	courseController CourseControllerInterface,
	classController ClassControllerInterface,
) *WILModuleProxy {
	return &WILModuleProxy{
		WILProjectController:            CreateWILProjectController(db),
		WILProjectApplicationController: CreateWILProjectApplicationController(db),
		WILProjectCurriculumController:  CreateWILProjectCurriculumFacadeController(db, courseController, classController),
		IndependentStudyController:      CreateIndependentStudyController(db),
	}
}
