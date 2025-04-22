// MEP-1010 Work Integrated Learning (WIL)
package controller

import (
	"ModEd/core"
	"ModEd/curriculum/model"

	"gorm.io/gorm"
)

type WILProjectCurriculumController struct {
	connector        *gorm.DB
	courseController CourseControllerInterface
	classController  ClassControllerInterface
}

type WILProjectCurriculumControllerInterface interface {
	RegisterWILProjects(projects []core.RecordInterface)
}

func CreateWILProjectCurriculumController(
	connector *gorm.DB,
	courseController CourseControllerInterface,
	classController ClassControllerInterface,
) *WILProjectCurriculumController {

	return &WILProjectCurriculumController{
		connector:        connector,
		courseController: courseController,
		classController:  classController,
	}
}

func (controller WILProjectCurriculumController) CreateNewWILCourse(course *model.Course, semester string) (uint, error) {
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

func (controller WILProjectCurriculumController) CreateNewWILClass(class *model.Class) (uint, error) {
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

func (controller WILProjectCurriculumController) RetrieveAllWILCourses() ([]model.Course, error) {
	var wilProjectCourses []model.WILProjectCourse

	if err := controller.connector.Find(&wilProjectCourses).Error; err != nil {
		return nil, err
	}

	var courses []model.Course

	for _, wilProjectCourse := range wilProjectCourses {
		course, err := controller.courseController.GetCourseByID(wilProjectCourse.CourseId)
		if err != nil {
			return nil, err
		}
		courses = append(courses, *course)
	}

	return courses, nil
}

func (controller WILProjectCurriculumController) RetrieveAllWILClasses() ([]model.Class, error) {
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
