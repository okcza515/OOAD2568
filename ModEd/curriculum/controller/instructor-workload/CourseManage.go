package controller

import (
	curriculumModel "ModEd/curriculum/model"
	model "ModEd/curriculum/model/instructor-workload"
	"fmt"

	"gorm.io/gorm"
)

type CourseManageController struct {
	Connector *gorm.DB
}

func CreateCourseManageController(connector *gorm.DB) *CourseManageController {
	return &CourseManageController{
		Connector: connector,
	}
}

func (c CourseManageController) GetScheduleByCourse(courseId string) ([]*model.ClassSchedule, error) {
	schedules := []*model.ClassSchedule{}
	result := c.Connector.Find(&schedules, "course_id = ?", courseId)
	if result.Error != nil {
		return nil, result.Error
	}
	return schedules, nil
}

func (courseUpdate *CourseManageController) UpdateCourseName(courseNameUpdate *model.CourseNameUpdate) error {
	var course curriculumModel.Course
	result := courseUpdate.Connector.First(&course, "id = ?", courseNameUpdate.CourseId)
	if result.Error != nil {
		return fmt.Errorf("course not found: %w", result.Error)
	}

	course.Name = courseNameUpdate.NewName
	result = courseUpdate.Connector.Save(&course)
	if result.Error != nil {
		return fmt.Errorf("failed to update course name: %w", result.Error)
	}

	return nil
}

func (courseUpdate *CourseManageController) UpdateCoursePrerequisite(coursePrerequisiteUpdate *model.CoursePrerequisiteUpdate) error {
	var course curriculumModel.Course
	result := courseUpdate.Connector.First(&course, "id = ?", coursePrerequisiteUpdate.CourseId)
	if result.Error != nil {
		return fmt.Errorf("course not found: %w", result.Error)
	}

	course.Prerequisite = []curriculumModel.Course{}
	result = courseUpdate.Connector.Save(&course)
	if result.Error != nil {
		return fmt.Errorf("failed to update course prerequisites: %w", result.Error)
	}

	return nil
}
