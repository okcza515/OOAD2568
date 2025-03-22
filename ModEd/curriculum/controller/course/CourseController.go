package controller

import (
	"ModEd/curriculum/model"

	"gorm.io/gorm"
)

type ICourseController interface {
	CreateCourse(course model.Course) (courseId uint, err error)
	GetCourseByID(courseId uint) (course *model.Course, err error)
	UpdateCourse(updatedCourse model.Course) (course *model.Course, err error)
	DeleteCourse(courseId uint) (course *model.Course, err error)
	ListCourses() (courses []*model.Course, err error)
}

type CourseController struct {
	db *gorm.DB
}

func NewCourseController(db *gorm.DB) ICourseController {
	return &CourseController{db: db}
}

func (c *CourseController) CreateCourse(course model.Course) (courseId uint, err error) {
	if err := c.db.Create(&course).Error; err != nil {
		return 0, err
	}
	return course.ID, nil
}

func (c *CourseController) GetCourseByID(courseId uint) (course *model.Course, err error) {
	course = &model.Course{}
	if err := c.db.Preload("Prerequisite").First(course, courseId).Error; err != nil {
		return nil, err
	}
	return course, nil
}

func (c *CourseController) UpdateCourse(updatedCourse model.Course) (course *model.Course, err error) {
	course = &model.Course{}
	if err := c.db.First(course, updatedCourse.ID).Error; err != nil {
		return nil, err
	}
	// Update fields
	course.Name = updatedCourse.Name
	course.Description = updatedCourse.Description
	course.Optional = updatedCourse.Optional
	course.CourseStatus = updatedCourse.CourseStatus

	if err := c.db.Save(course).Error; err != nil {
		return nil, err
	}
	return course, nil
}

func (c *CourseController) DeleteCourse(courseId uint) (course *model.Course, err error) {
	course = &model.Course{}
	if err := c.db.First(course, courseId).Error; err != nil {
		return nil, err
	}
	if err := c.db.Delete(course).Error; err != nil {
		return nil, err
	}
	return course, nil
}

func (c *CourseController) ListCourses() (courses []*model.Course, err error) {
	if err := c.db.Preload("Prerequisite").Find(&courses).Error; err != nil {
		return nil, err
	}
	return courses, nil
}
