package controller

import (
	"ModEd/curriculum/model"
	"ModEd/utils/deserializer"
	"fmt"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type CourseController struct {
	db *gorm.DB
}

func NewCourseController(db *gorm.DB) *CourseController {
	return &CourseController{db: db}
}

func (c *CourseController) CreateCourse(course *model.Course) (courseId uint, err error) {
	if err := c.db.Create(&course).Error; err != nil {
		return 0, err
	}
	return course.CourseId, nil
}

func (c *CourseController) GetCourseByID(courseId uint) (course *model.Course, err error) {
	course = &model.Course{}
	if err := c.db.Preload("Prerequisite").First(course, courseId).Error; err != nil {
		return nil, err
	}
	return course, nil
}

func (c *CourseController) UpdateCourse(updatedCourse *model.Course) (*model.Course, error) {
	course := &model.Course{}
	if err := c.db.First(course, updatedCourse.CourseId).Error; err != nil {
		return nil, err
	}

	// Update fields
	course.Name = updatedCourse.Name
	course.Description = updatedCourse.Description
	course.Optional = updatedCourse.Optional
	course.CourseStatus = updatedCourse.CourseStatus
	// course.Prerequisite = updatedCourse.Prerequisite
	// course.ClassList = updatedCourse.ClassList

	if err := c.db.Updates(course).Error; err != nil {
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

func (c *CourseController) CreateSeedCourse(path string) (courses []*model.Course, err error) {
	deserializer, err := deserializer.NewFileDeserializer(path)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create file deserializer")
	}

	if err := deserializer.Deserialize(&courses); err != nil {
		return nil, errors.Wrap(err, "failed to deserialize courses")
	}

	for _, course := range courses {
		_, err := c.CreateCourse(course)
		if err != nil {
			return nil, errors.Wrap(err, "failed to create course")
		}
	}
	fmt.Println("Create Course Seed Successfully")
	return courses, nil

}
