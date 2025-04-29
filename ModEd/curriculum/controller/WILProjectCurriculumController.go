// MEP-1010 Work Integrated Learning (WIL)
package controller

import (
	"ModEd/core"
	"ModEd/curriculum/model"

	"gorm.io/gorm"
)

type WILProjectCurriculumFacadeController struct {
	connector        *gorm.DB
	courseController CourseControllerInterface
	classController  ClassControllerInterface
}

type WILProjectCurriculumFacadeControllerInterface interface {
	RegisterWILProjects(projects []core.RecordInterface)
}

func NewWILProjectCurriculumFacadeController(
	connector *gorm.DB,
	courseController CourseControllerInterface,
	classController ClassControllerInterface,
) *WILProjectCurriculumFacadeController {

	return &WILProjectCurriculumFacadeController{
		connector:        connector,
		courseController: courseController,
		classController:  classController,
	}
}

func (controller WILProjectCurriculumFacadeController) CreateNewWILCourse(course *model.Course, semester string) (uint, error) {
	courseId, err := controller.courseController.CreateCourse(course)
	if err != nil {
		return 0, err
	}

	storeCourse := &model.WILProjectCourse{
		CourseId: courseId,
		Semester: semester,
	}

	if err := controller.connector.Create(storeCourse).Error; err != nil {
		return 0, err
	}

	return courseId, nil
}

func (controller WILProjectCurriculumFacadeController) CreateNewWILClass(class *model.Class) (uint, error) {
	classId, err := controller.classController.CreateClass(class)
	if err != nil {
		return 0, err
	}

	storeClass := &model.WILProjectClass{
		CourseId: class.CourseId,
		ClassId:  classId,
	}

	if err := controller.connector.Create(storeClass).Error; err != nil {
		return 0, err
	}

	return classId, nil
}

func (controller WILProjectCurriculumFacadeController) RetrieveAllWILCourses() ([]model.Course, error) {
	var wilProjectCourses []model.WILProjectCourse

	if err := controller.connector.Find(&wilProjectCourses).Error; err != nil {
		return nil, err
	}

	var courses []model.Course

	for _, wilProjectCourse := range wilProjectCourses {
		course, err := controller.courseController.GetCourse(wilProjectCourse.CourseId)
		if err != nil {
			return nil, err
		}
		courses = append(courses, *course)
	}

	return courses, nil
}

func (controller WILProjectCurriculumFacadeController) RetrieveAllWILClasses() ([]model.Class, error) {
	var wilProjectClasses []model.WILProjectClass

	if err := controller.connector.Find(&wilProjectClasses).Error; err != nil {
		return nil, err
	}

	var classes []model.Class

	for _, wilProjectCourse := range wilProjectClasses {
		class, err := controller.classController.GetClass(wilProjectCourse.ClassId)
		if err != nil {
			return nil, err
		}
		classes = append(classes, *class)
	}

	return classes, nil
}
