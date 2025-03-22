package controller

import (
	"ModEd/curriculum/model"
	"errors"

	"gorm.io/gorm"
)

type ICourseController interface {
	// Put methods here
	// eg. CreateCourse(course *modelCurriculum.Course) error
	CreateCourse(course *model.Course) error
	GetCourseByID(id uint) (*model.Course, error)
	UpdateCourse(id uint, updatedCourse *model.Course) error
	DeleteCourse(id uint) error
	ListCourses() ([]model.Course, error)
}

type CourseController struct {
	db *gorm.DB
}

func NewCourseController(db *gorm.DB) ICourseController {
	return &CourseController{db: db}
}

// Create Course
func (c *CourseController) CreateCourse(course *model.Course) error {
	result := c.db.Create(course)
	return result.Error
}

// Get Course by ID
func (c *CourseController) GetCourseByID(id uint) (*model.Course, error) {
	var course model.Course
	result := c.db.Preload("Prerequisite").First(&course, id) // Preload prerequisites
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("course not found")
	}
	return &course, result.Error
}

// Update Course
func (c *CourseController) UpdateCourse(id uint, updatedCourse *model.Course) error {
	var course model.Course
	result := c.db.First(&course, id)
	if result.Error != nil {
		return errors.New("course not found")
	}

	// Update fields
	course.Name = updatedCourse.Name
	course.Description = updatedCourse.Description
	course.Optional = updatedCourse.Optional
	course.CourseStatus = updatedCourse.CourseStatus

	// Save updated course
	return c.db.Save(&course).Error
}

// Delete Course
func (c *CourseController) DeleteCourse(id uint) error {
	result := c.db.Delete(&model.Course{}, id)
	if result.RowsAffected == 0 {
		return errors.New("course not found")
	}
	return result.Error
}

// List All Courses
func (c *CourseController) ListCourses() ([]model.Course, error) {
	var courses []model.Course
	result := c.db.Preload("Prerequisite").Find(&courses)
	return courses, result.Error
}
