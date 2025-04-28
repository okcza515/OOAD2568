// MEP-1008
package controller

import (
	"gorm.io/gorm"
)

type InstructorWorkloadModuleWrapper struct {
	CoursePlanController     *CoursePlanController
	ClassLectureController   *ClassLectureController
	ClassMaterialController  *ClassMaterialController
	SeniorProjectController  *ProjectController
	StudentRequestController *StudentWorkloadController
	MeetingController        *MeetingController
}

func NewInstructorWorkloadModuleWrapper(
	db *gorm.DB,
	courseController CourseControllerInterface,
	classController ClassControllerInterface,
	curriculumController CurriculumControllerInterface,
) *InstructorWorkloadModuleWrapper {
	return &InstructorWorkloadModuleWrapper{
		CoursePlanController:     CreateCoursePlanController(db).(*CoursePlanController),
		ClassLectureController:   CreateClassLectureController(db).(*ClassLectureController),
		ClassMaterialController:  CreateClassMaterialController(db).(*ClassMaterialController),
		SeniorProjectController:  CreateProjectController(db).(*ProjectController),
		StudentRequestController: CreateStudentWorkloadController(db).(*StudentWorkloadController),
		MeetingController:        CreateMeetingController(db).(*MeetingController),
	}
}
