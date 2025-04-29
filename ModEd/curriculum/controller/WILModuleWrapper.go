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
		WILProjectController:            NewWILProjectController(db),
		WILProjectApplicationController: NewWILProjectApplicationController(db),
		WILProjectCurriculumController:  NewWILProjectCurriculumFacadeController(db, courseController, classController),
		IndependentStudyController:      NewIndependentStudyController(db),
	}
}
