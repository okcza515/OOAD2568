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
	ClassLectureController   *ClassLectureController
	ClassMaterialController  *ClassMaterialController
	SeniorProjectController  *ProjectController
	StudentRequestController *StudentWorkloadController
	MeetingController        *MeetingController
	WorkloadReportController *WorkloadReportController
	WorkloadReportFacade     *WorkloadReportFacade
	WorkloadReportBuilder    *WorkloadReportBuilder
}

func NewInstructorWorkloadModuleWrapper(
	db *gorm.DB,
	courseController CourseControllerInterface,
	classController ClassControllerInterface,
	curriculumController CurriculumControllerInterface,
) *InstructorWorkloadModuleWrapper {
	return &InstructorWorkloadModuleWrapper{
		CoursePlanController:     CreateCoursePlanController(db),
		ClassLectureController:   CreateClassLectureController(db),
		ClassMaterialController:  CreateClassMaterialController(db),
		SeniorProjectController:  CreateProjectController(db),
		StudentRequestController: CreateStudentWorkloadController(db),
		MeetingController:        CreateMeetingController(db),
		WorkloadReportController: CreateWorkloadReportController(db),
	}
}
