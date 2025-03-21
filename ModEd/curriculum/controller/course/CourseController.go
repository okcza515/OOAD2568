package controller

import (
	"gorm.io/gorm"
)

type ICourseController interface {
	// Put methods here
	// eg. CreateCourse(course *modelCurriculum.Course) error
}

type CourseController struct {
	db *gorm.DB
}

func NewCourseController(db *gorm.DB) ICourseController {
	return &CourseController{db: db}
}
