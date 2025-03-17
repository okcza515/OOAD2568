package controller

import (
	curriculumModel "ModEd/curriculum/model"
	model "ModEd/curriculum/model/instructor-workload"
	"fmt"

	"gorm.io/gorm"
)

type CourseUpdateController struct {
	Connector *gorm.DB
}

type CourseNameUpdate struct {
	CourseId      string
	NewCourseName string
}

type CoursePrerequisiteUpdate struct {
	CourseId              string
	NewPrerequisiteCourse string
}

func CreateCourseUpdateController(connector *gorm.DB) *CourseUpdateController {
	// connector.AutoMigrate(&curriculumModel.Course{})
	return &CourseUpdateController{
		Connector: connector,
	}
}

func (courseUpdate *CourseUpdateController) CourseUpdate(courseUpdateRequest *model.CourseUpdateRequest) error {
	var updater model.CourseUpdateStrategy

	switch courseUpdateRequest.RequestType {
	case model.NAME:
		// updater = &CourseNameUpdate{}
	case model.PREREQUISITE:
		// updater = &CoursePrerequisiteUpdate{}
	default:
		return fmt.Errorf("invalid request type")
	}

	courseUpdateRequest.SetCourseUpdateStrategy(updater)

	return courseUpdateRequest.ExecuteUpdate()
}

func (c *CourseNameUpdate) Update(connector *gorm.DB) error {
	var course curriculumModel.Course
	result := connector.First(&course, "id = ?", c.CourseId)
	if result.Error != nil {
		return fmt.Errorf("course not found: %w", result.Error)
	}
	course.Name = c.NewCourseName
	result = connector.Save(&course)
	if result.Error != nil {
		return fmt.Errorf("failed to update course name: %w", result.Error)
	}

	return nil
}

func (c *CoursePrerequisiteUpdate) Update(connector *gorm.DB) error {
	var course []curriculumModel.Course
	result := connector.First(&course, "id = ?", c.CourseId)
	if result.Error != nil {
		return fmt.Errorf("course not found: %w", result.Error)
	}
	result = connector.Save(&course)
	if result.Error != nil {
		return fmt.Errorf("failed to update course prerequisites: %w", result.Error)
	}

	return nil
}
