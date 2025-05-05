// MEP-1008
package controller

import (
	"gorm.io/gorm"
)

type InstructorWorkloadModuleWrapper struct {
	CourseController         *CourseController
	ClassController          *ClassController
	CurriculumController     *CurriculumController
	CoursePlanController     *CoursePlanController
	ClassMaterialController  *ClassMaterialController
	SeniorProjectController  *ProjectController
	StudentRequestController *StudentRequestController
	MeetingController        *MeetingController
}

func NewInstructorWorkloadModuleWrapper(
	db *gorm.DB,
	courseController CourseControllerInterface,
	classController ClassControllerInterface,
	curriculumController CurriculumControllerInterface,
) *InstructorWorkloadModuleWrapper {
	return &InstructorWorkloadModuleWrapper{
		CoursePlanController:    NewCoursePlanController(db),
		ClassMaterialController: NewClassMaterialController(db),
		SeniorProjectController: NewProjectController(db),
		MeetingController:       NewMeetingController(db),
	}
}
